package numbers

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; i++ {
		tank := 0
		for j := i; j < i+n; j++ {
			k := j % n
			tank += gas[k] - cost[k]
			if tank < 0 {
				break
			}
		}
		if tank >= 0 {
			return i
		}
	}
	return -1
}

func canCompleteCircuitSkip(gas []int, cost []int) int {
	n := len(gas)
	j := 0
	for i := 0; i < n; i = j + 1 {
		tank := 0
		for j = i; j < i+n; j++ {
			k := j % n
			tank += gas[k] - cost[k]
			if tank < 0 {
				break
			}
		}
		if tank >= 0 {
			return i
		}
	}
	return -1
}

func canCompleteCircuitGreedy(gas []int, cost []int) int {
	n := len(gas)
	start := 0
	total := 0
	tank := 0
	for i := 0; i < n; i++ {
		total += gas[i] - cost[i]
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	if total < 0 {
		return -1
	} else {
		return start
	}
}
