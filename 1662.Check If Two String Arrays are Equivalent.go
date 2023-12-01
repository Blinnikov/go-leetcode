func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    var b strings.Builder

    for _, w := range word1 {
        b.WriteString(w)
    }

    str1 := b.String()

    b.Reset()

    for _, w := range word2 {
        b.WriteString(w)
    }

    str2 := b.String()

    return str1 == str2
}
