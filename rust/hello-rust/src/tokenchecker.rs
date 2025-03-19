use std::collections::HashMap;
use std::time::{Duration, Instant};
use std::string::String;

pub struct checker {
    
    map : HashMap <String, Instant>,
}

// Implementation of checker - these are methods on the checker type. 
impl checker {
    // fn functionName(argName:argType , argName:argType) -> (returnType, returnType) {}
    pub fn check(&mut self, token: String) -> (Duration, bool) {
        if self.map.contains_key(&token) {
            let d = self.map[&token] - Instant::now();
            return (d, true);
        } else {
            self.map.insert(token, Instant::now());
            return (Duration::from_millis(0), false);
        }
    }

}

pub fn newChecker() -> checker {
    checker {
        map: HashMap::new(),
    }
}

// Write unit tests here 
// #[cfg(test)]  <--  This is required attribute/macro. Check for attribue/macro in rust. 
#[cfg(test)]
// mod name is fixed to "tests" as required for unit tests by rust
mod tests {
    // use crate::tokenchecker;
    use super::*;
    
    // #[test] is required attribute/macro for test functions
    #[test]
    fn test_check() {
        let mut c = newChecker();
        let t = String::from("token");
        let t2 = String::from("token");
        assert_eq!(c.check(t), (Duration::from_millis(0), false));
        assert_eq!(c.check(t2), (Duration::from_millis(0), true));
        println!("test_check completed");
    }
}