package data-struct

func numIslands(grid [][]byte) int {
    count := 0
    width := 0
    height := len(grid)
    if height > 0 {
        width = len(grid[0])
    }
    for j:=0; j<height; j++ {
        for i:=0; i<width; i++ {
            count += BreadthFirstSearch(grid, j, i, height, width)
        }
    }
    return count
}

/*深度优先搜索*/
func DepthFirstSearch(grid [][]byte, j int, i int, height int, width int) int {
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
func BreadthFirstSearch(grid [][]byte, j int, i int, height int, width int) int {
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
