/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	low := 0
	high := n
	for low <= high {
		mid := low + (high-low)/2

		res := guess(mid)
		switch res {
		case 0:
			return mid
		case -1:
			high = mid - 1
		case 1:
			low = mid + 1
		}
	}
	return -1
}
