use std::collections::HashSet;

use utils::{read_file_to_int_grid, run_and_score};

fn main() {
    run_and_score("Part 1", || p1());//Result =  782. Ran For = 7240 us
    run_and_score("Part 2", || p2());//Result = 1694. Ran For = 4285 us
}


#[derive(Debug, Clone, Copy, Hash, PartialEq, Eq)]
struct Position(i32, i32);

impl Position {
    fn evaluate_for(&self, grid: &Grid) -> i32 {
        grid[self.0 as usize][self.1 as usize]
    }
    fn out_of_bounds(&self, size : i32) -> bool {
        self.0 < 0 || self.0 >= size || self.1 < 0 || self.1 >= size
    }

    fn add(&self, other: Position) -> Self {

        Self (
            self.0 + other.0,
            self.1 + other.1,
        )

    }
}

type Grid = Vec<Vec<i32>>;
type Visited = HashSet<Position>;

fn build_data() -> (Grid, Vec<Position>) {

    let grid: Vec<Vec<i32>> = read_file_to_int_grid("day10.txt");

    let mut trailheads: Vec<Position> = Vec::new();

    for (rix, row) in grid.iter().enumerate() {
        for (cix, col) in row.iter().enumerate() {

            if *col != 0 {
                continue;
            }

            trailheads.push(Position(rix as i32, cix as i32));
        }
    }

    return (grid, trailheads)
}

fn p1() -> i32 {
    let (grid, trailheads )= build_data();

    let mut sum = 0;

    for trailhead in trailheads {
        let mut visited: Visited = Visited::new();

        sum += dfs(&grid, trailhead, &mut visited, false);
    }

    sum
}
fn p2() -> i32 {
    let (grid, trailheads )= build_data();

    let mut sum = 0;

    for trailhead in trailheads {
        let mut visited: Visited = Visited::new();

        sum += dfs(&grid, trailhead, &mut visited, true);
    }

    sum
}

const MOVES : [Position; 4] = [
    Position(-1, 0),
    Position(1, 0),
    Position(0, -1),
    Position(0, 1),
];

fn dfs(grid: &Grid, position: Position, visited: &mut Visited, all_paths: bool) -> i32 {

    let mut sum = 0;

    let curr_height = position.evaluate_for(grid);

    if curr_height == 9 {
        return 1
    }

    for mv in MOVES {

        let new_pos = position.add(mv);

        if new_pos.out_of_bounds(grid.len() as i32) {
            continue;
        }
        if !all_paths && visited.contains(&new_pos) {
            continue;
        }
        if new_pos.evaluate_for(grid) - curr_height != 1 {
            continue;
        }
        visited.insert(new_pos);
        sum += dfs(grid, new_pos, visited, all_paths);
    }
    sum
}