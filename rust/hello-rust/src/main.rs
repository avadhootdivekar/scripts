mod packageA;
mod packageB;
mod tokenchecker;

use std::{thread};
use std::sync::mpsc;
use packageA::{testPackageA , forloop, ifelse};
use packageB::{write_to_channel, read_from_channel};
use tokenchecker::checker;

fn main() {
    println!("Hello, world!");
    A();
    testPackageA();
    forloop();
    ifelse();
    B();
}

fn 
A() { let c = 5 + 8; println!("100 {c} " ); }

fn B(){
    let (tx, rx) = mpsc::channel();
    println!("In function B");
    // Spawn a new thread to write to the channel
    thread::spawn(move || write_to_channel(tx));
    // Read from the channel in the main thread
    read_from_channel(rx);
}