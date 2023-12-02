extern crate utils;

use std::time::SystemTime;

const RADIX : u32 = 10;

fn main() {
    let start = match SystemTime::now().duration_since(SystemTime::UNIX_EPOCH) {
        Ok(n) => n.as_micros(),
        Err(_) => panic!("oops")
    };

    let lines = utils::read_file_to_vec("day1.txt");

    println!("Part 1 = {}", part1(&lines));//55130
    println!("Part 2 = {}", part2(&lines));//54985
    let end = match SystemTime::now().duration_since(SystemTime::UNIX_EPOCH) {
        Ok(n) => n.as_micros(),
        Err(_) => panic!("oops")
    };

    println!("TIme = {}", (end - start));

}

fn part1(lines : &Vec<String>) -> u32 {
    let mut sum : u32 = 0;

    for line in lines {
        sum += get_combined_str(line);
    }

    sum
}

fn part2(lines : &Vec<String>) -> u32 {
    let mut sum : u32 = 0;

    for line in lines {

        let result = line
            .replace("one", "o1e")
            .replace("two", "t2o")
            .replace("three", "t3e")
            .replace("four", "f4r")
            .replace("five", "f5e")
            .replace("six", "s6x")
            .replace("seven", "s7n")
            .replace("eight", "e8t")
            .replace("nine", "n9e");
        
        sum += get_combined_str(&result.to_string());
    }

    sum
}

fn get_combined_str(line : &String) -> u32 {

    let mut first : u32 = 0;
    let mut last : u32 = 0;

    
    for c in line.chars() {

        if c.is_ascii_digit() && first == 0 {
            first = c.to_digit(RADIX).expect(&format!("Could Not Parse {c}"));
            last = first;
            continue;
        }

        if c.is_ascii_digit() {
            last = c.to_digit(RADIX).expect(&format!("Could Not Parse {c}"));
        }
    }
    
    let mut combined = first.to_string();
    combined.push_str(&last.to_string());
    combined.parse::<u32>().expect(&format!("Could Not Convert {combined} to Int"))
}