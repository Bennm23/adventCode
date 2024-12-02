use std::{
    fs,
    time::SystemTime,
};

const TXT_PATH : &str = "/home/benn/CODE/adventCode/2024/";

pub fn read_file_to_vec(day_file : &str) -> Vec<String> {
    let file_path = format!("{TXT_PATH}{day_file}");
    let contents = fs::read_to_string(file_path)
        .expect("Could Not Parse File");

    let lines: Vec<String> = contents.split("\n")
        .map(|s : &str| s.to_string())
        .collect();
    lines
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


pub fn run_and_print_duration(solver : &dyn Fn()) {
    let start = match SystemTime::now().duration_since(SystemTime::UNIX_EPOCH) {
        Ok(n) => n.as_micros(),
        Err(_) => panic!("oops")
    };
    
    solver();

    let end = match SystemTime::now().duration_since(SystemTime::UNIX_EPOCH) {
        Ok(n) => n.as_micros(),
        Err(_) => panic!("oops")
    };

    println!("TIme = {}", (end - start));
}