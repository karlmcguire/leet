struct Solution {}

impl Solution {
    pub fn func(n: i32) -> i32 {
        let mut i: i32 = n;
        let mut r: i32 = 1;

        loop {
            i -= r;
            if i < 0 {
                r -= 1;
                break;
            }
            r += 1;
        }

        return r;
    }
}

fn main() {
    assert_eq!(0, Solution::func(0));
    assert_eq!(1, Solution::func(1));
    assert_eq!(2, Solution::func(5));
    assert_eq!(3, Solution::func(8));
    assert_eq!(13, Solution::func(100));
}
