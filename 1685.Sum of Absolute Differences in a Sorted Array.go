func getSumAbsoluteDifferences(nums []int) []int {
    ps := make([]int, len(nums))
    ss := make([]int, len(nums))

    ps[0] = nums[0]
    ss[len(nums) - 1] = nums[len(nums) - 1]

    for i:= 1; i < len(nums); i++ {
        ps[i] = ps[i - 1] + nums[i]
        ss[len(nums) - 1 - i] = ss[len(nums) - i] + nums[len(nums) - 1 - i]
    }

    res := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        res[i] = (nums[i] * i - ps[i]) + (ss[i] - (len(nums) - i - 1) * nums[i])
    }

    return res
}
