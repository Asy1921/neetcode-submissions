
// @lc code=start
import "slices"

func shipWithinDays(weights []int, days int) int {
	maxWeight := slices.Max(weights)
	if len(weights) <= days {
		return maxWeight
	}
	totalWeight := 0
	for _, weight := range weights {
		totalWeight += weight
	}
	tryWeight := func(maxWeight int) bool {
		daySpent := 1
		for i := 0; i < len(weights); {
			dailyCap := weights[i]
			i++
			for true {
				if i < len(weights) && dailyCap+weights[i] <= maxWeight {
					dailyCap += weights[i]
					i++
				} else {
					break
				}
			}
			if i == len(weights) {
				return true
			}
			daySpent++
			if daySpent > days {
				return false
			}
		}
		return true
	}
	l := maxWeight
	r := totalWeight
	res := totalWeight
	for l <= r {
		mid := l + (r-l)/2
		if tryWeight(mid) {
			if mid < res {
				res = mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res

}

// @lc code=end
