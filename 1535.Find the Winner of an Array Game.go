func getWinner(arr []int, k int) int {
    curr := arr[0]
    win := 0

    for i := 1; i < len(arr); i++ {
		if arr[i] > curr {
            curr = arr[i]
            win = 0
        }

        win++

        if win == k {
            break
        }
	}

    return curr
}
