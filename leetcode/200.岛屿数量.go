/*
 * @lc app=leetcode.cn id=200 lang=golang
 * @lcpr version=30203
 *
 * [200] 岛屿数量
 */

// @lc code=start

var n, m int

func numIslands(grid [][]byte) int {
	n = len(grid)
	if n == 0 {
		return 0
	}
	m = len(grid[0])
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				dfs(&grid, i, j)
				ans++
			}
		}
	}

	return ans
}

// dfs 函数接收一个指向二维切片的指针
func dfs(grid *[][]byte, i int, j int) {
	// 安全检查：索引越界或当前单元格不是 '1'（岛屿）
	if i < 0 || j < 0 || i >= n || j >= m || (*grid)[i][j] != '1' {
		return
	}
	// 标记当前单元格为已访问（例如，置为 '0'）
	(*grid)[i][j] = '0'
	// 递归访问四个方向
	dfs(grid, i+1, j) // 下
	dfs(grid, i-1, j) // 上
	dfs(grid, i, j+1) // 右
	dfs(grid, i, j-1) // 左
}

// @lc code=end

/*
// @lcpr case=start
// [\n['1','1','1','1','0'],\n['1','1','0','1','0'],\n['1','1','0','0','0'],\n['0','0','0','0','0']\n]\n
// @lcpr case=end

// @lcpr case=start
// [\n['1','1','0','0','0'],\n['1','1','0','0','0'],\n['0','0','1','0','0'],\n['0','0','0','1','1']\n]\n
// @lcpr case=end

*/

