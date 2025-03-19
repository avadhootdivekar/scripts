use std::sync::mpsc;
use std::thread;
use std::time;

// Function to write to the channel
pub fn write_to_channel(tx: mpsc::Sender<i32>) {
    for i in 0..10 {
        thread::sleep(time::Duration::from_secs(1));
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