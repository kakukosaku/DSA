#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 19/10/07

# GitHub:
#
#   https://github.com/kakukosaku
#
# © 2019-2022 Kaku Kosaku All Rights Reserved
#
# Usage:
#   python -m python_implement.bubble_sort

from .array import random_array, show_array, swap

ARRAY_SIZE = 10


def bubble_sort(arr, array_size):
    """Optimized Bubble Sort

    Notes:
        1. If some circle not move elements position means array is sorted!
        2. In Python arguments pass by reference to mutable variables, needn't return arr.
        3. Replace for loop with while since Python for loop is not friendly to use index.
        4. Pass array size to function `even in Python can get array(list) len on runtime`.

    """
    i = 0
    moved = True
    while i < array_size:
        if moved:
            j = 0
            moved = False
            while j < array_size - i - 1:
                # from small to large
                if arr[j] > arr[j + 1]:
                    swap(arr, j, j + 1)
                    moved = True
                j += 1
        else:
            break


if __name__ == '__main__':
    arr = random_array(ARRAY_SIZE)
    show_array(arr, "Original Array")
    bubble_sort(arr, len(arr))
    show_array(arr, "After Bubble Sort")
