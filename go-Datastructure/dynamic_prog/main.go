package main

//动态规划demo

/*
	给你六种面额 1、5、10、20、50、100 元的纸币，
	假设每种币值的数量都足够多，编写程序求组成N元（N为0~10000的非负整数）的不同组合的个数。
	[参考](https://www.cnblogs.com/freephp/p/11997846.html)
*/

func main() {

}

func getTheNum(num int) int {
	var dp [5][10000]int
	moneys := [...]int{1, 5, 10, 20, 50, 100}
	for i := 0; i < 5; i++ {
		dp[i][0] = 1 // 用前i种面额拼凑1rmb的方法数均为1
	}
	for i := 0; i <= num; i++ {
		dp[0][i] = 1 //用1rmb面额拼凑0金额的方法数都为1
	}

	for i := 0; i < 5; i++ {
		for j := 1; j <= num; j++ {
			if j-moneys[i] >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-moneys[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[4][num]
}

func SimpleGetNum(num int) int {
	var dp [10000]int
	moneys := [...]int{1, 5, 10, 20, 50, 100} // 面额数组
	for i := 0; i <= num; i++ {               // 用1rmb面额拼凑0金额的方法数都为1
		dp[i] = 1
	}

	for j := 1; j <= 5; j++ {
		for i := 1; i <= num; i++ {
			if i >= moneys[j] { // 当前面金额大于面额的时候需要计算前i种面额组合出i - moneys[j]的方法数
				dp[i] += dp[i-moneys[j]]
			}
		}
	}

	return dp[num]
}
