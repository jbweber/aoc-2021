use std::fs::File;
use std::io::Read;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let (calls, mut boards) = load_input();

    let mut winners = vec![];
    let mut callxs = vec![];

    for &callx in calls.iter() {
        if boards.is_empty() {
            break;
        }

        println!("{}", boards.len());

        call(callx, &mut boards);

        let mut remove_idx: Option<usize> = None;

        for (idx, board) in boards.iter_mut().enumerate() {
            if board.is_winner() {
                remove_idx = Some(idx);
                break;
            }
        }

        match remove_idx {
            Some(idx) => {
                println!("board.len() less next row");
                let last_board = boards.remove(idx);
                println!("{:?}", last_board);
                winners.push(last_board);
                callxs.push(callx);
            }
            None => {}
        }
    }

    // print_board(&winners[winners.len() - 1]);

    println!(
        "{}",
        calc_winner(callxs[callxs.len() - 1], &winners[winners.len() - 1])
    )
}

fn calc_winner(callx: i32, board: &BingoBoard) -> i32 {
    let mut sum: i32 = 0;

    if !board.one_0.called {
        sum += board.one_0.val
    }
    if !board.one_1.called {
        sum += board.one_1.val
    }
    if !board.one_2.called {
        sum += board.one_2.val
    }
    if !board.one_3.called {
        sum += board.one_3.val
    }
    if !board.one_4.called {
        sum += board.one_4.val
    }

    if !board.two_0.called {
        sum += board.two_0.val
    }
    if !board.two_1.called {
        sum += board.two_1.val
    }
    if !board.two_2.called {
        sum += board.two_2.val
    }
    if !board.two_3.called {
        sum += board.two_3.val
    }
    if !board.two_4.called {
        sum += board.two_4.val
    }

    if !board.three_0.called {
        sum += board.three_0.val
    }
    if !board.three_1.called {
        sum += board.three_1.val
    }
    if !board.three_2.called {
        sum += board.three_2.val
    }
    if !board.three_3.called {
        sum += board.three_3.val
    }
    if !board.three_4.called {
        sum += board.three_4.val
    }

    if !board.four_0.called {
        sum += board.four_0.val
    }
    if !board.four_1.called {
        sum += board.four_1.val
    }
    if !board.four_2.called {
        sum += board.four_2.val
    }
    if !board.four_3.called {
        sum += board.four_3.val
    }
    if !board.four_4.called {
        sum += board.four_4.val
    }

    if !board.five_0.called {
        sum += board.five_0.val
    }
    if !board.five_1.called {
        sum += board.five_1.val
    }
    if !board.five_2.called {
        sum += board.five_2.val
    }
    if !board.five_3.called {
        sum += board.five_3.val
    }
    if !board.five_4.called {
        sum += board.five_4.val
    }

    return sum * callx;
}

fn call(num: i32, boards: &mut Vec<BingoBoard>) {
    for board in boards.iter_mut() {
        board.one_0.call(num);
        board.one_1.call(num);
        board.one_2.call(num);
        board.one_3.call(num);
        board.one_4.call(num);

        board.two_0.call(num);
        board.two_1.call(num);
        board.two_2.call(num);
        board.two_3.call(num);
        board.two_4.call(num);

        board.three_0.call(num);
        board.three_1.call(num);
        board.three_2.call(num);
        board.three_3.call(num);
        board.three_4.call(num);

        board.four_0.call(num);
        board.four_1.call(num);
        board.four_2.call(num);
        board.four_3.call(num);
        board.four_4.call(num);

        board.five_0.call(num);
        board.five_1.call(num);
        board.five_2.call(num);
        board.five_3.call(num);
        board.five_4.call(num);

        // println!("{:?}", board);
    }
}

fn find_winner(boards: &Vec<BingoBoard>) -> Option<&BingoBoard> {
    for board in boards.iter() {
        if board.is_winner() {
            return Some(board);
        }
    }
    return None;
}

fn load_input() -> (Vec<i32>, Vec<BingoBoard>) {
    let file = std::env::args().nth(1).expect("no filename given");
    let mut path_buf = std::env::current_dir().unwrap();

    path_buf.push(file);

    let path = path_buf.as_path().to_str().unwrap();

    let input = std::fs::read_to_string(path).unwrap();

    let mut lines = input.lines().collect::<Vec<&str>>();

    let calls_str = lines[0];
    let calls: Vec<i32> = calls_str
        .split(",")
        .map(|s| s.parse().unwrap())
        .collect::<Vec<i32>>();

    lines.remove(0);

    let mut keep_going = true;
    let mut boards = vec![];

    while keep_going {
        if lines.len() < 6 {
            panic!("expected at least 6 lines, got {}", lines.len());
        }

        let mut board = BingoBoard::new();

        lines.remove(0);
        let mut one = create_board(&mut lines);
        board.one_0 = one.remove(0);
        board.one_1 = one.remove(0);
        board.one_2 = one.remove(0);
        board.one_3 = one.remove(0);
        board.one_4 = one.remove(0);

        let mut two = create_board(&mut lines);
        board.two_0 = two.remove(0);
        board.two_1 = two.remove(0);
        board.two_2 = two.remove(0);
        board.two_3 = two.remove(0);
        board.two_4 = two.remove(0);

        let mut three = create_board(&mut lines);
        board.three_0 = three.remove(0);
        board.three_1 = three.remove(0);
        board.three_2 = three.remove(0);
        board.three_3 = three.remove(0);
        board.three_4 = three.remove(0);

        let mut four = create_board(&mut lines);
        board.four_0 = four.remove(0);
        board.four_1 = four.remove(0);
        board.four_2 = four.remove(0);
        board.four_3 = four.remove(0);
        board.four_4 = four.remove(0);

        let mut five = create_board(&mut lines);
        board.five_0 = five.remove(0);
        board.five_1 = five.remove(0);
        board.five_2 = five.remove(0);
        board.five_3 = five.remove(0);
        board.five_4 = five.remove(0);

        // let board = vec![one, two, three, four, five];

        boards.push(board);
        keep_going = lines.len() > 1;
    }

    (calls, boards)
}

fn create_board(lines: &mut Vec<&str>) -> Vec<BingoCell> {
    lines
        .remove(0)
        .trim()
        .split_ascii_whitespace()
        .map(|s| BingoCell {
            val: s.parse().unwrap(),
            called: false,
        })
        .collect::<Vec<BingoCell>>()
}

#[derive(Debug)]
struct BingoCell {
    val: i32,
    called: bool,
}

impl BingoCell {
    fn call(&mut self, mat: i32) {
        if mat == self.val {
            self.called = true;
        }
    }
}

#[derive(Debug)]
struct BingoBoard {
    one_0: BingoCell,
    one_1: BingoCell,
    one_2: BingoCell,
    one_3: BingoCell,
    one_4: BingoCell,
    two_0: BingoCell,
    two_1: BingoCell,
    two_2: BingoCell,
    two_3: BingoCell,
    two_4: BingoCell,
    three_0: BingoCell,
    three_1: BingoCell,
    three_2: BingoCell,
    three_3: BingoCell,
    three_4: BingoCell,
    four_0: BingoCell,
    four_1: BingoCell,
    four_2: BingoCell,
    four_3: BingoCell,
    four_4: BingoCell,
    five_0: BingoCell,
    five_1: BingoCell,
    five_2: BingoCell,
    five_3: BingoCell,
    five_4: BingoCell,
}

impl BingoBoard {
    fn new() -> BingoBoard {
        return BingoBoard {
            one_0: BingoCell {
                val: 0,
                called: false,
            },
            one_1: BingoCell {
                val: 0,
                called: false,
            },
            one_2: BingoCell {
                val: 0,
                called: false,
            },
            one_3: BingoCell {
                val: 0,
                called: false,
            },
            one_4: BingoCell {
                val: 0,
                called: false,
            },
            two_0: BingoCell {
                val: 0,
                called: false,
            },
            two_1: BingoCell {
                val: 0,
                called: false,
            },
            two_2: BingoCell {
                val: 0,
                called: false,
            },
            two_3: BingoCell {
                val: 0,
                called: false,
            },
            two_4: BingoCell {
                val: 0,
                called: false,
            },
            three_0: BingoCell {
                val: 0,
                called: false,
            },
            three_1: BingoCell {
                val: 0,
                called: false,
            },
            three_2: BingoCell {
                val: 0,
                called: false,
            },
            three_3: BingoCell {
                val: 0,
                called: false,
            },
            three_4: BingoCell {
                val: 0,
                called: false,
            },
            four_0: BingoCell {
                val: 0,
                called: false,
            },
            four_1: BingoCell {
                val: 0,
                called: false,
            },
            four_2: BingoCell {
                val: 0,
                called: false,
            },
            four_3: BingoCell {
                val: 0,
                called: false,
            },
            four_4: BingoCell {
                val: 0,
                called: false,
            },
            five_0: BingoCell {
                val: 0,
                called: false,
            },
            five_1: BingoCell {
                val: 0,
                called: false,
            },
            five_2: BingoCell {
                val: 0,
                called: false,
            },
            five_3: BingoCell {
                val: 0,
                called: false,
            },
            five_4: BingoCell {
                val: 0,
                called: false,
            },
        };
    }

    fn is_winner(&self) -> bool {
        // check by row
        let r1 = self.one_0.called
            && self.one_1.called
            && self.one_2.called
            && self.one_3.called
            && self.one_4.called;

        let r2 = self.two_0.called
            && self.two_1.called
            && self.two_2.called
            && self.two_3.called
            && self.two_4.called;

        let r3 = self.three_0.called
            && self.three_1.called
            && self.three_2.called
            && self.three_3.called
            && self.three_4.called;

        let r4 = self.four_0.called
            && self.four_1.called
            && self.four_2.called
            && self.four_3.called
            && self.four_4.called;

        let r5 = self.five_0.called
            && self.five_1.called
            && self.five_2.called
            && self.five_3.called
            && self.five_4.called;

        // check by column
        let c1 = self.one_0.called
            && self.two_0.called
            && self.three_0.called
            && self.four_0.called
            && self.five_0.called;

        let c2 = self.one_1.called
            && self.two_1.called
            && self.three_1.called
            && self.four_1.called
            && self.five_1.called;

        let c3 = self.one_2.called
            && self.two_2.called
            && self.three_2.called
            && self.four_2.called
            && self.five_2.called;

        let c4 = self.one_3.called
            && self.two_3.called
            && self.three_3.called
            && self.four_3.called
            && self.five_3.called;

        let c5 = self.one_4.called
            && self.two_4.called
            && self.three_4.called
            && self.four_4.called
            && self.five_4.called;

        return (r1 || r2 || r3 || r4 || r5) || (c1 || c2 || c3 || c4 || c5);
    }
}
