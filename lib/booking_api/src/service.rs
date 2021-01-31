use crate::models;
use crate::postgres::Postgres;
use async_trait::async_trait;
use chrono::format::Numeric::Timestamp;
use chrono::{DateTime, Datelike, Duration, NaiveDate, NaiveTime, TimeZone, Timelike, Utc};
use protobuf::booking::api::booking_api_server::BookingApi;
use protobuf::booking::api::GetSlotResponse;
use protobuf::booking::models::{Booking, Slot, SlotInput};
use protobuf::venue::api::table_api_client::TableApiClient;
use protobuf::venue::api::table_api_server::TableApi;
use protobuf::venue::api::venue_api_client::VenueApiClient;
use protobuf::venue::api::venue_api_server::VenueApi;
use protobuf::venue::api::{GetTablesRequest, GetVenueRequest};
use std::borrow::Borrow;
use std::collections::HashMap;
use std::ops::{Add, Deref};
use std::sync::Arc;
use tonic::{Request, Response, Status};
use uuid::Uuid;

pub trait Repository {
    fn get_bookings_by_date(
        &self,
        venue: &Uuid,
        day: &NaiveDate,
    ) -> Result<Vec<models::Booking>, Status>;

    fn create_booking(&self, new_booking: &models::Booking) -> Result<(), Status>;
}

pub struct BookingService {
    repository: Box<dyn Repository + Send + Sync + 'static>,
    venue_client: Box<VenueApiClient<tonic::transport::Channel>>,
    table_client: Box<TableApiClient<tonic::transport::Channel>>,
}

impl BookingService {
    pub fn new(
        repository: Box<dyn Repository + Send + Sync + 'static>,
        venue_client: Box<VenueApiClient<tonic::transport::Channel>>,
        table_client: Box<TableApiClient<tonic::transport::Channel>>,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        Ok(BookingService {
            repository,
            venue_client,
            table_client,
        })
    }
}

#[async_trait]
impl BookingApi for BookingService {
    async fn get_slot(&self, req: Request<SlotInput>) -> Result<Response<GetSlotResponse>, Status> {
        let slot = req.into_inner();

        let slot_length = slot.duration as i64;

        let starts = DateTime::parse_from_rfc3339(&slot.starts_at).map_err(|e| {
            log::error!("could not parse date : {}", e);
            Status::internal("could not parse date")
        })?;
        let day = NaiveDate::from_ymd(starts.year(), starts.month(), starts.day());

        let venue = &self
            .venue_client
            .clone()
            .get_venue(GetVenueRequest {
                id: slot.venue_id.clone(),
            })
            .await?
            .into_inner();

        let opening_hours_specification = venue
            .opening_hours
            .iter()
            .filter(|&hours| hours.day_of_week == day.weekday().number_from_monday())
            .next()
            .ok_or_else(|| Status::invalid_argument("venue not open on given date"))?;

        let opens = NaiveTime::parse_from_str(&opening_hours_specification.opens, "%H:%M")
            .map_err(|e| {
                log::error!("could not parse opens time : {}", e);
                Status::internal("could not parse opens time")
            })
            .map(|o| {
                Utc.ymd(starts.year(), starts.month(), starts.day())
                    .and_hms(o.hour(), o.minute(), o.second())
            })?;

        let closes = NaiveTime::parse_from_str(&opening_hours_specification.closes, "%H:%M")
            .map_err(|e| {
                log::error!("could not parse closes time : {}", e);
                Status::internal("could not parse closes time")
            })
            .map(|c| {
                Utc.ymd(starts.year(), starts.month(), starts.day())
                    .and_hms(c.hour(), c.minute(), c.second())
            })?;

        if starts < opens || starts + Duration::minutes(slot_length) > closes {
            return Err(Status::invalid_argument("venue is closed"));
        }

        let tables_with_capacity: Vec<String> = self
            .table_client
            .clone()
            .get_tables(GetTablesRequest {
                venue_id: slot.venue_id.clone(),
            })
            .await?
            .into_inner()
            .tables
            .iter()
            .filter(|table| table.capacity >= slot.people)
            .map(|table| table.id.clone())
            .collect();

        if tables_with_capacity.is_empty() {
            return Err(Status::invalid_argument(
                "restaurant does not have tables that large",
            ));
        }

        log::info!("getting bookings from database");
        let bookings: Vec<models::Booking> = self
            .repository
            .get_bookings_by_date(
                &Uuid::parse_str(&slot.venue_id).map_err(|e| {
                    log::error!("could not parse uuid : {}", e);
                    Status::internal("could not parse uuid")
                })?,
                &day,
            )?
            .iter()
            .filter(|&booking| tables_with_capacity.contains(&booking.table_id.to_string()))
            .map(|booking| booking.clone())
            .collect();

        let mut free_time_slots = HashMap::new();
        let mut t = opens;
        while t <= closes - Duration::minutes(slot_length) {
            let free_table_id = tables_with_capacity
                .iter()
                .filter(|table_id| {
                    bookings
                        .iter()
                        .filter(|booking| booking.table_id.to_string() == **table_id)
                        .all(|b| {
                            !(t < b.ends_at && b.starts_at < t + Duration::minutes(slot_length))
                        })
                })
                .next();

            if let Some(id) = free_table_id {
                free_time_slots.insert(t, id);
            }

            t = t + Duration::minutes(30);
        }

        let other_available_slots: Vec<Slot> = free_time_slots
            .iter()
            .map(|(time, table_id)| Slot {
                venue_id: slot.venue_id.clone(),
                email: slot.email.clone(),
                people: slot.people,
                starts_at: time.to_rfc3339(),
                ends_at: (*time + Duration::minutes(slot_length)).to_rfc3339(),
                duration: slot.duration,
            })
            .collect();

        Ok(Response::new(GetSlotResponse {
            r#match: free_time_slots
                .get(&starts.with_timezone(&Utc))
                .map(|_| Slot {
                    venue_id: slot.venue_id,
                    email: slot.email,
                    people: slot.people,
                    starts_at: slot.starts_at,
                    ends_at: (starts + Duration::minutes(slot_length)).to_rfc3339(),
                    duration: slot.duration,
                }),
            other_available_slots,
        }))
    }

    async fn create_booking(&self, req: Request<SlotInput>) -> Result<Response<Booking>, Status> {
        let slot = req.into_inner();

        let slot_length = slot.duration as i64;

        let starts = DateTime::parse_from_rfc3339(&slot.starts_at).map_err(|e| {
            log::error!("could not parse date : {}", e);
            Status::internal("could not parse date")
        })?;
        let day = NaiveDate::from_ymd(starts.year(), starts.month(), starts.day());

        let venue = &self
            .venue_client
            .clone()
            .get_venue(GetVenueRequest {
                id: slot.venue_id.clone(),
            })
            .await?
            .into_inner();

        let opening_hours_specification = venue
            .opening_hours
            .iter()
            .filter(|&hours| hours.day_of_week == day.weekday().number_from_monday())
            .next()
            .ok_or_else(|| Status::invalid_argument("venue not open on given date"))?;

        let opens = NaiveTime::parse_from_str(&opening_hours_specification.opens, "%H:%M")
            .map_err(|e| {
                log::error!("could not parse opens time : {}", e);
                Status::internal("could not parse opens time")
            })
            .map(|o| {
                Utc.ymd(starts.year(), starts.month(), starts.day())
                    .and_hms(o.hour(), o.minute(), o.second())
            })?;

        let closes = NaiveTime::parse_from_str(&opening_hours_specification.closes, "%H:%M")
            .map_err(|e| {
                log::error!("could not parse closes time : {}", e);
                Status::internal("could not parse closes time")
            })
            .map(|c| {
                Utc.ymd(starts.year(), starts.month(), starts.day())
                    .and_hms(c.hour(), c.minute(), c.second())
            })?;

        if starts < opens || starts + Duration::minutes(slot_length) > closes {
            return Err(Status::invalid_argument("venue is closed"));
        }

        let tables_with_capacity: Vec<String> = self
            .table_client
            .clone()
            .get_tables(GetTablesRequest {
                venue_id: slot.venue_id.clone(),
            })
            .await?
            .into_inner()
            .tables
            .iter()
            .filter(|table| table.capacity >= slot.people)
            .map(|table| table.id.clone())
            .collect();

        if tables_with_capacity.is_empty() {
            return Err(Status::invalid_argument(
                "restaurant does not have tables that large",
            ));
        }

        let bookings: Vec<models::Booking> = self
            .repository
            .get_bookings_by_date(
                &Uuid::parse_str(&slot.venue_id).map_err(|e| {
                    log::error!("could not parse uuid : {}", e);
                    Status::internal("could not parse uuid")
                })?,
                &day,
            )?
            .iter()
            .filter(|&booking| tables_with_capacity.contains(&booking.table_id.to_string()))
            .map(|booking| booking.clone())
            .collect();

        let free_table_id = tables_with_capacity
            .iter()
            .filter(|table_id| {
                bookings
                    .iter()
                    .filter(|booking| booking.table_id.to_string() == **table_id)
                    .all(|b| {
                        !(starts < b.ends_at
                            && b.starts_at < starts + Duration::minutes(slot_length))
                    })
            })
            .next();

        if let Some(table_id) = free_table_id {
            let id = uuid::Uuid::new_v4();
            let new_booking = models::Booking {
                id: id.clone(),
                customer_email: slot.email.clone(),
                venue_id: Uuid::parse_str(&slot.venue_id)
                    .map_err(|_| Status::invalid_argument("could not parse uuid"))?,
                table_id: Uuid::parse_str(table_id)
                    .map_err(|_| Status::internal("could not parse table uuid"))?,
                people: slot.people as i32,
                date: day,
                starts_at: starts.with_timezone(&Utc),
                ends_at: starts
                    .with_timezone(&Utc)
                    .add(Duration::minutes(slot_length)),
                duration: slot.duration as i32,
            };
            log::info!("{:?}", new_booking);
            self.repository.create_booking(&new_booking)?;

            Ok(Response::new(Booking {
                id: id.to_string(),
                venue_id: slot.venue_id.clone(),
                email: slot.email.clone(),
                people: slot.people,
                starts_at: slot.starts_at,
                ends_at: (starts + Duration::minutes(slot_length)).to_rfc3339(),
                duration: slot.duration,
                table_id: table_id.to_string(),
            }))
        } else {
            Err(Status::not_found("could not find a free slot"))
        }
    }
}

fn get_ends_at(starts_at: &str, duration: u32) -> Result<String, Status> {
    DateTime::parse_from_rfc3339(starts_at)
        .map(|dt| dt.add(Duration::minutes(duration as i64)).to_rfc3339())
        .map_err(|e| Status::internal(e.to_string()))
}