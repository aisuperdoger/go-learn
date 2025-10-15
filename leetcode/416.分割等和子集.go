/*
 * @lc app=leetcode.cn id=416 lang=golang
 * @lcpr version=30203
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	
	sum:= 0
	for _,v :=range nums {
		sum += v
	}

	if sum%2 ==1 {
		return false
	}

	sum = sum/2

    n := len(nums)
	dp := make([][]int,n+1)

	for i:=range dp{
		dp[i] = make([]int,sum+1)
	}

	for i:= 1;i<=n;i++ {
		for j:=0;j<=sum;j++{
			if j >= nums[i-1]{
				if dp[i-1][j] > dp[i-1][j-nums[i-1]]+nums[i-1]{
					dp[i][j] = dp[i-1][j]
				}else{
					dp[i][j] =dp[i-1][j-nums[i-1]]+nums[i-1]
				}
			}else{
				dp[i][j] = dp[i-1][j]
			}
		}
	}
		
	return dp[n][sum]==sum

}
// @lc code=end



/*
// @lcpr case=start
// [1,5,11,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,5]\n
// @lcpr case=end

 */

