func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
    dist := max(math.Abs(float64(fx - sx)), math.Abs(float64(fy - sy)))

    fmt.Println(dist)

    if dist == 0 && t != 1 {
        return true
    }

    if dist > 0 {
        return float64(t) >= dist
    }

    return false
}
