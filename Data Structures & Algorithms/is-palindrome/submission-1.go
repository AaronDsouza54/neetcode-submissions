func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	
	if len(s) <= 1 {
        return true
    }

    for start < end {
	    char1 := s[start]
	    char2 := s[end]
	    
	    if !isAcceptedChar(char1) {
	        start++
	        continue
	    } else if !isAcceptedChar(char2) {
	        end--
	        continue
	    }
	    
	    if char1 >= 'A' && char1 <= 'Z' {
	        char1 += 32
	    } 
	    
	    if char2 >= 'A' && char2 <= 'Z' {
	        char2 += 32
	    }
	    
	    if char1 != char2 {
	        return false
	    }
	    
	    start++
	    end--
	    
	} 
	
	return true
}

func isAcceptedChar(char uint8) bool {
    return (char >= '0' && char <= '9') || (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')
}