func largestGoodInteger(num string) string {
    res := ""

    for i := 0; i < len(num) - 2; i++ {
        if num[i] == num[i + 1] && num[i + 1] == num[i + 2] {
            candidate := fmt.Sprintf("%s%s%s", string(num[i]), string(num[i + 1]), string(num[i + 2]))
            res = max(res, candidate)
        }
    }

    return res
}
