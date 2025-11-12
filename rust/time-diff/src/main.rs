mod node;

fn main(){
    println!("Hello World");
    test();
    println!("test completed ");
}


fn test(){
    let mut n = node::Node::new(10);
    n.insert(5);
    n.insert(15);
    n.insert(3);
    n.insert(7);
    n.insert(12);
    n.insert(18);
    let s = n.stringify();
    println!("Initial tree: {}", s);
}