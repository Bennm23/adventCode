use std::collections::HashMap;

use utils::{log, read_file_to_vec, run_and_score_both};

fn main() {
//  Part 1: Result = 306.
//  Part 2: Result = 604622004681855.
//  Took 11705608 us
    run_and_score_both(solve);
}

fn build_input() -> (Vec<Vec<char>>, Vec<Vec<char>>) {

    let lines = read_file_to_vec("day19.txt");

    let options = lines[0].split(", ").map(|s| s.chars().collect()).collect();

    let mut combinations : Vec<Vec<char>> = Vec::new();
    for i in 1 .. lines.len() {
        combinations.push(lines[i].chars().collect());
    }

    (combinations, options)
}

fn count_possibilities(
    options : &Vec<Vec<char>>,
    towel   : Vec<char>,
    visited : &mut HashMap<Vec<char>, usize>
) -> usize {

    if let Some(found) = visited.get(&towel) {
        return *found
    }
    if towel.is_empty() {
        return 1
    }
    //At each char in the towel
    //Check curr in options
    //  if so, slice string and recurse
    //  else increment char
    //if towel == "" return 1

    let mut total = 0;
    for c in 0 ..= towel.len() {

        let curr = &towel[0..c];

        if options.contains(&curr.to_vec()) {
            total += count_possibilities(&options, towel[c..].to_vec(), visited);
        }
    }
    visited.insert(towel, total);

    total
}

fn solve() -> (usize, usize) {

    let (combinations, options) = build_input();

    log!("Combinations");
    let mut p1 = 0;
    let mut p2 = 0;
    for combo in combinations {
        log!(&format!("Combo = {:?}", combo));

        let mut visited : HashMap<Vec<char>, usize> = HashMap::new();

        let res = count_possibilities(&options, combo, &mut visited);
    
        if res > 0 {
            p1 += 1;
        }
        p2 += res;
    }
    (p1, p2)
}
