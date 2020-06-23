struct Solution {}

impl Solution {
    pub fn func(a: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut v: Vec<Vec<i32>> = a.clone();
        
        for i in 0..v.len() {
            v[i].reverse();
            for j in 0..v[i].len() {
                v[i][j] = if v[i][j] == 0 { 1 } else { 0 };    
            }
        }

        v
    }
}

fn main() {
    assert_eq!(vec![
        vec![1, 0, 0],
        vec![0, 1, 0],
        vec![1, 1, 1],
    ], Solution::func(vec![
        vec![1, 1, 0],
        vec![1, 0, 1],
        vec![0, 0, 0],
    ]));
    assert_eq!(vec![
        vec![1, 1, 0, 0],
        vec![0, 1, 1, 0],
        vec![0, 0, 0, 1],
        vec![1, 0, 1, 0],
    ], Solution::func(vec![
        vec![1, 1, 0, 0],
        vec![1, 0, 0, 1],
        vec![0, 1, 1, 1],
        vec![1, 0, 1, 0],
    ]));
}
