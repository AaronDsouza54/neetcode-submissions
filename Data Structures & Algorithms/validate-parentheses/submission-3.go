func isValid(s string) bool {
	oppositeParen := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	if len(s) % 2 == 1 {
		return false
	}

	middleChar := len(s) / 2
	
    stack := NewStack(middleChar)


	for i := 0; i < len(s); i++ {
		_, exists := oppositeParen[s[i]]
		
		fmt.Printf("Opposite paren: %c\n", oppositeParen[s[i]])
		if !exists {
			if stack.IsEmpty() || oppositeParen[stack.Pop()] != s[i] {
			    return false
			}
		} else {
		    stack.Push(s[i])
		}
	}

	return stack.Sp == -1
}

type Stack struct {
    Arr         []byte
    Sp          int
}

func NewStack(capacity int) *Stack {
    temp := make([]byte, capacity)
    return &Stack {
        temp,
        -1,
    }
}

func (s *Stack) IsEmpty() bool {
    return s.Sp == -1    
}

func (s *Stack) Push(value byte) {
    s.Sp++
    if s.Sp == len(s.Arr) {
        s.Arr = append(s.Arr, value)
    } else {
        s.Arr[s.Sp] = value
    }
}

func (s *Stack) Pop() byte {
    temp := s.Arr[s.Sp]
    s.Sp--
    return temp
}