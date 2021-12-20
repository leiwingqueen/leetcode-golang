package binarysearch

//冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
//
//在加热器的加热半径范围内的每个房屋都可以获得供暖。
//
//现在，给出位于一条水平线上的房屋 houses 和供暖器 heaters 的位置，请你找出并返回可以覆盖所有房屋的最小加热半径。
//
//说明：所有供暖器都遵循你的半径标准，加热的半径也一样。
//
// 
//
//示例 1:
//
//输入: houses = [1,2,3], heaters = [2]
//输出: 1
//解释: 仅在位置2上有一个供暖器。如果我们将加热半径设为1，那么所有房屋就都能得到供暖。
//示例 2:
//
//输入: houses = [1,2,3,4], heaters = [1,4]
//输出: 1
//解释: 在位置1, 4上有两个供暖器。我们需要将加热半径设为1，这样所有房屋就都能得到供暖。
//示例 3：
//
//输入：houses = [1,5], heaters = [2]
//输出：3
// 
//
//提示：
//
//1 <= houses.length, heaters.length <= 3 * 104
//1 <= houses[i], heaters[i] <= 109
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/heaters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findRadius(houses []int, heaters []int) int {
	//找到房子的左边界和右边界
	left := houses[0]
	right := houses[len(houses)-1]
	//随便找一台供暖器，能够到达[min,max]这两个房子，最大值
	dis := max(abs(heaters[0]-left), abs(heaters[0]-right))
	if dis == 0 {
		return 0
	}
	//二分查找右边界
	l := 0
	r := dis - 1
	for l < r {
		mid := l + (r-l)/2
		if check(houses, heaters, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func check(houses []int, heaters []int, radius int) bool {
	//TODO:bst?
	return true
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}
