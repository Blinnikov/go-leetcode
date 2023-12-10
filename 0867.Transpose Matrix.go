func transpose(matrix [][]int) [][]int {
    res := make([][]int, len(matrix[0]))

    for row := range matrix {
        for col := range matrix[row] {
            if row == 0 {
                res[col] = make([]int, len(matrix))
            }
            res[col][row] = matrix[row][col]
        }
    }

    return res
}
