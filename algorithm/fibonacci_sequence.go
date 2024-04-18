package algorithm

import "math"

/*
动态规划实现：
从数学角度来看，斐波那契数列的递推关系可以定义为：F(0) = 0，F(1) = 1，以及对于n > 1，F(n) = F(n-1) + F(n-2)。
斐波那契数列的递推公式和通项公式如下：
1. 递推公式。F(n)=F(n-1)+F(n-2)（n≥2），F(0)=0，F(1)=1。
2. 通项公式。2.1 封闭公式。Fn=((1+√5)/2)^n-((1-√5)/2)^n]/√5。
3. 通项公式。2.2 扩展形式。Fn=((1+√5)/2)^(n-1)-((1-√5)/2)^(n-1))/√5。

算法应用：
1. 爬楼梯：假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
	适用于n比较小的情况
*/
// 递归实现斐波那契数列，时间复杂度：O(n^2)
func FibonacciSequenceRecursion(n int) int {
	if n < 2 {
		return n
	}

	return FibonacciSequenceRecursion(n-1) + FibonacciSequenceRecursion(n-2)
}

// 非递归实现斐波那契数列，时间复杂度：O(n)
func FibonacciSequence(n int) int {
	if n < 2 {
		return n
	}

	prev1, prev2 := 0, 1

	for i := 2; i < n; i++ {
		temp := prev1 + prev2
		prev1 = prev2
		prev2 = temp
	}

	return prev2
}

// 按照斐波那契数列通向公式实现
func FibonacciSequenceCommon(n int) int {
	sqrt5 := math.Sqrt(5)
	pow1 := math.Pow((1+sqrt5)/2, float64(n))
	pow2 := math.Pow((1-sqrt5)/2, float64(n))
	return int(math.Round((pow1 - pow2) / sqrt5))
}
