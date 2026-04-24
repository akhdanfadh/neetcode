type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var sb strings.Builder
    for _, s := range strs {
        fmt.Fprintf(&sb, "%d#%s", len(s), s)
    }
    return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
    var res []string
    i := 0
    for i < len(encoded) {
        j := i
        for encoded[j] != '#' { j++ } // find separator
        length, _ := strconv.Atoi(encoded[i:j]) // read length
        res = append(res, encoded[j+1:j+1+length]) // read exact bytes
        i = j+1+length
    }
    return res
}
