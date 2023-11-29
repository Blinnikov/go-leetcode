func hammingWeight(num uint32) int {
    var res uint32 = 0

    for i := 0; i < 32; i++ {
        res += num & 1
        num >>= 1
    }

    return int(res)
}
