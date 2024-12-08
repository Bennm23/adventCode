use std::collections::{HashMap, HashSet};

use utils::{log, read_file_to_grid, run_and_score};

type AntennaMap = HashMap<char, Vec<Position>>;
type Grid = Vec<Vec<char>>;

#[derive(Debug, Hash, Clone, Copy, PartialEq, Eq)]
struct Position(i32, i32);

const GRID_LEN : usize = 50;

impl Position {

    fn in_bounds(&self) -> bool {
        if self.0 < 0 || self.1 < 0 {
            return false;
        }
        return (self.0 as usize) < GRID_LEN && (self.1 as usize) < GRID_LEN;
    }
}

fn p1() -> i32 {
    let (grid, antenna_mapping) = build_state();

    let antinodes = find_antinodes(&antenna_mapping,true);

    print_state(&grid, &antinodes);

    antinodes.len() as i32
}

fn p2() -> i32 {
    let (grid, antenna_mapping) = build_state();

    let antinodes = find_antinodes(&antenna_mapping, false);

    print_state(&grid, &antinodes);

    antinodes.len() as i32
}

fn build_state() -> (Grid, AntennaMap) {
    let grid : Grid = read_file_to_grid("day8.txt");

    let mut antenna_mapping : AntennaMap = HashMap::new();

    for (rix, row) in grid.iter().enumerate() {
        for (cix, col) in row.iter().enumerate() {
            
            if *col == '.' {
                continue;
            }
            antenna_mapping.entry(*col)
                .or_insert_with(Vec::new)
                .push(Position(rix as i32, cix as i32));
        }
    }
    return (grid, antenna_mapping)
}
fn check_antennas_fixed(a1 : Position, a2 : Position, antinodes : &mut HashSet<Position>) {
    let row_diff = a1.0 - a2.0;
    let col_diff = a1.1 - a2.1;

    log!(&format!("Row Diff = {row_diff}, Col Diff = {col_diff}"));

    let antinode = Position(
        a1.0 + row_diff,
        a1.1 + col_diff,
    );

    if antinode.in_bounds() {
        antinodes.insert(antinode);
        
    }
}

fn check_antennas(a1 : Position, a2 : Position, fixed_distance : bool, antinodes : &mut HashSet<Position>) {
    let row_diff = a1.0 - a2.0;
    let col_diff = a1.1 - a2.1;

    log!(&format!("Row Diff = {row_diff}, Col Diff = {col_diff}"));

    let mut antinode = Position(
        a1.0 + row_diff,
        a1.1 + col_diff,
    );

    if !fixed_distance {
        antinodes.insert(a1);
        
    }

    while antinode.in_bounds() {
        log!(&format!("Inserting Antinode at {:?}", antinode));
        antinodes.insert(antinode);

        if fixed_distance {
            break;
        }

        antinode = Position(
            antinode.0 + row_diff,
            antinode.1 + col_diff,
        );
    }
}

fn find_antinodes(antenna_mapping : &AntennaMap, fixed_distance : bool) -> HashSet<Position> {
    let mut antinodes : HashSet<Position> = HashSet::new();

    for (frequency, antennas) in antenna_mapping {

        log!("----------------");
        log!(&format!("Frequencey = {}", frequency));
        for i in 0 .. antennas.len() {

            for j in i + 1 .. antennas.len() {
                check_antennas(antennas[i], antennas[j], fixed_distance, &mut antinodes);
                check_antennas(antennas[j], antennas[i], fixed_distance, &mut antinodes);
            }
        }
    }
    return antinodes;
}

fn print_state(grid : &Grid, antinodes : &HashSet<Position>) {

    for (rix, row) in grid.iter().enumerate() {

        for (cix, col) in row.iter().enumerate() {
            
            if antinodes.contains(&Position(rix as i32, cix as i32)) && *col != '.' {
                log!(&format!("#{} ", *col), false);
            }
            else if antinodes.contains(&Position(rix as i32, cix as i32)) {
                log!("#  ", false);
            } else {
                log!(&format!("{}  ", *col), false);
            }
        }
        log!()
    }
}

fn main() {
    run_and_score("Part 1", || p1());//Total  308, 2925 us
    run_and_score("Part 2", || p2());//Total 1147, 4262 us
}