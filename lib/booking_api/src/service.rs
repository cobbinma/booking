use async_trait::async_trait;
use chrono::{DateTime, Duration};
use protobuf::booking::api::booking_api_server::BookingApi;
use protobuf::booking::api::GetSlotResponse;
use protobuf::booking::models::{Booking, Slot, SlotInput};
use std::ops::Add;
use tonic::{Request, Response, Status};
use uuid::Uuid;

#[derive(Debug, Default)]
pub struct BookingService {}

#[async_trait]
impl BookingApi for BookingService {
    async fn get_slot(&self, req: Request<SlotInput>) -> Result<Response<GetSlotResponse>, Status> {
        let slot = req.into_inner();
        let starts_at = slot.starts_at;
        let duration = slot.duration;
        let ends_at = get_ends_at(&starts_at, duration)?;
        let venue_id = slot.venue_id;
        let email = slot.email;
        let people = slot.people;
        Ok(Response::new(GetSlotResponse {
            r#match: Some(Slot {
                venue_id,
                email,
                people,
                starts_at,
                ends_at,
                duration,
            }),
            other_available_slots: vec![],
        }))
    }

    async fn create_booking(&self, req: Request<SlotInput>) -> Result<Response<Booking>, Status> {
        let slot = req.into_inner();
        let starts_at = slot.starts_at;
        let duration = slot.duration;
        let ends_at = get_ends_at(&starts_at, duration)?;
        let venue_id = slot.venue_id;
        let email = slot.email;
        let people = slot.people;
        Ok(Response::new(Booking {
            id: Uuid::new_v4().to_string(),
            venue_id,
            email,
            people,
            starts_at,
            ends_at,
            duration,
            table_id: Uuid::new_v4().to_string(),
        }))
    }
}

fn get_ends_at(starts_at: &str, duration: u32) -> Result<String, Status> {
    DateTime::parse_from_rfc3339(starts_at)
        .map(|dt| dt.add(Duration::minutes(duration as i64)).to_rfc3339())
        .map_err(|e| Status::internal(e.to_string()))
}
