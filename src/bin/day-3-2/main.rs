use std::io::Read;
use std::iter::Zip;
use std::slice::Iter;

fn main() {
    let mut input = String::new();
    std::io::stdin().read_to_string(&mut input).unwrap();

    //     let input = "00100
    // 11110
    // 10110
    // 10111
    // 10101
    // 01111
    // 00111
    // 11100
    // 10000
    // 11001
    // 00010
    // 01010";

    let mut lines = input.lines().collect::<Vec<&str>>();

    let line_len = lines[0].len();

    for idx in 0..line_len {
        let filtered_lines = filter_input(idx, &lines, '0', '1');

        lines = filtered_lines;

        if lines.len() == 1 {
            break;
        }
    }

    let oxy_num = lines[0];

    let mut lines = input.lines().collect::<Vec<&str>>();

    for idx in 0..line_len {
        let filtered_lines = filter_input(idx, &lines, '1', '0');

        lines = filtered_lines;

        if lines.len() == 1 {
            break;
        }
    }

    let co2_num = lines[0];

    println!("{:?}", oxy_num);
    println!("{:?}", co2_num);

    let oxy = isize::from_str_radix(oxy_num, 2).unwrap();
    let co2 = isize::from_str_radix(co2_num, 2).unwrap();

    println!("{}", oxy * co2);
}

fn count(input: &Vec<&str>) -> Vec<ZerosOnes> {
    let mut counts = vec![ZerosOnes { ones: 0, zeroes: 0 }];

    for &entry in input.iter() {
        for (idx, ch) in entry.chars().enumerate() {
            if idx >= counts.len() {
                counts.push(ZerosOnes { ones: 0, zeroes: 0 })
            }

            let mut zo = counts.get_mut(idx).unwrap();

            match ch {
                '0' => zo.zeroes += 1,
                '1' => zo.ones += 1,
                _ => {}
            }
        }
    }

    counts
}

fn filter_input<'filter>(
    column: usize,
    input: &Vec<&'filter str>,
    m1: char,
    m2: char,
) -> Vec<&'filter str> {
    let counts = count(&input);
    let count = &counts[column];

    let most = if count.zeroes > count.ones { m1 } else { m2 };
    let filtered_input = input
        .iter()
        .filter(|l| l.chars().nth(column).unwrap() == most)
        .cloned()
        .collect::<Vec<&str>>();

    return filtered_input;
}

#[derive(Debug)]
struct ZerosOnes {
    ones: i32,
    zeroes: i32,
}
