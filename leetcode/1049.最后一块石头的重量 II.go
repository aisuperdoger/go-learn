/*
 * @lc app=leetcode.cn id=1049 lang=golang
 * @lcpr version=30203
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
func lastStoneWeightII(stones []int) int {
    sum := 0

	for _,v :=range stones {
		sum+= v
	}

	half := sum/2


	dp := make([]int,half+1)

	for i:=0;i<len(stones);i++{
		for j:=len(dp)-1;j>=stones[i];j-- {
			dp[j] = max(dp[j],dp[j-stones[i]]+stones[i])
		}
	}

	return sum-2*dp[len(dp)-1] 

}


func max(a,b int) int{
	if a>b {
		return a
	}else{
		return b
	}
}
// @lc code=end



/*
// @lcpr case=start
// [2,7,4,1,8,1]\n
// @lcpr case=end

// @lcpr case=start
// [31,26,33,21,40]\n
// @lcpr case=end

 */

