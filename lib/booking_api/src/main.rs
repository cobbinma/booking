use alcoholic_jwt::{token_kid, validate, Validation, JWKS};
use async_trait::async_trait;
use chrono::{DateTime, Duration};
use protobuf::booking::api::booking_api_server::{BookingApi, BookingApiServer};
use protobuf::booking::api::GetSlotResponse;
use protobuf::booking::models::{Booking, Slot, SlotInput};
use std::ops::Add;
use tonic::transport::{Identity, Server, ServerTlsConfig};
use tonic::{Request, Response, Status};
use uuid::Uuid;

#[derive(Debug, Default)]
pub struct BookingService {}

fn get_ends_at(starts_at: &str, duration: u32) -> Result<String, Status> {
    DateTime::parse_from_rfc3339(starts_at)
        .map(|dt| dt.add(Duration::minutes(duration as i64)).to_rfc3339())
        .map_err(|e| Status::internal(e.to_string()))
}

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

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    femme::with_level(femme::LevelFilter::Info);

    dotenv::dotenv().ok();

    let cert = tokio::fs::read("localhost.crt").await?;
    let key = tokio::fs::read("localhost.key").await?;

    let identity = Identity::from_pem(cert, key);

    let addr = "[::1]:6969".parse()?;
    let service = BookingService::default();

    log::info!("listening on port {}", &addr);

    Server::builder()
        .tls_config(ServerTlsConfig::new().identity(identity))?
        .add_service(BookingApiServer::with_interceptor(service, check_auth))
        .serve(addr)
        .await?;

    Ok(())
}

fn check_auth(req: Request<()>) -> Result<Request<()>, Status> {
    let md = req
        .metadata()
        .get("authorization")
        .ok_or_else(|| Status::unauthenticated("no valid auth token"))?;

    let token = md
        .to_str()
        .map_err(|_| Status::invalid_argument("could not parse token"))?;

    if validate_token(token)? {
        Ok(req)
    } else {
        Err(Status::unauthenticated("could not valididate auth token"))
    }
}

pub fn validate_token(token: &str) -> Result<bool, Status> {
    let token = token.trim_start_matches("Bearer ");
    let authority = std::env::var("AUTHORITY").expect("AUTHORITY must be set");
    let jwks = fetch_jwks(&format!(
        "{}{}",
        authority.as_str(),
        ".well-known/jwks.json"
    ))?;
    let validations = vec![Validation::Issuer(authority), Validation::SubjectPresent];
    let kid = token_kid(&token)
        .map_err(|_| Status::internal("failed to fetch jwts"))?
        .ok_or_else(|| Status::invalid_argument("failed to decode kid"))?;
    let jwk = jwks
        .find(&kid)
        .ok_or_else(|| Status::invalid_argument("specified key not found in set"))?;
    let res = validate(token, jwk, validations);

    Ok(res.is_ok())
}

fn fetch_jwks(uri: &str) -> Result<JWKS, Status> {
    let mut res = reqwest::get(uri).map_err(|_| Status::internal("could not get jwks"))?;
    let val = res
        .json::<JWKS>()
        .map_err(|_| Status::internal("could not unmarshall jwks"))?;

    Ok(val)
}