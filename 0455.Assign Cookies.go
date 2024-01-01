func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)

    gIdx := len(g) - 1
    sIdx := len(s) - 1

    content := 0

    for gIdx >= 0 && sIdx >= 0 {
        if s[sIdx] >= g[gIdx] {
            content++
            sIdx--
            gIdx--
        } else {
            gIdx--
        }
    }

    return content
}
