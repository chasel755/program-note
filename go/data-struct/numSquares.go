package data-struct

/**
 * 完全平方数
 * 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
 * 输入: n = 12
 * 输出: 3 
 * 解释: 12 = 4 + 4 + 4.
 */
func numSquares(n int) int {
    if n <= 0 {
        return 0
    }
    //----查找n以内的所有平方数squares队列----
    var squares []int
    var sq = int(math.Sqrt(float64(n)))
    for i:=sq; i>=1; i-- {
        squares = append(squares, i * i)
    }
    //----广度优先遍历队列squares----
    result := BreadthFirstSearchNumSquares(n, squares)
    //fmt.Println(result[1:])
    return result[1]
}

//----广度优先遍历队列squares----
func BreadthFirstSearchNumSquares(n int, squares []int) []int {
    queue := [][]int{}
    for _, sq := range squares {
        left := n - sq
        pathNum := 1
        newItem := []int{left, pathNum, sq}
        if left == 0 {
            return newItem
        }
        queue = append(queue, newItem)
    }
    for len(queue) > 0 {
        item := queue[0]
        queue = queue[1:]
        for _, sq := range squares {
            left := item[0] - sq
            if left >= 0 {
                newItem := []int{left, item[1]+1, sq}
                newItem = append(newItem, item[2:]...)
                if left == 0 {
                    return newItem
                }
                queue = append(queue, newItem)
            }
        }
    }
    return []int{}
}