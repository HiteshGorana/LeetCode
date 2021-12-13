class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if not s:
            return 0
        output = 0
        for index, i in enumerate(s):
            l = 0
            for start in s[index:]:
                if i == start:
                    break
                else:
                    l+=1
            if l > output:
                output = l
        return output


sol = Solution()
print(sol.lengthOfLongestSubstring("pwwkew"))
