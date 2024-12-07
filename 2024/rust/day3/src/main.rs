use lazy_static::lazy_static;
use regex::Regex;
use utils::{log, logs, read_file_to_vec, remove_between_or_after, run_and_score};


#[derive(Debug)]
struct Pair(i32, i32);

impl Pair {
    fn new(str : &str) -> Self {
        lazy_static! {
            static ref RE_INTS: Regex = Regex::new(r"\d+").unwrap();
        }

        let vals: Vec<i32> = RE_INTS.find_iter(str)
            .filter_map(|digits| digits.as_str().parse().ok())
            .collect();

        Self(vals[0], vals[1])
    }
}

fn count_matches(superstring : String) -> i32 {
    lazy_static! {
        static ref RE_PAIRS: Regex = Regex::new(r"mul\(\d{1,3},\d{1,3}\)").unwrap();
    }
    let matches : Vec<Pair> = RE_PAIRS.find_iter(&superstring)
        .map(|s| Pair::new(s.as_str()))
        .collect();

    let mut sum = 0;
    for m in matches {

        log!(&format!("Match = {:?}", m));
        sum += m.0 * m.1;
    }
    sum
}

fn build_superstring() -> String {
    let chunks = read_file_to_vec("day3.txt");

    let mut superstring: String = String::new();

    for chunk in chunks {
        logs!(&chunk);
        superstring.push_str(&chunk);
    }

    log!("--------------");
    logs!(&superstring);

    superstring
}

fn p1() -> i32 {

    let superstring: String = build_superstring();
    count_matches(superstring)
}

fn p2() -> i32 {
    let mut superstring: String = build_superstring();

    superstring = remove_between_or_after(superstring.as_str(), "don't()", "do()");

    log!("----------------");
    log!(&superstring);

    count_matches(superstring)
}

fn main() {
    run_and_score("Part 1", || p1());//Total: 180233229. Run time 5969 us
    run_and_score("Part 2", || p2());//Total: 95411583.  Run time 3557 us
}
