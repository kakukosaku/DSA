#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 19/10/08
#
# GitHub:
#
#   https://github.com/kakukosaku
#
# © 2019-2022 Kaku Kosaku All Rights Reserved

from typing import List


def quick_sort_test_wrap(arr: List[int], array_size: int) -> List[int]:
    """for invoker test_case reason..."""
    return quick_sort(arr, 0, array_size - 1)


def quick_sort(arr: List[int], low: int, high: int) -> List[int]:
    """Quick Sort, c style :)

    Notes:
        1. Arguments pass by reference for mutable variables, needn't return arr.
        2. Pass array size to function although `can get array len on runtime in Python`.
        3. Use while loop instead for loop as for loop is not friendly to use index.

    """
    if low < high:
        pivot = partition(arr, low, high)
        quick_sort(arr, low, pivot - 1)
        quick_sort(arr, pivot + 1, high)
    
    return arr


def partition(arr: List[int], low: int, high: int) -> int:
    pivot_val = arr[low]
    while low < high:
        while low < high and arr[high] >= pivot_val:
            high -= 1
        arr[low] = arr[high]
        while low < high and arr[low] <= pivot_val:
            low += 1
        arr[high] = arr[low]

    arr[low] = pivot_val
    return low
