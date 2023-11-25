const mod = 1_000_000_007

func countNicePairs(nums []int) int {
    m := make(map[int]int)

    for _, n := range nums {
        m[n - reverse(n)] += 1
    }

    res := 0

    for _, v := range m {
        perm := (v * (v - 1)) / 2
        res = (res + perm) % mod
    }

    return res
}

func reverse(num int) int {
    res := 0

    for num > 0 {
        res = res * 10 + num % 10
        num /= 10
    }

    return res
}
