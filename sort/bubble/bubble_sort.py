#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 19/10/07
#
# GitHub:
#
#   https://github.com/kakukosaku
#
# © 2019-2022 Kaku Kosaku All Rights Reserved

from utils.sort_utils import swap
from typing import List, NoReturn


def bubble_sort(arr: List[int], array_size: int) -> List[int]:
    """Optimized Bubble Sort, c style :)

    Notes:
        1. Arguments pass by reference for mutable variables, needn't return arr.
        2. Pass array size to function although `can get array len on runtime in Python`.
        3. Use while loop instead for loop as for loop is not friendly to use index.

        4. If one loop not move elements position means array is sorted!

    """
    sorted_num = 0
    moved = True
    while sorted_num < array_size:
        if moved:
            check_idx = 0
            moved = False
            while check_idx < array_size - sorted_num - 1:
                # from small to large
                if arr[check_idx] > arr[check_idx + 1]:
                    swap(arr, check_idx, check_idx + 1)
                    moved = True
                check_idx += 1
            sorted_num += 1
        else:
            break

    return arr