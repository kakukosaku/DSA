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

from .array import swap
from typing import List, NoReturn


def bubble_sort(arr: List[int], array_size: int) -> NoReturn:
    """Optimized Bubble Sort, c style :)

    Notes:
        1. In Python arguments pass by reference to mutable variables, needn't return arr.
        2. Pass array size to function `even in Python can get array(list) len on runtime`.
        3. Replace for loop with while since Python for loop is not friendly to use index.

        4. If some circle not move elements position means array is sorted!

    """
    i = 0
    moved = True
    while i < array_size:
        if moved:
            j = 0
            moved = False
            while j < array_size - i - 1:
                # from small to big
                if arr[j] > arr[j + 1]:
                    swap(arr, j, j + 1)
                    moved = True
                j += 1
        else:
            break
