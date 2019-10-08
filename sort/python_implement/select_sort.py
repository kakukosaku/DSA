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
from .array import swap


def select_sort(arr: List[int], array_size: int) -> NoReturn:
    """Select Sort

    Notes:
        1. In Python arguments pass by reference to mutable variables, needn't return arr.
        2. Pass array size to function `even in Python can get array(list) len on runtime`.
        3. Replace for loop with while since Python for loop is not friendly to use index.

        4. Find the Right position; Move for space; Put it in.

    """
    i = 0
    while i < array_size:
        min_idx = i
        j = i + 1
        while j < array_size:
            # from small to big
            if arr[j] < arr[min_idx]:
                min_idx = j
            j += 1

        swap(arr, i, min_idx)
        i += 1
