import "slices"
func shipWithinDays(weights []int, days int) int {
	maxWeight := slices.Max(weights)
	if len(weights) <= days {
		return maxWeight
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
	for true {
		if tryWeight(maxWeight) == true {
			return maxWeight
		} else {
			maxWeight++
		}
	}
	return maxWeight

}

