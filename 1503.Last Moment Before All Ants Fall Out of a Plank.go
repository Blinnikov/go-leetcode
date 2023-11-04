func getLastMoment(n int, left []int, right []int) int {
    var res = 0

    for _, v := range left {
        res = max(res, v)
	}

    for _, v := range right {
        res = max(res, n - v)
	}

    return res
}
