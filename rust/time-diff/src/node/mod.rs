
pub struct Node {
    Value: i32,
    Left :Option<Box<Node>>,
    Right :Option<Box<Node>>,
}

impl Node {
    pub fn new(value: i32) -> Self {
        Node {
            Value: value,
            Left: None,
            Right: None,
        }
    }

    pub fn insert(&mut self, value: i32) {
        if value < self.Value {
            match self.Left {
                Some(ref mut left) => left.insert(value),
                None => self.Left = Some(Box::new(Node::new(value))),
            }
        } else {
            match self.Right {
                Some(ref mut right) => right.insert(value),
                None => self.Right = Some(Box::new(Node::new(value))),
            }
        }
    }
    pub fn stringify(&self) -> String {
        let left_str = match &self.Left {
            Some(left) => left.stringify(),
            None => String::from(""),
        };
        let right_str = match &self.Right {
            Some(right) => right.stringify(),
            None => String::from(""),
        };
        format!("[{} {} {}]", left_str, self.Value, right_str)
    }
}