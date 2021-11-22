class Solution:
    def majorityElement(self, nums):
        _count = {}
        for i in nums:
            if _count.get(i):
                _count[i] += 1
            else:
                _count[i] = 1
        return max(_count, key=_count.get)


# AFTER SUBMISSION 
# Runtime: 172 ms, faster than 52.14% of Python3 online submissions for Majority Element.
# Memory Usage: 15.4 MB, less than 82.70% of Python3 online submissions for Majority Element.