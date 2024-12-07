use std::{collections::{HashSet}, fs, hash::Hash};

use utils::{get_advent_path, log, run_and_score};

#[derive(Debug, PartialEq, Eq, Clone, Copy, Hash)]
struct Position(usize, usize);

impl Position {

    fn in_bounds(&self, size : usize) -> bool {

        self.0 < size && self.1 < size
    }
}

#[derive(Debug, PartialEq, Eq, Clone, Copy, Hash)]
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
    fn int_val(&self) -> usize {
        if self.0 == -1 {
            return 0;
        } else if self.0 == 1 {
            return 1;
        } else if self.1 == -1 {
            return 2;
        } else {
            return 3;
        }
    }
}

#[derive(Debug, PartialEq, Eq, Clone, Copy, Hash)]
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

fn get_guard_path(guard : &mut Guard, grid : &Vec<Vec<char>>) -> HashSet<Position> {

    let mut path : HashSet<Position> = HashSet::new();

    while guard.make_move(&grid) {

        path.insert(guard.position);
    }

    path
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

    let visited = get_guard_path(&mut guard, &grid);

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
    let path = get_guard_path(&mut initial_guard.clone(), &initial_grid.clone());

    let mut sum = 0;
    let path_len = path.len();
    
    for pos in path {
        if initial_grid[pos.0][pos.1] == '#'  || initial_grid[pos.0][pos.1] == '^' {
            continue;
        }

        let mut grid = initial_grid.clone();
        grid[pos.0][pos.1] = '#';

        sum += solve_grid2d(&mut initial_guard.clone(), &grid, path_len);
        // sum += solve_grid_3d(&mut initial_guard.clone(), &grid);
    }
    sum
}
const SIZE : usize = 130;

fn solve_grid2d(guard : &mut Guard, grid : &Vec<Vec<char>>, path_len : usize) -> i32 {

    let mut visited : [[bool; SIZE]; SIZE] = [[false; SIZE]; SIZE];

    let mut last_pos = guard.position;

    let mut loop_counter = 0;

    while guard.make_move(&grid) {

        if guard.position != last_pos && visited[guard.position.0][guard.position.1] {
            loop_counter += 1;
        }
        if loop_counter > path_len {
            return 1;
        }
        visited[guard.position.0][guard.position.1] = true;

        last_pos = guard.position;
    }
    0
}

fn solve_grid_3d(guard : &mut Guard, grid : &Vec<Vec<char>>) -> i32 {

    let mut visited : [[[bool; 4]; SIZE]; SIZE] = [[[false; 4]; SIZE]; SIZE];

    let mut last_pos = guard.position;

    while guard.make_move(&grid) {

        if guard.position != last_pos && visited[guard.position.0][guard.position.1][guard.movement.int_val()] {
            return 1;
        }
        visited[guard.position.0][guard.position.1][guard.movement.int_val()] = true;

        last_pos = guard.position;
    }
    0
}

fn main() {
    run_and_score("Part 1", || { p1() });//4515, 4240 us
    //1309, naive 4367754 us. 
    //HashSet path only 764746 us.
    //2d array path only 67579 us 59443 us after caching
    //3d array path only 78024 us 54058 us after caching
    run_and_score("Part 2", || { p2() });
}
