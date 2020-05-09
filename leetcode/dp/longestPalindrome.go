package main

// 最长回文子串

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	maxLen, from, to := 0, 0, 0
	for length := 1; length <= len(s); length++ {
		for start := 0; start <= len(s)-length; start++ {
			end := start + length - 1
			dp[start][end] = s[start] == s[end] && (length == 1 || length == 2 || dp[start+1][end-1])
			if dp[start][end] && length > maxLen {
				maxLen = end - start + 1
				from = start
				to = end
			}
		}
	}

	return s[from : to+1]
}

func main() {

}
