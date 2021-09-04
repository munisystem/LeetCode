impl Solution {
    pub fn tribonacci(n: i32) -> i32 {
        if n < 2 {
            return n;
        } else if n == 2 {
            return 1;
        }
        let mut n0 = 0;
        let mut n1 = 1;
        let mut n2 = 1;
        let mut tmp = 0;
        for _ in 0..n - 2 {
            tmp = n0 + n1 + n2;
            n0 = n1;
            n1 = n2;
            n2 = tmp;
        }
        return tmp;
    }
}
