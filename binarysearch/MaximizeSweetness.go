package binarysearch

import "math"

//1231. 分享巧克力
//你有一大块巧克力，它由一些甜度不完全相同的小块组成。我们用数组 sweetness 来表示每一小块的甜度。
//
//你打算和 K 名朋友一起分享这块巧克力，所以你需要将切割 K 次才能得到 K+1 块，每一块都由一些 连续 的小块组成。
//
//为了表现出你的慷慨，你将会吃掉 总甜度最小 的一块，并将其余几块分给你的朋友们。
//
//请找出一个最佳的切割策略，使得你所分得的巧克力 总甜度最大，并返回这个 最大总甜度。
//
// 
//
//示例 1：
//
//输入：sweetness = [1,2,3,4,5,6,7,8,9], K = 5
//输出：6
//解释：你可以把巧克力分成 [1,2,3], [4,5], [6], [7], [8], [9]。
//示例 2：
//
//输入：sweetness = [5,6,7,8,9,1,2,3,4], K = 8
//输出：1
//解释：只有一种办法可以把巧克力分成 9 块。
//示例 3：
//
//输入：sweetness = [1,2,2,1,2,2,1,2,2], K = 2
//输出：5
//解释：你可以把巧克力分成 [1,2,2], [1,2,2], [1,2,2]。
// 
//
//提示：
//
//0 <= K < sweetness.length <= 10^4
//1 <= sweetness[i] <= 10^5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/divide-chocolate
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func maximizeSweetness(sweetness []int, k int) int {
	min := math.MaxInt32
	sum := 0
	for _, sweet := range sweetness {
		if sweet < min {
			min = sweet
		}
		sum += sweet
	}
	l := min
	r := sum
	for l < r {
		mid := l + (r-l+1)/2
		if check(sweetness, k, mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func check(sweetness []int, k int, s int) bool {
	idx := 0
	cnt := 0
	for idx < len(sweetness) && cnt < k+1 {
		sum := 0
		for idx < len(sweetness) && sum < s {
			sum += sweetness[idx]
			idx++
		}
		if sum < s {
			return false
		}
		cnt++
	}
	return cnt >= k+1
}
