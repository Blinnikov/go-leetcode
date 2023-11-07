func longestCommonPrefix(strs []string) string {
    sort.Strings(strs)

    idx := 0
    first := strs[0]
    last := strs[len(strs) - 1]

    for idx < len(first) && idx < len(last) {
        if first[idx] != last[idx] {
            break
        }

        idx += 1
    }

    return first[:idx]
}
