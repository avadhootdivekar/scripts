extern crate rand;

use rand::prelude::*;

pub fn testPackageA() {
    println!("This is test package A.")
}

pub fn forloop(){
    for i in 0..10 {
        println!("{}" , i);
    }
}

pub fn ifelse(){
    let a: u32 = rand::thread_rng().gen_range(0..=10);
    if a > 5 {
        println!("a is greater than 5");
    }
}