extern crate utils;

use std::{
    string,
};

use regex::Regex;

fn main() {
    solve()
}

const RMAX : i32 = 12;
const GMAX : i32 = 13;
const BMAX : i32 = 14;

fn solve() {

    let lines = utils::read_file_to_vec("day2.txt");

    let mut score : i32 = 0;

    let mut rc : i32 = 0;
    for line in lines {
        rc += 1;

        println!("{line}");
        if gamePasses(&line) {
            score += rc;
        }



        break;
        
    }

    //part1 find if any game count exceeds the max for that color

}

fn gamePasses(game : &String) -> bool {

    let re = Regex::new(r"\d+").expect("Failed to create Regex");

    for cap in re.captures_iter(game) {

        println!("Cap = {:?}", cap);
    }


    false
}
