func canCompleteCircuit(gas []int, cost []int) int {
	max, idx, accum := 0, 0, 0
	for i := len(gas) - 1; i > -1; i-- {
		accum += gas[i] - cost[i]
		if accum > max {
			max, idx = accum, i
		}
	}
	if accum < 0 {
		return -1
	}
	return idx
}
