use std::io::Read;

fn main() {
    let mut input = String::new();
    std::io::stdin().read_to_string(&mut input).unwrap();

    let mut zeroes = vec![0];
    let mut ones = vec![0];

    for line in input.lines() {
        for (x, y) in line.chars().enumerate() {
            if x >= zeroes.len() {
                zeroes.push(0);
            };

            if x >= ones.len() {
                ones.push(0);
            };

            let zero = zeroes.get(x).unwrap();
            let one = ones.get(x).unwrap();

            match y {
                '0' => zeroes[x] = zero + 1,
                '1' => ones[x] = one + 1,
                _ => {}
            }
        }
    }

    let mut eps = String::new();
    let mut gam = String::new();

    for idx in 0..zeroes.len() {
        let zero = zeroes[idx];
        let one = ones[idx];

        if zero > one {
            eps.push('0');
            gam.push('1');
        } else {
            eps.push('1');
            gam.push('0');
        }
    }

    let gamma = isize::from_str_radix(&gam, 2).unwrap();
    let epislon = isize::from_str_radix(&eps, 2).unwrap();

    println!("{}", gamma * epislon);
}
