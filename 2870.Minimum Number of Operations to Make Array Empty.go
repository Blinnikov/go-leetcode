func minOperations(nums []int) int {
    m := make(map[int]int)

    for _, n := range nums {
        m[n]++
    }

    res := 0

    for _, v := range m {
        if v == 1 {
            return -1
        }

        res += v / 3

        if v % 3 != 0 {
            res++
        }
    }

    return res
}
