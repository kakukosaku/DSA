#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 18/05/08

# GitHub:
#
#   https://github.com/kakuchange
#
# 快排 Python 实现, 转载请附原链接


from random import randint


def quick_sort(low, high, ls):
    """快排

    Args:
        low:
        high:
        ls:

    Returns:

    """
    def partitioin(low, high, ls):
        pivot_value = ls[high]
        while low < high:
            while low < high and ls[low] <= pivot_value:
                low += 1
            ls[low], ls[high] = ls[high], ls[low]
            while low < high and ls[high] > pivot_value:
                high -= 1
            ls[low], ls[high] = ls[high], ls[low]
        return high

    pivot = partitioin(low, high, ls)
    if low < high:
        quick_sort(low, pivot - 1, ls)
        quick_sort(pivot + 1, high, ls)
    return ls


# def quick_sort(arr, first, last):
#     """ Quicksort
#         Complexity: best O(n log(n)) avg O(n log(n)), worst O(N^2)
#     """
#     if first < last:
#         pos = partition(arr, first, last)
#         # Start our two recursive calls
#         quick_sort(arr, first, pos-1)
#         quick_sort(arr, pos+1, last)
#     return arr
#
#
# def partition(arr, first, last):
#     wall = first
#     for pos in range(first, last):
#         if arr[pos] < arr[last]: # last is the pivot
#             arr[pos], arr[wall] = arr[wall], arr[pos]
#             wall += 1
#     arr[wall], arr[last] = arr[last], arr[wall]
#     return wall


if __name__ == '__main__':
    ls = list()
    for _ in range(10):
        ls.append(randint(0, 100))
    # 完全反序用以模拟最差情况
    # ls = [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]
    ls = [3, 5, 4, 10, 2, 1, 9, 8, 7, 6]
    print('old:')
    print(ls)
    sorted_ls = quick_sort(0, len(ls) - 1, ls)
    # sorted_ls = quick_sort(ls, 0, len(ls) - 1)
    print('sorted:')
    print(sorted_ls)
