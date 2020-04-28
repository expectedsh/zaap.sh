use tonic::{Request, Status};
use tonic::metadata::MetadataValue;

pub fn check_auth(raw_token: &str) -> impl Fn(Request<()>) -> Result<Request<()>, Status> {
    let token = MetadataValue::from_str(raw_token).unwrap();

    return move |req| {
        match req.metadata().get("authorization") {
            Some(t) if token == t => Ok(req),
            _ => Err(Status::unauthenticated("Invalid token.")),
        }
    }
}
