use atoi::atoi;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let input_file = File::open("day-1-1-input.txt")?;
    let reader = BufReader::new(input_file);

    let mut previous: Option<i32> = None;
    let mut increases = 0;

    for line in reader.lines() {
        match line {
            Ok(v) => {
                let current = atoi::<i32>(v.as_bytes()).unwrap();

                match previous {
                    Some(p) => {
                        if current > p {
                            increases += 1;
                        }

                        previous = Some(current);
                    }
                    None => previous = Some(current),
                }
            }

            Err(err) => {
                return Err(err.into());
            }
        }
    }

    println!("{}", increases);

    Ok(())
}
