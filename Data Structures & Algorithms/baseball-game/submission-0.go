func calPoints(operations []string) int {
     records := []int{}
	for _, op := range operations {
		if num, err := strconv.Atoi(op); err == nil {
			records = append(records, num)
		} else if op == "C" {
			//Disregard previous score
			records = records[:len(records)-1]
		} else if op == "D" {
			//Double previous score
			records = append(records, records[len(records)-1]*2)
		} else if op == "+" {
			//Sum of previous two scores
			records = append(records, records[len(records)-1]+records[len(records)-2])
		}
	}
	sum := 0
	for _, score := range records {
		sum += score
	}
	return sum

}
