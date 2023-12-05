func numberOfMatches(n int) int {
    return n - 1

    res := 0

    for n > 1 {
        res += n / 2 
        n = (n + 1) / 2
    }

    return res
}
