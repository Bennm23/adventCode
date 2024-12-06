use std::{collections::HashSet, fs};

use utils::{get_advent_path, log, run_and_score};

#[derive(Debug, PartialEq, Eq, Clone, Copy, Hash)]
struct Position(usize, usize);

impl Position {

    fn in_bounds(&self, size : usize) -> bool {

        self.0 < size && self.1 < size
    }
}

#[derive(Debug, PartialEq, Eq, Clone, Copy)]
struct Move(i32, i32);

impl Move {
    fn apply(&self, pos : Position) -> Position {
        let mut res : Position = pos;

        if self.0 < 0 {
            res.0 = res.0.wrapping_sub(self.0.abs() as usize);
        } else {
            res.0 = res.0.wrapping_add(self.0.abs() as usize);
        }

        if self.1 < 0 {
            res.1 = res.1.wrapping_sub(self.1.abs() as usize);
        } else {
            res.1 = res.1.wrapping_add(self.1.abs() as usize);
        }

        res
    }
    fn turn(&mut self) {
        //up
        if self.0 == -1 {
            self.0 = 0;
            self.1 = 1;
        } else if self.0 == 1 {
            self.0 = 0;
            self.1 = -1;
        } else if self.1 == -1 {
            self.0 = -1;
            self.1 = 0;
        } else if self.1 == 1 {
            self.0 = 1;
            self.1 = 0;
        }
    }
}

#[derive(Debug, PartialEq, Eq, Clone, Copy)]
struct Guard {
    position : Position,
    movement : Move,
}
impl Guard {
    fn at_pos(&self, pos : Position) -> bool {
        self.position.0 == pos.0 && self.position.1 == pos.1
    }

    fn visualize(&self) -> &str {
        match self.movement {
            Move(-1,0) => {return "U "},
            Move(1,0) => {return "D "},
            Move(0,-1) => {return "L "},
            Move(0,1) => {return "R "},
            _ => { return "?"}
        }
    }

    fn make_move(&mut self, grid : &Vec<Vec<char>>) -> bool {

        let destination = self.movement.apply(self.position);

        if !destination.in_bounds(grid.len()) {
            return false;
        }

        if grid[destination.0][destination.1] == '#' {
            self.movement.turn();
        } else {
            self.position = destination;
        }

        true
    }
}

fn p1() -> i32 {
    let contents = fs::read_to_string(get_advent_path("day6.txt"))
        .expect("Could Not Parse File");

    let grid: Vec<Vec<char>> = contents.split("\n")
        .map(|s : &str| s.chars().collect())
        .filter(|f : &Vec<char>| !f.is_empty())
        .collect();

    let size = grid.len();

    let mut guard = Guard{position: Position(0,0), movement: Move(-1,0)};

    'outer: for i in 0 .. size {
        for j in 0 .. size {

            if grid[i][j] == '^' {

                guard.position = Position(i, j);
                break 'outer;
            }
        }
    }

    print_state(guard, &grid);

    let mut visited : HashSet<Position> = HashSet::new();

    while guard.make_move(&grid) {

        visited.insert(guard.position);
    }

    visited.len() as i32
}

fn print_state(guard : Guard, grid : &Vec<Vec<char>>) {

    for i in 0 .. grid.len() {
        for j in 0 .. grid.len() {
            
            if guard.at_pos(Position(i, j)) {
                log!(guard.visualize(), false);
            } else {
                log!(&format!("{} ", grid[i][j]), false);
            }
        }
        log!()
    }

}

fn p2() -> i32 {
    let contents = fs::read_to_string(get_advent_path("day6.txt"))
        .expect("Could Not Parse File");

    let initial_grid: Vec<Vec<char>> = contents.split("\n")
        .map(|s : &str| s.chars().collect())
        .filter(|f : &Vec<char>| !f.is_empty())
        .collect();

    let size = initial_grid.len();


    let mut initial_guard = Guard{position: Position(0,0), movement: Move(-1,0)};

    'outer: for i in 0 .. size {
        for j in 0 .. size {

            if initial_grid[i][j] == '^' {

                initial_guard.position = Position(i, j);
                break 'outer;
            }
        }
    }

    let mut sum = 0;
    'outer: for i in 0 .. size {
        for j in 0 .. size {

            if initial_grid[i][j] == '#'  || initial_grid[i][j] == '^' {
                continue;
            }

            let mut grid = initial_grid.clone();
            grid[i][j] = '#';

            sum += solve_grid2(&mut initial_guard.clone(), &grid);
        }
    }

    sum

}

fn solve_grid2(guard : &mut Guard, grid : &Vec<Vec<char>>) -> i32 {

    let mut visited : HashSet<Position> = HashSet::new();

    let mut last_pos = guard.position;

    let mut loop_counter = 0;

    while guard.make_move(&grid) {

        if guard.position != last_pos && visited.contains(&guard.position) {
            loop_counter += 1;
        }
        if loop_counter > visited.len() {
            return 1;
        }
        visited.insert(guard.position);
        last_pos = guard.position;
    }

    0
}

fn main() {

    run_and_score("Part 1", || { p1() });//4515, 4240 us
    run_and_score("Part 2", || { p2() });//1309, 4367754 us
}
