use atoi::atoi;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let input_file = File::open("day-1-1-input.txt").unwrap();
    let reader = BufReader::new(input_file);

    let mut depths: Vec<i32> = Vec::new();

    for line in reader.lines() {
        let current = atoi::<i32>(line.unwrap().as_bytes()).unwrap();
        depths.push(current);
    }

    let length = depths.len();

    let mut aggregate: Vec<i32> = Vec::new();

    for x in 0..length {
        let end = if x + 3 > length { length } else { x + 3 };

        let slice = &depths[x..end];
        aggregate.push(slice.iter().sum());
    }

    let mut previous: Option<i32> = None;
    let mut increases = 0;

    for x in aggregate {
        match previous {
            Some(p) => {
                if x > p {
                    increases += 1;
                }

                previous = Some(x);
            }
            None => previous = Some(x),
        }
    }

    println!("{}", increases);
}
