use std::{
    string,
    fs,
    env, fmt::format
};

const TXT_PATH : &str = "/home/benn/CODE/adventCode/";

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
