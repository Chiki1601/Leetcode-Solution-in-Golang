func mincostTickets(days, costs []int) int {
	mem := make(map[int]int, 365)
	var recursion func([]int, int, int, int) int
	recursion = func(days []int, day, availableDays, spent int) int {
		if len(days) == 0 {
			return spent
		}
		if availableDays != 0 {
			if days[0] == day {
				days = days[1:]
			}
			return recursion(days, day+1, availableDays-1, spent)
		}
		if days[0] != day {
			return recursion(days, day+1, 0, spent)
		}
		days = days[1:]
		if ret, ok := mem[day]; ok {
			return spent + ret
		}
		buy1 := recursion(days, day, availableDays+1, spent+costs[0])
		buy7 := recursion(days, day, availableDays+7, spent+costs[1])
		buy30 := recursion(days, day, availableDays+30, spent+costs[2])
		out := min(min(buy1, buy7), buy30)
		mem[day] = out - spent
		return out
	}
	return recursion(days, 1, 0, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
