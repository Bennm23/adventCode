use utils::{log, read_file_to_vec, run_and_score};


struct Line {
    total : i64,
    opts  : Vec<i64>,
}

impl Line {

    fn new(file_line : String) -> Self {
        let split : Vec<&str> = file_line.split(": ").collect();

        log!(&format!("File Line = {}", file_line));
        log!(&format!("Split[0] = {}, Split[1] = {}", split[0], split[1]));
        let total: i64 = split[0].parse().expect("Parse Total Failed");

        let opts : Vec<i64> = split[1]
                .split(" ")
                .map(|s| {s.parse::<i64>().expect("Failed to parse opt")})
                .collect();

        Self {
            total,
            opts
        }
    }

    fn degug(&self) {
        log!(&format!("Sum = {}: Opts = {:?}", self.total, self.opts))
    }

}

//Can't define new lambda each recurse, get overflow. Also fn is
//not the same as Fn. fn is a type, Fn is a trait. 
const ADD : fn(i64, i64) -> i64 = |a, b| {a + b};
const MUL : fn(i64, i64) -> i64 = |a, b| {a * b};
const OR  : fn(i64, i64) -> i64 = |a, b| {
    let mut string_res : String = a.to_string();
    string_res.push_str(b.to_string().as_str());
    string_res.parse::<i64>().expect("Failed to Merge L/R")
};

#[derive(Debug)]
enum Func {
    ADDED,
    MULLED,
    ORD,
}

fn solve<F: Fn(i64,i64) -> i64>(
    p2 : bool, total : i64,
    left : i64, right_index : usize,
    opts : &Vec<i64>, operator : F,
    debugging : &mut Vec<Func>
) -> bool {
    
    if right_index == opts.len() {
        return false;
    }
    let res = operator(left, opts[right_index]);

    if res > total {
        return false;
    }
    if res == total && right_index == (opts.len() - 1) {
        return true;
    }

    if p2 {
        let add = solve(p2, total, res, right_index + 1, opts, ADD, debugging);
        if add {
            debugging.insert(0, Func::ADDED);
            return true;
        }
        let mul = solve(p2, total, res, right_index + 1, opts, MUL, debugging);
        if mul {
            debugging.insert(0, Func::MULLED);
            return true;
        }
        let or = solve(p2, total, res, right_index + 1, opts, OR, debugging);
        if or {
            debugging.insert(0, Func::ORD);
            return true;
        }
        return false;
    }
    return solve(p2, total, res, right_index + 1, opts, ADD, debugging) ||
            solve(p2, total, res, right_index + 1, opts, MUL, debugging);

}
fn p1() -> i64 {

    let data = build_data();
    let mut sum : i64 = 0;

    for entry in &data {
        entry.degug();
    
        let mut debugging : Vec<Func> = Vec::new();
        let res = solve(false, entry.total, entry.opts[0], 0, &entry.opts, |a,_b| { a }, &mut debugging);

        if res {
            sum += entry.total;
        }
    }
    sum
}

fn p2() -> i64 {

    let data = build_data();
    let mut sum : i64 = 0;

    log!("-----------\nSolved");
    for entry in data {
        let mut debugging : Vec<Func> = Vec::new();
        let res = solve(true, entry.total, entry.opts[0], 0, &entry.opts, |a,_b| { a }, &mut debugging);

        if res {
            sum += entry.total;

            log!("----------");
            entry.degug();
            log!("Steps: ", false);
            for step in debugging {
                log!(&format!("{:?} ", step), false);
            }
            log!();
        }
    }
    sum
}

fn build_data() -> Vec<Line> {
    let mut res: Vec<Line> = Vec::new();
    for line in  read_file_to_vec("day7.txt") {
        res.push(Line::new(line));
    }
    res
}

fn main() {
    run_and_score("Part 1", || {p1()});//Total: 3119088655389.   Time: 2009 us
    run_and_score("Part 2", || {p2()});//Total: 264184041398847. Time: 333088 us
}