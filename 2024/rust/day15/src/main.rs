use utils::{read_file_to_groups, run_and_score, Pair};

fn main() {
    run_and_score("Part 1", || p1());//Result = 1451928. Ran For = 620 us
    run_and_score("Part 2", || p2());//Result = 1462788. Ran For = 512 us
}

type Grid = Vec<Vec<char>>;

fn get_initial_data(expand : bool) -> (Vec<Pair>, Grid, Pair) {
    let groups = read_file_to_groups("day15.txt", "");
    let grid_vec = groups[0].clone();
    let moves_vec = groups[1].clone();

    let mut grid : Grid = Vec::new();

    for row in grid_vec {
        grid.push(row.chars().collect());
    }

    let mut position = Pair(0, 0);

    if expand {
        let mut new_grid : Grid = Vec::new();
        for row in &grid {
            let mut row_vec : Vec<char> = Vec::new();
            for col in row {
                if *col == 'O' {
                    row_vec.push('[');
                    row_vec.push(']');
                }
                if *col == '.' {
                    row_vec.push('.');
                    row_vec.push('.');
                }
                if *col == '#' {
                    row_vec.push('#');
                    row_vec.push('#');
                }

                if *col == '@' {
                    row_vec.push('@');
                    row_vec.push('.');
                }
            }
            new_grid.push(row_vec);
        }
        grid = new_grid;
    }
    for (rix, row) in grid.iter().enumerate() {
        for (cix, col) in row.iter().enumerate() {
            if *col == '@' {
                position = Pair(rix as i32, cix as i32);
            }
        }
    }

    let mut moves : Vec<Pair> = Vec::new();

    for row in moves_vec {
    
        row.chars().for_each(|c: char| {
            if c == '>' {
                moves.push(Pair(0, 1));
            } else if c == '<' {
                moves.push(Pair(0, -1));
            } else if c == '^' {
                moves.push(Pair(-1, 0));
            } else {
                moves.push(Pair(1, 0));
            }
        });
    }
    (moves, grid, position)

}

fn p1() -> usize {

    let (moves, mut grid, mut position) = get_initial_data(false);

    for mv in &moves {
        if try_move(position, &mut grid, mv, false) {
            try_move(position, &mut grid, mv, true);
            position = position.add(mv);
        }
    }

    score_grid('O', &grid)
}

fn p2() -> usize {
    let (moves, mut grid, mut position) = get_initial_data(true);

    for mv in &moves {
        if try_move(position, &mut grid, mv, false) {
            try_move(position, &mut grid, mv, true);
            position = position.add(mv);
        }
    }
    score_grid('[', &grid)
}

fn score_grid(boulder : char, grid : &Grid) -> usize {
    let mut sum = 0;

    for (rix, row) in grid.iter().enumerate() {
        for (cix, col) in row.iter().enumerate() {
            if *col == boulder {
                sum += 100*rix + cix
            }
        }
    }
    sum
}

fn try_move(position : Pair, grid : &mut Grid, mv : &Pair, apply : bool) -> bool {
    let new_position = position.add(mv);
    
    //No need to worry about bounds because the border is blocked
    match grid[new_position.ia()][new_position.ib()] {
        '#' => {
            return false;
        }
        '.' => {
            if apply {
                grid[new_position.ia()][new_position.ib()] = grid[position.ia()][position.ib()];
                grid[position.ia()][position.ib()] = '.';
            }
            return true;
        }
        ']' | '[' => {
            //Moving up/down
            if mv.0 != 0 {
                let mut other_pos = new_position.add(&Pair(0, -1));
                if grid[new_position.ia()][new_position.ib()] == '[' {
                    other_pos = new_position.add(&Pair(0, 1));
                }

                if try_move(new_position, grid, mv, apply) && 
                    try_move(other_pos, grid, mv, apply) {

                    if apply {
                        grid[new_position.ia()][new_position.ib()] = grid[position.ia()][position.ib()];
                        grid[position.ia()][position.ib()] = '.';
                    }
                    return true;
                }
                return false;
            } else {
                //We can treat left/right as the same
                if try_move(new_position, grid, mv, apply) {
                    if apply {
                        grid[new_position.ia()][new_position.ib()] = grid[position.ia()][position.ib()];
                        grid[position.ia()][position.ib()] = '.';
                    }
                    return true;
                }
                return false;
            }
        }
        'O' => {
            //P1
            if try_move(new_position, grid, mv, apply) {
                if apply {
                    grid[new_position.ia()][new_position.ib()] = grid[position.ia()][position.ib()];
                    grid[position.ia()][position.ib()] = '.';
                }
                return true;
            }
            return false;
        }

        _ => { unreachable!() }
    }
}