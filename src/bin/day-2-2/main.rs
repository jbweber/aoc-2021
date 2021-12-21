use std::io::Read;

fn main() {
    let mut input = String::new();
    std::io::stdin().read_to_string(&mut input).unwrap();

    let mut aim = 0;
    let mut position = 0;
    let mut depth = 0;

    for line in input.lines() {
        let parts: Vec<&str> = line.split(' ').collect();

        let command = parts[0];
        let val: i32 = parts[1].parse().unwrap();

        match command {
            "forward" => {
                position = position + val;
                depth = depth + (aim * val);
            }

            "down" => {
                aim = aim + val;
            }

            "up" => {
                aim = aim - val;
            }

            _ => {
                println!("found bad direction {}", parts[0]);
            }
        }
    }

    println!("{}", depth * position);
}
