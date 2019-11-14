package data-struct

/**
岛屿数量
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

示例 1:
输入:
11110
11010
11000
00000
输出: 1

示例 2:
输入:
11000
11000
00100
00011
输出: 3
*/
func numIslands(grid [][]byte) int {
    count := 0
    width := 0
    height := len(grid)
    if height > 0 {
        width = len(grid[0])
    }
    for j:=0; j<height; j++ {
        for i:=0; i<width; i++ {
            count += breadthFirstSearchNumIslands(grid, j, i, height, width)
        }
    }
    return count
}

/*深度优先搜索*/
func depthFirstSearchNumIslands(grid [][]byte, j int, i int, height int, width int) int {
    if j<0 || j >= height || i<0 || i>=width {
        return 0
    }
    if grid[j][i] == '1' {
        grid[j][i] = '0'
        DepthFirstSearch(grid, j-1, i, height, width)
        DepthFirstSearch(grid, j+1, i, height, width)
        DepthFirstSearch(grid, j, i-1, height, width)
        DepthFirstSearch(grid, j, i+1, height, width)
        return 1
    }
    return 0
}

/*广度优先搜索*/
func breadthFirstSearchNumIslands(grid [][]byte, j int, i int, height int, width int) int {
    if j<0 || j >= height || i<0 || i>=width {
        return 0
    }
    if grid[j][i] != '1' {
        return 0
    }
    
    queueHead := 0
    queueTail := 0
    queue := make(map[int][]int)
    queue[queueHead] = []int{j, i}
    for queueHead <= queueTail  {
        j = queue[queueHead][0]
        i = queue[queueHead][1]
        delete(queue, queueHead)
        queueHead++
        if grid[j][i] == '0' { //跳过重复放入队列的项
            continue
        }
        grid[j][i] = '0'
        if i > 0 && grid[j][i-1] == '1' {
            queueTail++
            queue[queueTail] = []int{j, i-1}
        }
        if i < width-1 && grid[j][i+1] == '1' {
            queueTail++
            queue[queueTail] = []int{j, i+1}
        }
        if j > 0 && grid[j-1][i] == '1' {
            queueTail++
            queue[queueTail] = []int{j-1, i}
        }
        if j < height-1 && grid[j+1][i] == '1' {
            queueTail++
            queue[queueTail] = []int{j+1, i}
        }
    }
    return 1
}
