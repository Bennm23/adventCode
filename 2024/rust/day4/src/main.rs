
use utils::{log, read_file_to_grid};

extern crate utils;

#[derive(Debug, Clone, Copy)]
struct Point(usize, usize);

fn main() {
    let grid : Vec<Vec<char>> = read_file_to_grid("day4.txt");

    let mut xs : Vec<Point> = Vec::new();

    for (rx, r) in grid.iter().enumerate() {
        for (cx, c) in r.iter().enumerate() {
            log!(c.to_string().as_str(), false);

            if *c == 'X' {
                xs.push(Point(rx, cx));
            }
        }
        log!();
    }

    log!("-----------");
    let mut p1 = 0;
    for x in xs {
        log!(format!("X at {:?}", x).as_str());
    
        p1 += eval_x(x, &grid);
    }
    println!("P1 Total = {p1}");//2547

}

//fn check_dir(x : Point, grid : &Vec<Vec<char>>, mover : &dyn Fn(Point) -> Point) -> bool {
fn check_dir<F>(x : Point, grid : &Vec<Vec<char>>, mover : F) -> bool 
where 
    F: Fn(Point) -> Point
{
    let mut found = false;

    log!("------------");
    let mut new_x = x;
    let mut expected : char = 'M';
    for _ in 0 .. 3 {
        new_x = mover(new_x);
        log!(format!("Moved X = {:?}",new_x).as_str());
        
        if new_x.0 >= grid.len() || new_x.1 >= grid[0].len() {
            break;
        }

        let val : char = grid[new_x.0][new_x.1];
        log!(format!("VAL AT NEW X = {val}").as_str());

        if val == expected {
            
            if val == 'S' {
                found = true;
                log!("Found End");
                break;
            } else if val == 'A' {
                expected = 'S';
                log!("Found A");
            } else if val == 'M' {
                expected = 'A';
                log!("Found M");
            }
        } else {
            break;
        }
    }

    found
}

fn eval_x(x : Point, grid : &Vec<Vec<char>>) -> u32 {
    let mut count = 0;

    log!("================");
    log!(format!("Evaluating X = {:?}",x).as_str());

    //Right
    if check_dir(x, grid, |x: Point| {Point(x.0, x.1 + 1)}) {
        count += 1;
    }
    //Left
    if check_dir(x, grid, |x: Point| {Point(x.0, x.1.wrapping_sub(1))}) {
        count += 1;
    }
    //Up
    if check_dir(x, grid, |x: Point| {Point(x.0.wrapping_sub(1), x.1)}) {
        count += 1;
    }
    //Down
    if check_dir(x, grid, |x: Point| {Point(x.0 + 1, x.1)}) {
        count += 1;
    }

    //Up Right
    if check_dir(x, grid, |x: Point| {Point(x.0.wrapping_sub(1), x.1 + 1)}) {
        count += 1;
    }
    //Down Right
    if check_dir(x, grid, |x: Point| {Point(x.0 + 1, x.1 + 1)}) {
        count += 1;
    }
    //Up Left
    if check_dir(x, grid, |x: Point| {Point(x.0.wrapping_sub(1), x.1.wrapping_sub(1))}) {
        count += 1;
    }
    //Down Left
    if check_dir(x, grid, |x: Point| {Point(x.0 + 1, x.1.wrapping_sub(1))}) {
        count += 1;
    }

    count
}
