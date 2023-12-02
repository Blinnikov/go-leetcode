func countCharacters(words []string, chars string) int {
    p := toArray(chars)

    res := 0

    for _, word := range words {
        s := toArray(word)
        if contains(&p, &s) {
            res += len(word)
        }
    }

    return res
}

func toArray(s string) [26]int {
    res := [26]int{}

    for i := 0; i < len(s); i++ {
        res[s[i] - 'a']++
    }

    return res
}

func contains(p *[26]int, s *[26]int) bool {
    for i := 0; i < 26; i++ {
        if p[i] < s[i] {
            return false
        }
    }

    return true
}
