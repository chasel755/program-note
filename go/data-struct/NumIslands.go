package data-struct

func NumIslands(grid [][]byte) int {
    count := 0
    width := 0
    height := len(grid)
    if height >0 {
        width = len(grid[0])
    }
    for j:=0; j<height; j++ {
        for i:=0; i<width; i++ {
            count += _numIslands(grid, j, i, height, width)
        }
    }
    return count
}

func _numIslands(grid [][]byte, j int, i int, height int, width int) int {
    if j<0 || j >= height || i<0 || i>=width {
        return 0
    }
    if grid[j][i] == '1' {
        grid[j][i] = 0
        _numIslands(grid, j-1, i, height, width)
        _numIslands(grid, j+1, i, height, width)
        _numIslands(grid, j, i-1, height, width)
        _numIslands(grid, j, i+1, height, width)
        return 1
    }
    return 0
}