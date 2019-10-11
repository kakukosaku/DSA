#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 19/10/08

# GitHub:
#
#   https://github.com/kakukosaku
#
# © 2019-2022 Kaku Kosaku All Rights Reserved

from typing import List, NoReturn


def quick_sort(arr: List[int], array_size: int) -> NoReturn:
    _quick_sort(arr, 0, array_size - 1)


def _quick_sort(arr: List[int], low: int, high: int) -> NoReturn:
    """Quick Sort, c style :)

    Notes:
        1. In Python arguments pass by reference to mutable variables, needn't return arr.
        2. Pass array size to function `even in Python can get array(list) len on runtime`.
        3. Replace for loop with while since Python for loop is not friendly to use index.

    """
    if low < high:
        pivot = partition(arr, low, high)
        _quick_sort(arr, low, pivot - 1)
        _quick_sort(arr, pivot + 1, high)


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
