impl Solution {
    pub fn is_sum_equal(first_word: String, second_word: String, target_word: String) -> bool {
        let mut first;
        let mut second;
        let mut target;
        for c in first_word.chars() {
            first = first * 10 + c as i32 - 97;
        }
        for c in second_word.chars() {
            second = second * 10 + c as i32 - 97;
        }
        for c in target_word.chars() {
            target = target * 10 + c as i32 - 97;
        }
        return first + second == target;
    }
}
