package main

import "fmt"

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
	moneys := [...]int{1,5,10,20,100}
	for i := 0 ; i < 5 ; i++ {
		dp[i][0] = 1 // 用前i种面额拼凑1rmb的方法数均为1
	}
	for i := 0 ; i < num ; i++ {
		dp[0][i] = 1 //用1rmb面额拼凑0金额的方法数都为1
	}
}