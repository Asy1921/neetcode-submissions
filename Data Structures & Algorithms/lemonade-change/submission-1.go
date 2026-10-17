func lemonadeChange(bills []int) bool {
	hand:= make(map[int]int)
	for i:=0;i<len(bills);i++ {
		// Now we have to keep 5$ and return rest if any
		hand[bills[i]]++
		if bills[i]==5{
			continue
		} else if bills[i]==10 {
			if hand[5]>0 {
				hand[5]--
				continue
			} else {
			
				return false
			}
		} else if bills[i]==20 {
			if hand[10]>0 && hand[5]>0 {
				hand[10]--
				hand[5]--
				continue
			} else if hand[5]>2 {
				hand[5]=hand[5]-3
				continue
			} else {
			
				return false
			}
		}
		 
	}
	return true
}
