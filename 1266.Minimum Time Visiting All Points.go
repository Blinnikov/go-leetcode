func minTimeToVisitAllPoints(points [][]int) int {
    res := 0

    for i := 0; i < len(points) - 1; i++ {
        res += getTime(points[i], points[i + 1])
    }
    
    return res
}

func getTime(from []int, to []int) int {
    return max(
        int(math.Abs(float64(from[0] - to[0]))),
        int(math.Abs(float64(from[1] - to[1]))))
}
