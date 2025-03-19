// client.rs
use tokio::net::tcp::TcpStream;
use tokio::prelude::*;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
struct Message {
    topic: String,
    data: String,
}

struct Client {
    stream: TcpStream,
}

impl Client {
    async fn new() -> Self {
        let stream = TcpStream::connect("127.0.0.1:8080").await.unwrap();
        Client { stream }
    }

    async fn send(&mut self, topic: String, data: String) {
        let msg = Message { topic, data };
        let json = serde_json::to_string(&msg).unwrap();
        self.stream.write_all(json.as_bytes()).await.unwrap();
    }

    async fn receive(&mut self) -> (String, String) {
        let mut buf = [0; 1024];
        match self.stream.read(&mut buf).await {
            Ok(n) => {
                let msg: Message = serde_json::from_slice(&buf[..n]).unwrap();
                (msg.topic, msg.data)
            }
            Err(e) => {
                println!("Error reading from connection: {}", e);
                (String::new(), String::new())
            }
        }
    }
}

#[tokio::main]
async fn main() {
    let mut client = Client::new().await;
    client.send("topic1".to_string(), "Hello, world!".to_string()).await;
    let (topic, data) = client.receive().await;
    println!("Received on topic {}: {}", topic, data);
}