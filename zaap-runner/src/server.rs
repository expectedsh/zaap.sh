use tonic::{Request, Response, Status};
use protocol::runner_server::Runner;
use protocol::{GetConfigurationRequest, GetConfigurationResponse, RunnerType};

#[derive(Debug, Default)]
pub struct KubernetesRunner {}

pub mod protocol {
    tonic::include_proto!("protocol");
}

#[tonic::async_trait]
impl Runner for KubernetesRunner {
    async fn get_configuration(
        &self,
        request: Request<GetConfigurationRequest>,
    ) -> Result<Response<GetConfigurationResponse>, Status> {
        println!("Got a request: {:?}", request);

        let reply = protocol::GetConfigurationResponse {
            r#type: RunnerType::DockerSwarm.into(),
            external_ips: vec!["dommage".to_string()].into(),
        };

        Ok(Response::new(reply))
    }
}
