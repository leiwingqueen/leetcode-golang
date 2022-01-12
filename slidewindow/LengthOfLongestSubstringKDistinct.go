package slidewindow

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	mp := make(map[byte]int)
	cnt := 0
	l := 0
	r := 0
	res := 0
	for r < len(s) {
		//窗口右移
		if mp[s[r]] == 0 {
			cnt++
		}
		mp[s[r]]++
		r++
		if cnt <= k {
			if r-l > res {
				res = r - l
			}
		} else {
			//窗口左移
			for l < r && cnt > k {
				mp[s[l]]--
				if mp[s[l]] == 0 {
					cnt--
				}
				l++
			}
		}
	}
	return res
}
