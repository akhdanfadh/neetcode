func canConstruct(ransomNote string, magazine string) bool {
    count := [26]int{}
    for _, c := range ransomNote {
        count[c-'a']++
    }

    countM := [26]int{}
    for _, c := range magazine {
        countM[c-'a']++
    }

    for i, c := range count {
        if countM[i] - c < 0 {
            return false
        }
    }
    return true
}