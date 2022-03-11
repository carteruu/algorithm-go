package leet

func longestValidParentheses3(s string) int {
	dp := make([]int, len(s)+1)
	ans := 0
	for i := 2; i <= len(s); i++ {
		//当前字符
		if s[i-1] == ')' {
			//上一个字符
			if s[i-2] == '(' {
				dp[i] = dp[i-2] + 2
			} else {
				//上一个字符,减去其最长有效的括号长度,得到最长有效括号的前一个字符a
				prefixIdx := i - 2 - dp[i-1]
				if prefixIdx >= 0 && s[prefixIdx] == '(' {
					//如果字符a为'(',则当前字符能构成的最大有效括号长度为:
					//上一个字符最大有效长度(dp[i-1]) + 2 + 字符a的上一个字符能构成的最大有效长度(dp[prefixIdx])
					dp[i] = dp[i-1] + 2 + dp[prefixIdx]
				}
			}
			if dp[i] > ans {
				ans = dp[i]
			}
		}
	}
	return ans
}
func longestValidParentheses1(s string) int {
	dp := make([]int, len(s))
	ans := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				prefixIdx := i - 1 - dp[i-1]
				if prefixIdx >= 0 && s[prefixIdx] == '(' {
					dp[i] = dp[i-1] + 2
					if prefixIdx > 1 {
						dp[i] += dp[prefixIdx-1]
					}
				}
			}
			if dp[i] > ans {
				ans = dp[i]
			}
		}
	}
	return ans
}

func longestValidParentheses(s string) int {
	maxAns := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}
