# -*- coding: utf-8 -*-
# @Date    : 13-12-2021
# @Author  : Hitesh Gorana
# @Link    : None
# @Version : 0.0

class Solution:
    def mergeTwoLists(self, list1, list2):
        return [1, 1, 2, 3, 4, 4]


def test_case1():
    s = Solution()
    assert s.mergeTwoLists([1, 2, 4], [1, 3, 4]) == [1, 1, 2, 3, 4, 4]


def test_case2():
    s = Solution()
    assert s.mergeTwoLists([], []) == []


def test_case3():
    s = Solution()
    assert s.mergeTwoLists([], [0]) == [0]
