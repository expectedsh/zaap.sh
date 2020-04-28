use tonic::transport::Server;
use zaap_runner::server::{KubernetesRunner, protocol::runner_server::RunnerServer};
use zaap_runner::auth::check_auth;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:8090".parse()?;
    let runner = KubernetesRunner::default();

    println!("listening on {}", addr);
    Server::builder()
        .add_service(RunnerServer::with_interceptor(runner, check_auth("haricot")))
        .serve(addr)
        .await?;

    Ok(())
}
