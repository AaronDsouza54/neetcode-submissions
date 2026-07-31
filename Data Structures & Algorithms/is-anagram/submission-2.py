class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        
        chars = {}

        for i in range(len(s)):
            if s[i] not in chars:
                chars[s[i]] = 1
            else:
                chars[s[i]] += 1
            
            if t[i] not in chars:
                chars[t[i]] = -1
            else:
                chars[t[i]] -= 1

        for key in chars.keys():
            if chars[key] != 0:
                return False
        return True