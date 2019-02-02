package array

func IsCanCompleteCircuit(gas []int, cost []int) int {
	// find valid start pos
	for i := 0; i < len(gas); i++ {
		if gas[i] >= cost[i] {
			startPos := i
			save := gas[startPos] - cost[startPos]
			var j int
			for j = startPos + 1; j < len(gas) && j != startPos; j = (j + 1) % len(gas) {
				if gas[j]+save < cost[j] {
					break
				} else {
					save += gas[j] + save - cost[i]
				}
			}
			if j == startPos {
				return startPos
			}
		}
	}
	return -1
}
