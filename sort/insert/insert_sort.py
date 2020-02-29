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

from typing import List


def insert_sort(arr: List[int], array_size: int) -> List[int]:
    """Insert Sort, c style :)

    Notes:
        1. Arguments pass by reference for mutable variables, needn't return arr.
        2. Pass array size to function although `can get array len on runtime in Python`.
        3. Use while loop instead for loop as for loop is not friendly to use index.

        4. Find the Right position; Move for space; Put it in.

    """
    sorted_num = 0
    while sorted_num < array_size:
        check_idx = sorted_num + 1
        # from small to large
        if check_idx < array_size and arr[check_idx] < arr[check_idx - 1]:
            temp = arr[check_idx]
            while check_idx > 0 and arr[check_idx - 1] > temp:
                # move check_elem to next position.
                arr[check_idx] = arr[check_idx - 1]
                check_idx -= 1
            arr[check_idx] = temp

        sorted_num += 1
    return arr
