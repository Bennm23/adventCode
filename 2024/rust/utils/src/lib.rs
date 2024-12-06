use std::{
    fmt::Display, fs, time::SystemTime
};

use home::home_dir;

pub fn get_advent_path(day_file : &str) -> String {
    
    match home::home_dir() {
        Some(path) => {
            return format!("{}/CODE/adventCode/2024/inputs/{}", path.display(), day_file)
        }
        None => panic!("Home Dir not set")
    }
}

pub fn read_file_to_vec(day_file : &str) -> Vec<String> {
    let file_path = get_advent_path(day_file);
    let contents = fs::read_to_string(file_path)
        .expect("Could Not Parse File");

    let lines: Vec<String> = contents.split("\n")
        .map(|s : &str| s.to_string())
        .collect();
    lines
}

pub fn read_file_to_grid(day_file : &str) -> Vec<Vec<char>> {
    let file_path = get_advent_path(day_file);
    let contents = fs::read_to_string(file_path)
        .expect("Could Not Parse File");

    let lines: Vec<Vec<char>> = contents.split("\n")
        .map(|s : &str| s.chars().collect())
        .filter(|f : &Vec<char>| !f.is_empty())
        .collect();
    lines
}

const DEBUG : bool = false;

pub fn log(s : &str, newline : bool, always : bool) {
    if DEBUG || always {
        
        if newline {
            println!("{s}");
        } else {
            print!("{s}")
        }
    }

}

#[macro_export]
macro_rules! log {

    ($s:expr, $newline: expr) => {
        log($s, $newline, false);
    };
    ($s: expr) => {
        log($s, true, false);
    };
    () => {
        log("", true, false);
    };
}
pub fn logs(s : &String, newline : bool, always : bool) {
    if DEBUG || always {
        if newline {
            println!("{s}");
        } else {
            print!("{s}")
        }
    }

}

#[macro_export]
macro_rules! logs {

    ($s:expr, $newline: expr) => {
        logs($s, $newline, false);
    };
    ($s: expr) => {
        logs($s, true, false);
    };
    () => {
        logs("", true, false);
    };
}


// pub fn add(left: usize, right: usize) -> usize {
//     left + right
// }

// #[cfg(test)]
// mod tests {
//     use super::*;

//     #[test]
//     fn it_works() {
//         let result = add(2, 2);
//         assert_eq!(result, 4);
//     }
// }


pub fn run_and_print_duration(title : &str, solver : &dyn Fn()) {
    let start = match SystemTime::now().duration_since(SystemTime::UNIX_EPOCH) {
        Ok(n) => n.as_micros(),
        Err(_) => panic!("oops")
    };
    
    solver();

    let end = match SystemTime::now().duration_since(SystemTime::UNIX_EPOCH) {
        Ok(n) => n.as_micros(),
        Err(_) => panic!("oops")
    };

    println!("{title} Ran For = {} us", (end - start));
}

pub fn run_and_score<R : Display, F :Fn() -> R>(title : &str, solver : F) {
    let start = match SystemTime::now().duration_since(SystemTime::UNIX_EPOCH) {
        Ok(n) => n.as_micros(),
        Err(_) => panic!("oops")
    };
    
    let res = solver();

    let end = match SystemTime::now().duration_since(SystemTime::UNIX_EPOCH) {
        Ok(n) => n.as_micros(),
        Err(_) => panic!("oops")
    };

    println!("{title}: Result = {}. Ran For = {} us", res, (end - start));
}