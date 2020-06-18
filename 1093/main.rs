struct Solution {}

impl Solution {
    pub fn func(count: Vec<i32>) -> Vec<f64> {
        let mut min: f64 = 255.0;
        let mut max: f64 = 0.0;
        let mut avg: f64 = 0.0;
        let mut med: f64 = 0.0;
        let mut rep: f64 = 0.0;

        // pointers
        let mut l: usize = 0;
        let mut r: usize = 255;

        // new pointers
        let mut nl: usize = 0;
        let mut nr: usize = 0;

        // median
        let mut m1: usize = 0;
        let mut m2: usize = 0;

        while l <= r {
            // ignore zeros
            while count[l] == 0 { l += 1; }
            while count[r] == 0 { r -= 1; }
            if nl < nr {
                // average
                avg += count[l] as f64 * l as f64;
                // new left position
                nl += count[l] as usize;
                // mode
                if count[l] > count[rep as usize] { rep = l as f64; }
                // min
                if l < min as usize { min = l as f64; }
                // new left position
                m1 = l;
                l += 1;
            } else {
                // average
                avg += count[r] as f64 * r as f64;
                // new right position
                nr += count[r] as usize;
                // mode
                if count[r] > count[rep as usize] { rep = r as f64; }
                // max
                if r > max as usize { max = r as f64; }
                // new right position 
                m2 = r;
                r -= 1;
            }
        }

        // finish average
        avg /= nl as f64 + nr as f64;

        // median
        if nl < nr { med = m2 as f64; }
        else if nl > nr { med = m1 as f64; }
        else { med = ((m1 as f64) + (m2 as f64)) / 2.0 }

        return vec![min, max, avg, med, rep];
    }
}

fn main() {
    assert_eq!(
        vec![1.0, 3.0, 2.375, 2.5, 3.0], 
        Solution::func(vec![
            0,1,3,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
        ]),
    );
    assert_eq!(
        vec![1.0, 4.0, 2.1818181818181817, 2.0, 1.0],
        Solution::func(vec![
            0,4,3,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
            0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
        ]),
    );
}
