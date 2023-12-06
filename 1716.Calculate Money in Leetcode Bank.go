func totalMoney(n int) int {
    sum := 0
    prev := 0

    for i := 0; i < n; i++ {
        if i % 7 == 0 { // it's Monday
            prev = (i / 7) + 1
        } else {
            prev++
        }

        sum += prev
    }

    return sum
}
