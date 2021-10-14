/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

package main

func guessNumber(n int) int {
	left, right := 0, n
	var mid int
	for {
		mid := (left + right) / 2
		switch guess(mid) {
		case 0:
			return mid
		case 1:
			left = mid + 1
		case -1:
			right = mid
		}
	}
	return mid
}
