use utils::read_file_to_vec;



extern crate utils;
fn main() {
    println!("Hello, world!");

    let f = read_file_to_vec("day2.txt");

    let mut sum = 0;

    let mut invalid : Vec<Vec<i32>> = Vec::new();

    for line in f {
        println!("{line}");

        let split : Vec<i32> = line.split(" ").map(|x| x.parse::<i32>().unwrap()).collect();

        println!("{:?}", split);

        if validate(&split) {
            sum += 1;
        } else {
            invalid.push(split.clone());
        }

    }

    println!("Part 1 = {sum}");

    let mut p2 = 0;
    println!("Invalid Count = {:?}", invalid.len());

    for i in invalid {
        for j in 0 .. i.len() {

            let mut new = i.clone();
            new.remove(j);

            if validate(&new) {
                p2 += 1;
                break;
            }
        }
    }

    println!("Valid invalids = {p2}");
    println!("New Sum = {:?}", p2 + sum);
}

// fn validate(report : &Vec<i32>) -> bool {

//     let increasing = report[0] < report[1];

//     let mut diff = 0;
//     for i in 0 .. report.len() - 1 {

//         diff = report[i] - report[i + 1];

//         if diff.abs() < 1 || diff.abs() > 3 {
//             let mut new = report.clone();
//             new.remove(i);
//             return validate_again(new);
//         }

//         if increasing && diff > 0 {
//             let mut new = report.clone();
//             new.remove(i);
//             return validate_again(new);
//         } else if !increasing && diff < 0 {
//             let mut new = report.clone();
//             new.remove(i);
//             return validate_again(new);
//         }
//     }
//     true
// }

fn validate(report : &Vec<i32>) -> bool {

    let increasing = report[0] < report[1];

    let mut diff = 0;
    for i in 0 .. report.len() - 1 {

        diff = report[i] - report[i + 1];

        if diff.abs() < 1 || diff.abs() > 3 {
            return false;
        }

        if increasing && diff > 0 {
            return false;
        } else if !increasing && diff < 0 {
            return false;
        }
    }
    true
}