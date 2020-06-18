struct Solution {}

impl Solution {
    pub fn func(board: Vec<Vec<char>>) -> i32 {
        let mut c: i32 = 0;

        for i in 0..board.len() {
            for j in 0..board[i].len() {
                if (board[i][j] == 'X') &&
                   (i == 0 || board[i - 1][j] != 'X') &&
                   (j == 0 || board[i][j - 1] != 'X') {
                    c += 1;
                }
            }
        }

        c
    }
}

fn main() {
    assert_eq!(2, Solution::func(vec![
        vec!['X', '.', '.', 'X'],
        vec!['.', '.', '.', 'X'],
        vec!['.', '.', '.', 'X'],
    ]));
}
