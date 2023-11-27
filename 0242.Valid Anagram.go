func isAnagram(s string, t string) bool {
    sb := []byte(s)
    sort.Slice(sb, func(i, j int) bool {
        return sb[i] < sb[j]
    })

    s = string(sb)

    fmt.Println(string(sb))

    tb := []byte(t)
    sort.Slice(tb, func(i, j int) bool {
        return tb[i] < tb[j]
    })

    t = string(tb)

    return s == t
}
