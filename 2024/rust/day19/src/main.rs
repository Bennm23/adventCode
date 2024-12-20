use std::collections::HashMap;

use utils::{log, read_file_to_vec, run_and_score_both};

fn main() {
//  Part 1: Result = 306.
//  Part 2: Result = 604622004681855.
//  Took 209471 us
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

type VisitedCombos = HashMap<Vec<char>, usize>;

fn count_possibilities(
    options     : &Vec<Vec<char>>,
    combination : &[char],
    visited     : &mut VisitedCombos
) -> usize {

    if let Some(found) = visited.get(combination) {
        return *found
    }
    if combination.is_empty() {
        return 1
    }
    //At each char in the towel
    //Check curr in options
    //  if so, slice string and recurse
    //  else increment char
    //if towel == "" return 1

    let mut total = 0;
    for c in 0..=combination.len() {

        let curr = &combination[0..c];

        if options.contains(&curr.to_vec()) {
            total += count_possibilities(&options, &combination[c..], visited);
        }
    }
    visited.insert(combination.to_vec(), total);

    total
}

fn solve() -> (usize, usize) {

    let (combinations, options) = build_input();

    log!("Combinations");
    let mut p1 = 0;
    let mut p2 = 0;
    for combination in combinations {
        log!(&format!("Combination = {:?}", combination));

        let mut visited : VisitedCombos = HashMap::new();

        let res = count_possibilities(&options, &combination, &mut visited);
    
        if res > 0 {
            p1 += 1;
        }
        p2 += res;
    }
    (p1, p2)
}
