struct Solution {}

impl Solution {
    pub fn func(flowerbed: Vec<i32>, n: i32) -> bool {
        let mut a: i32 = n;
        let mut i: usize = 0;

        while i < flowerbed.len() {
            if flowerbed[i] == 1 { 
                i += 2;
            } else if i + 1 < flowerbed.len() && flowerbed[i + 1] == 1 { 
                i += 3;
            } else { 
                i += 2;
                a -= 1;
            }
        }

        a <= 0
    }
}

fn main() {
    assert_eq!(true, Solution::func(vec![1, 0, 0, 0, 1], 1));
    assert_eq!(false, Solution::func(vec![1, 0, 0, 0, 1], 2));
    assert_eq!(true, Solution::func(vec![1, 0, 0, 0, 1, 0, 0], 2));
}
