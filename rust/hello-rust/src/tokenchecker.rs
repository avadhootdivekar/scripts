use std::collections::HashMap;
use std::time::{Duration, Instance};
struct checker {
    map : HashMap<str, Instance>,
}

// Implementation of checker - these are methods on the checker type. 
impl checker {
    // fn functionName(argName:argType , argName:argType) -> returnType, returnType {}
    fn check(me: &self, token: str) -> Duration, bool {
        if me.map.contains_key(token) {
            d = me.map[token] - Instance::now();
            return d, true;
        } else {
            me.map[token] = Instance::now();
            return Duration::from_millis(0), false;
        }
    }

}

pub fn newChecker() -> &checker {
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
        assert_eq!(c.check("token"), (Duration::from_millis(0), false));
        assert_eq!(c.check("token"), (Duration::from_millis(0), true));
        println!("test_check completed");
    }
}