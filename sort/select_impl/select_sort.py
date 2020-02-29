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
from utils.sort_utils import swap


def select_sort(arr: List[int], array_size: int) -> List[int]:
    """Select Sort, c style :)

    Notes:
        1. Arguments pass by reference for mutable variables, needn't return arr.
        2. Pass array size to function although `can get array len on runtime in Python`.
        3. Use while loop instead for loop as for loop is not friendly to use index.

    """
    sorted_num = 0
    while sorted_num < array_size:
        curr_idx = sorted_num
        small_idx = curr_idx
        check_idx = curr_idx + 1
        while check_idx < array_size:
            # from small to large
            if arr[check_idx] < arr[small_idx]:
                small_idx = check_idx
            check_idx += 1

        swap(arr, curr_idx, small_idx)
        sorted_num += 1
    
    return arr
