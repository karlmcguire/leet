struct Solution {}

impl Solution {
    pub fn func(points: Vec<Vec<i32>>) -> f64 {
        fn area(v: [f64; 6]) -> f64 {
            ((v[0] * (v[3] - v[5])) +
             (v[2] * (v[5] - v[1])) +
             (v[4] * (v[1] - v[3]))) / 2.0
        }

        let mut m: f64 = 0.0;

        for i in 0..points.len() {
            for j in 0..points.len() {
                for k in 0..points.len() {
                    let a: f64 = area([
                        points[i][0] as f64, points[i][1] as f64,
                        points[j][0] as f64, points[j][1] as f64,
                        points[k][0] as f64, points[k][1] as f64,
                    ]); 

                    if a > m {
                        m = a;
                    }
                }
            }
        }

        return m;
    }
}

fn main() {
    assert_eq!(2.0, Solution::func(vec![
        vec![0, 0],
        vec![0, 1],
        vec![1, 0],
        vec![0, 2],
        vec![2, 0],
    ]));
}
