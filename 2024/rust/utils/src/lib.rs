use std::{
    fmt::Display, fs, time::SystemTime
};

use regex::Regex;

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
        .filter(|f : &String| !f.is_empty())
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
pub fn read_file_to_groups(day_file : &str, seperator : &str) -> Vec<Vec<String>> {
    let file_path = get_advent_path(day_file);
    let contents = fs::read_to_string(file_path)
        .expect("Could Not Parse File");

    let mut lines : Vec<Vec<String>> = Vec::new();

    let mut subgroup : Vec<String> = Vec::new();

    for line in contents.split("\n") {

        if line == seperator {
            lines.push(subgroup.clone());
            subgroup = Vec::new();
        } else {
            subgroup.push(line.to_owned());
        }
    }
    lines.push(subgroup.clone());
    lines
}

pub fn read_file_to_int_grid(day_file : &str) -> Vec<Vec<i32>> {
    let file_path = get_advent_path(day_file);
    let contents = fs::read_to_string(file_path)
        .expect("Could Not Parse File");

    let lines: Vec<Vec<i32>> = contents.split("\n")
        .filter(|s | !s.is_empty())
        .map(|s : &str| convert_string_to_ints(s))
        .collect();
    lines
}

pub fn convert_string_to_ints(string : &str) -> Vec<i32> {

    return string.chars()
        .map(|c: char| c.to_digit(10).expect("Failed to parse int") as i32)
        .collect()
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
pub fn run_and_score_both<R : Display, F :Fn() -> (R, R)>(solver : F) {
    let start = match SystemTime::now().duration_since(SystemTime::UNIX_EPOCH) {
        Ok(n) => n.as_micros(),
        Err(_) => panic!("oops")
    };
    
    let (p1, p2) = solver();

    let end = match SystemTime::now().duration_since(SystemTime::UNIX_EPOCH) {
        Ok(n) => n.as_micros(),
        Err(_) => panic!("oops")
    };

    println!("Part 1: Result = {}.", p1);
    println!("Part 2: Result = {}.", p2);
    println!("Took {} us", (end - start));
}

pub fn remove_between_or_after(text: &str, start: &str, end: &str) -> String {
    // Build the non-greedy regex pattern
    let pattern = format!(r"{}.*?{}|{}.*", regex::escape(start), regex::escape(end), regex::escape(start));
    let re = Regex::new(&pattern).unwrap();
    re.replace_all(text, "").to_string()
}

#[derive(Debug, Clone, Copy, Hash, PartialEq, Eq)]
pub struct Pair(pub i32, pub i32);

impl Pair {
    fn evaluate_for<T: Copy>(&self, grid: &Vec<Vec<T>>) -> T {
        grid[self.0 as usize][self.1 as usize]
    }
    pub fn out_of_bounds(&self, size : usize) -> bool {
        self.0 < 0 || self.0 >= size as i32 || self.1 < 0 || self.1 >= size as i32
    }

    pub fn add(&self, other: &Pair) -> Self {

        Self (
            self.0 + other.0,
            self.1 + other.1,
        )
    }

    pub fn ia(&self) -> usize {
        self.0 as usize
    }
    pub fn ib(&self) -> usize {
        self.1 as usize
    }
}