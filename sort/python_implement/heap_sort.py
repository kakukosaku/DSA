#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 18/05/08

# GitHub:
#
#   https://github.com/kakuchange
#
# 简单选择排序 Python 实现, 转载请附原链接


from random import randint


def heap_sort(lst):
    """堆排序 借鉴维基百科

    https://zh.wikipedia.org/wiki/%E5%A0%86%E6%8E%92%E5%BA%8F#Python

    Notes:

        1. 三元素中取最小, 如何防越界
        2. 时刻关注要排*元素*以及要将之放在哪里?

    """
    # 包含初始堆化与再堆化
    total_compare = 0
    total_swap = 0

    def heapnify(start, end):
        """堆化"""
        nonlocal total_compare, total_swap
        root = start
        while True:
            child = 2 * root + 1
            if child > end:
                break
            total_compare += 1
            if child + 1 <= end and lst[child] < lst[child + 1]:
                child += 1
            if lst[root] < lst[child]:
                total_swap += 1
                lst[root], lst[child] = lst[child], lst[root]
                root = child
            else:
                break

    for i in range((len(lst) - 2) // 2, -1, -1):
        heapnify(i, len(ls) - 1)

    for end in range(len(lst) - 1, 0, -1):
        lst[0], lst[end] = lst[end], lst[0]
        heapnify(0, end - 1)

    print('\ntotal compare:', total_compare, end='\n')
    print('\ntotal swap:', total_swap, end='\n\n')
    return lst


if __name__ == '__main__':
    ls = list()
    for _ in range(10):
        ls.append(randint(0, 100))
    # 完全反序用以模拟最差情况
    # ls = [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]
    print('old:')
    print(ls)
    sorted_ls = heap_sort(ls)
    print('sorted:')
    print(sorted_ls)
