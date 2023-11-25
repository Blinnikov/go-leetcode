func longestSubarray(nums []int) int {
    lo := 0
    zeroes := 1

    res := 0

    for hi, num := range nums {
        if num == 0 {
            zeroes -= 1
        }

        for zeroes < 0 {
            if nums[lo] == 0 {
                zeroes++
            }

            lo++
        }

        res = max(res, hi - lo)
    }

    return res
}
