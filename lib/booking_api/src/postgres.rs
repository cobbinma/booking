#[warn(unused_imports)]
use crate::diesel::RunQueryDsl;

use crate::models::Booking;
use crate::schema::bookings::dsl::bookings;
use crate::service::Repository;
use chrono::{DateTime, Datelike, Duration, NaiveDate};
use diesel::r2d2::{ConnectionManager, Pool};
use diesel::{ExpressionMethods, PgConnection, QueryDsl};
use protobuf::booking::models::SlotInput;
use std::env;
use tonic::Status;
use uuid::Uuid;

pub struct Postgres {
    pool: Pool<ConnectionManager<PgConnection>>,
}

impl Postgres {
    pub fn new() -> Result<Self, Box<dyn std::error::Error>> {
        let database_url = env::var("DATABASE_URL")?;
        let manager = ConnectionManager::<PgConnection>::new(database_url);
        let pool = diesel::r2d2::Builder::new().build(manager)?;

        Ok(Postgres { pool })
    }
}

impl Repository for Postgres {
    fn get_bookings_by_date(&self, venue: &Uuid, day: &NaiveDate) -> Result<Vec<Booking>, Status> {
        use crate::schema::bookings::dsl::*;

        let results: Vec<Booking> = bookings
            .filter(venue_id.eq(&venue))
            .filter(date.eq(&day))
            .get_results(&self.pool.get().map_err(|e| {
                log::error!("could not get database connection : {}", e);
                Status::internal("could not get database connection")
            })?)
            .map_err(|e| {
                log::error!("could not get get bookings from database : {}", e);
                Status::internal("could not get get bookings from database")
            })?;

        Ok(results)
    }

    fn create_booking(&self, new_booking: &Booking) -> Result<(), Status> {
        use crate::schema::bookings;
        diesel::insert_into(bookings::table)
            .values(new_booking)
            .execute(&self.pool.get().map_err(|e| {
                log::error!("{}", e);
                Status::internal("could not get database connection")
            })?)
            .map_err(|e| {
                log::error!("{}", e);
                Status::internal("could not create booking in database")
            })?;

        Ok(())
    }
}
