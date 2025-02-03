mod packageA;
use packageA::{testPackageA , forloop, ifelse};

fn main() {
    println!("Hello, world!");
    A();
    testPackageA();
    forloop();
    ifelse();
}

fn 
A() { let c = 5 + 8; println!("100 {c} " ); }