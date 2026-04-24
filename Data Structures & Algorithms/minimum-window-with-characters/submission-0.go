func minWindow(s string, t string) string {
	if t == "" { return "" }

	tCount := make(map[byte]int)
	for i := range len(t) { tCount[t[i]]++ }
	
	have, need := 0, len(tCount)
	res := [2]int{-1, -1}
	resLen := math.MaxInt32
	lId := 0
	winCount := make(map[byte]int)

	for rId := 0; rId < len(s); rId++ {
		if inT(tCount, s[rId]) {
			winCount[s[rId]]++
			if winCount[s[rId]] == tCount[s[rId]] { have++ }
		}

		for have == need { 
			if resLen > rId-lId+1 {
				res = [2]int{lId, rId}
				resLen = rId-lId+1
			}

			if inT(tCount, s[lId]) {
				winCount[s[lId]]--
				if winCount[s[lId]] < tCount[s[lId]] { have-- }
			}
			lId++
		}
	}

	if res[0] == -1 { return "" }
	return s[res[0]:res[1]+1]
}

func inT(m map[byte]int, c byte) bool {
	_, exists := m[c]
	return exists
}