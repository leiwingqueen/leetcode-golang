package dp

func change(amount int, coins []int) int {
	if amount <= 0 {
		return 1
	}
	dp := make([][]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = make([]int, len(coins))
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] == 0 {
				dp[i][j] = 1
			} else if i-coins[j] > 0 {
				for k := 0; k <= j; k++ {
					dp[i][j] += dp[i-coins[j]][k]
				}
			}
		}
	}
	sum := 0
	for i := 0; i < len(coins); i++ {
		sum += dp[amount][i]
	}
	return sum
}
