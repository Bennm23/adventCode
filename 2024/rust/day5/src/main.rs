use std::collections::{HashMap, HashSet};

use utils::{log, logs, read_file_to_vec, run_and_score};


#[derive(Debug, PartialEq, Eq, Clone, Copy)]
struct Pairing(i32, i32);

fn build_data() -> (Vec<Vec<i32>>, HashMap<i32, HashSet<i32>>) {
    let lines = read_file_to_vec("day5.txt");

    let mut rules : Vec<Pairing> = Vec::new();
    let mut updates : Vec<Vec<i32>> = Vec::new();

    for line in &lines {
        logs!(line);
    
        if line.contains("|") {
            
            let split : Vec<i32> = line.split("|").map(|i| i.parse::<i32>().expect("Rule Parse Bad")).collect();
            rules.push(Pairing(split[0], split[1]));

        } else if line.contains(",") {
            let split : Vec<i32> = line.split(",").map(|i| i.parse::<i32>().expect("Update Parse Bad")).collect();
            updates.push(split);
        }
    }

    let mut ruleset : HashMap<i32, HashSet<i32>> = HashMap::new();

    for rule in &rules {
        log!(&format!("Rule = {:?}", rule));
        ruleset.entry(rule.0).or_insert_with(HashSet::new).insert(rule.1);
    }

    (updates, ruleset)
}

fn update_is_valid(update : &Vec<i32>, ruleset : &HashMap<i32, HashSet<i32>>) -> bool {

    let mut prev : Vec<i32> = Vec::new();

    for curr in update {

        if prev.is_empty() {
            prev.push(*curr);
            continue;
        }

        let opt = ruleset.get(curr);
        if opt.is_none() {
            continue;
        }
        let rule = opt.unwrap();

        //For each previous value
        //if the current values ruleset contains a previous, this is invalid
        for p in &prev {

            if rule.contains(p) {
                log!(&format!("Found to be invalid due to {curr} with prev = {:?}", prev));
                return false;
            }
        }
        prev.push(*curr);
    }
    true
}


fn p1() -> i32 {
    
    let (updates, ruleset) = build_data();

    updates.iter()
        .filter(|u| update_is_valid(&u, &ruleset))
        .filter_map(|u| Some(u[u.len() / 2]))
        .sum()
}

fn p2() -> i32 {
    
    let (updates, ruleset) = build_data();


    let bad_updates : Vec<&Vec<i32>> = updates.iter()
                .filter(|x| !update_is_valid(x, &ruleset))
                .collect();

    let mut sum = 0;

    for bad in bad_updates {

        let mut clone = bad.clone();
        log!("--------------");
        log!(&format!("Evaluating {:?}", clone));


        for i in (0 .. bad.len()).rev() {

            while try_move(&mut clone, i, &ruleset) {
                
                log!(&format!("Swapped {:?}", clone));
            }
        }

        sum += clone[clone.len() / 2];
    }
    sum
}

/**
 * Try to swap the vector indices. This is assuming we are walking backwards from the end
 * given our ruleset.
 */
fn try_move(v : &mut Vec<i32>, index : usize, ruleset : &HashMap<i32, HashSet<i32>>) -> bool {

    let val = v[index];

    //walking forward from index 0
    // if val needs to come before v[i]
    //  swap index and i then return true
    let opt = ruleset.get(&val);
    if opt.is_none() {
        log!(&format!("{val} Is Not in Ruleset"));
        return false;
    }
    let rule = opt.unwrap();

    for i in 0 .. index {
        
        if rule.contains(&v[i]) {
        
            v.swap(i, index);
            return true;
        }
    }
    log!(&format!("{val} Does not need to be moved"));
    false
}

fn main() {
    run_and_score("Part 1", || { p1() }); //6041, 807 us
    run_and_score("Part 2", || { p2() }); //4884, 2766 us
}
