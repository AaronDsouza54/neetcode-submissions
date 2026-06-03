func isValid(s string) bool {

	if len(s) % 2 == 1 {
		return false
	}

	stack := make([]byte, 0, len(s))
	
	for i := 0; i < len(s); i++ {
		switch s[i] {
			case '(': 
				stack = append(stack, ')')
			case '{':
				stack = append(stack, '}')
			case '[':
				stack = append(stack, ']')
			default:
				n := len(stack)

				if n == 0 || stack[n-1] != s[i] {
					return false
				}

				stack = stack[0:n-1]
		}
	}

	return len(stack) == 0
}
