package stack


/**
栈解法
*/
func reverseParentheses(s string) string {
	stack := []byte{}
	str := []byte{}
	for i := range s {
		if len(stack) == 0 && s[i] != '(' {
			str = append(str, s[i])
		} else {
			if s[i] != ')' {
				stack = append(stack, s[i])
			} else {
				queue := []byte{}
				for len(stack) > 0 && stack[len(stack)-1] != '(' {
					queue = append(queue, stack[len(stack)-1])
					//等价于pop的写法
					stack = stack[:len(stack)-1]
				}
				//把左括号也一起pop
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					for len(queue) > 0 {
						str = append(str, queue[0])
						//出队
						queue = queue[1:]
					}
				} else {
					//重新入栈
					for len(queue) > 0 {
						stack = append(stack, queue[0])
						queue = queue[1:]
					}
				}
			}
		}
	}
	return string(str)
}
