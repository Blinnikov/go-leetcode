func largestSubmatrix(matrix [][]int) int {
    for row := 1; row < len(matrix); row++ {
        for col := 0; col < len(matrix[row]); col++ {
            if matrix[row][col] == 1 {
                matrix[row][col] = matrix[row - 1][col] + 1
            }
        }
    }

    res := 0

    for _, row := range matrix {
        n := len(row)
        sort.Ints(row)

        for col := n - 1; col >= 0 && row[col] > 0; col-- {
            area := row[col] * (n - col)
            res = max(res, area)
        }
    }

    

    // for row := 0; row < len(matrix); row++ {
        
    //     for col := 1; col <= n; col++ {
    //         area := col * matrix[row][n - col]
    //         res = max(res, area)
    //     }
    // }
    
    // fmt.Println(matrix)

    return res
}
