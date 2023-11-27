func findAnagrams(s string, p string) []int {
    pattern := make([]byte, 26, 26)
    segment := make([]byte, 26, 26)

    for i := 0; i < len(p); i++ {
        pattern[p[i] - 'a']++
    }

    res := make([]int, 0)

    for i := 0; i <= len(s) - len(p); i++ {
        updateSegment(segment, s, i, i + len(p))
        if isAnagram(segment, pattern) {
            res = append(res, i)
        }
    }

    return res
}

func updateSegment(segment []byte, s string, start int, end int) {
    if start == 0 {
        for i := start; i < end; i++ {
            segment[s[i] - 'a']++
        }
    } else {
        charRemoving := s[start - 1] - 'a'
        charAdding := s[end - 1] - 'a'

        segment[charRemoving]--
        segment[charAdding]++
    }
    fmt.Println("Inside update", segment)
}

func isAnagram(segment []byte, pattern []byte) bool {
    for i := 0; i < len(segment); i++ {
        if segment[i] != pattern[i] {
            return false
        }
    }
    return true
}
