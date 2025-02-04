use std::sync::mpsc;
use std::thread;

// Function to write to the channel
pub fn write_to_channel(tx: mpsc::Sender<i32>) {
    for i in 0..10 {
        tx.send(i).unwrap();
        println!("Sent: {}", i);
    }
}

// Function to read from the channel
pub fn read_from_channel(rx: mpsc::Receiver<i32>) {
    for _ in 0..10 {
        match rx.recv() {
            Ok(msg) => println!("Received: {}", msg),
            Err(_) => println!("Channel closed"),
        }
    }
}