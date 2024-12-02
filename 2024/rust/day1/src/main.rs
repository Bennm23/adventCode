use core::fmt;
use std::{collections::HashMap, fmt::Error, num::ParseIntError, str::FromStr};

use utils::read_file_to_vec;

extern crate utils;

fn main() {
    println!("Hello, world!");

    let file = read_file_to_vec("day1.txt");

    let mut left : Vec<i32> = Vec::new();
    let mut right : Vec<i32> = Vec::new();

    let mut map : HashMap<i32, i32> = HashMap::new();

    for l in file {
        let split : Vec<&str> = l.split(" ").collect();

        let mut lval = -1;
        let mut rval = -1;
        for s in split {
            let val: Result<i32, ParseIntError> = FromStr::from_str(s);

            if val.is_ok() {
                if lval < 0 {
                    lval = val.unwrap();
                } else {
                    rval = val.unwrap();
                }
            }
        }

        if lval > 0 && rval > 0 {
            left.push(lval);
            right.push(rval);
        }
    }

    left.sort();
    right.sort();

    for l in &left {
        map.insert(*l, 0);
    }

    for r in &right {
        if map.contains_key(&r) {
            map.insert(*r, map.get(&r).unwrap() + 1);
        }
    }

    let mut distance = 0;
    let mut similarity = 0;
    for i in 0 .. left.len() {
        println!("{:?}  {:?}", left[i], right[i]);
        distance += (left[i] - right[i]).abs();

        similarity += left[i] * map.get(&left[i]).unwrap();
    }

    println!("Length = {:?}, {:?}", left.len(), right.len());
    println!("Distance = {distance}");
    println!("Similarity = {similarity}");
}
