
use utils::{read_file_to_vec, run_and_score, Pair};

fn main() {
    run_and_score("Part 1", || p1());//Result = 1323.   Ran For = 129620  us
    run_and_score("Part 2", || p2());//Result = 983905. Ran For = 3812521 us
}

type Grid = Vec<Vec<bool>>;
type Path = Vec<Pair>;
const MOVES : [Pair; 4] = [
    Pair(-1, 0),
    Pair(1, 0),
    Pair(0, -1),
    Pair(0, 1),
];

fn build_grid() -> (Grid, Pair, Pair) {

    let mut start_pos = Pair(0, 0);
    let mut stop_pos = Pair(0, 0);
    let lines = read_file_to_vec("day20.txt");
    let mut char_grid: Vec<Vec<char>> = Vec::new();
    for line in lines {
        char_grid.push(line.chars().collect());
    }
    let mut grid: Grid = Vec::new();

    for (rix, row) in char_grid.iter().enumerate() {
        let mut row_vec = Vec::new();
        for (cix, col) in row.iter().enumerate() {
            if *col == '#' {
                row_vec.push(true);
            } else {
                row_vec.push(false);
            }
            if *col == 'S' {
                start_pos = Pair(rix as i32, cix as i32);
            } else if *col == 'E' {
                stop_pos = Pair(rix as i32, cix as i32);
            }
        }
        grid.push(row_vec);
    }

    (grid, start_pos, stop_pos)
}

fn find_orig_path(node: Pair, grid : &Grid, goal: Pair, path : &mut Path) {
    if node == goal {
        return;
    }
    for mv in MOVES {
        let res = node.add(&mv);
        if res.out_of_bounds(grid.len()) {
            continue;
        }
        if path.contains(&res) {
            continue;
        }
        //Means Rock
        if grid[res.ia()][res.ib()] {
            continue;
        }
        path.push(res);
        find_orig_path(res, grid, goal, path);
    }
}

fn p1() -> i32 {
    let (grid, start, goal) = build_grid();
    let mut orig_path = vec![start];
    find_orig_path(start, &grid, goal, &mut orig_path);

    let mut sum = 0;

    for (index, point) in orig_path.iter().enumerate() {
        sum += check_index(index, point, 2, &grid, orig_path.len() - 1, &orig_path);
    }
    sum
}

fn p2() -> i32 {
    let (grid, start, goal) = build_grid();
    let mut orig_path = vec![start];
    find_orig_path(start, &grid, goal, &mut orig_path);

    let mut sum = 0;

    for (index, point) in orig_path.iter().enumerate() {
        sum += check_index(index, point, 20, &grid, orig_path.len() - 1, &orig_path);
    }
    sum
}

fn check_index(
    jump_index : usize, point : &Pair,
    max_cheats : i32, grid : &Grid,
    orig_best : usize, orig_path : &Vec<Pair>
) -> i32 {

    let mut sum = 0;
    //for each index, try all points in a max_cheat grid around the curr index
    for rix in (point.0 - max_cheats)..=(point.0 + max_cheats) {
        for cix in (point.1 - max_cheats)..=(point.1 + max_cheats) {

            let new_pos = Pair(rix, cix);
        
            let steps_taken = (rix - point.0).abs() + (cix - point.1).abs();

            //If result is OB, or we have cheated too much or a # continue (not needed just makes it faste)
            if new_pos.out_of_bounds(grid.len()) || steps_taken > max_cheats || grid[new_pos.ia()][new_pos.ib()] {
                continue;
            }

            //Else this is a valid cross attempt
            let mut landing_index = jump_index;

            for (ix, p) in orig_path.iter().enumerate() {
                if *p == new_pos {
                    landing_index = ix;
                    break;
                }
            }
            //New score is the jump index(already 0 indexed), + jump length + distance from land to end
            let new_score = jump_index + steps_taken as usize + (orig_best - landing_index);

            let improvement = orig_best as i32 - new_score as i32;
            // if improvement > 0 {
            //     println!("Found Shorter Path, Len = {}", (no_ghost_best - new_score));
            // }
            if improvement >= 100 {
                sum += 1;
            }
        }
    }
    sum
}