#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 19/10/10
#
# GitHub:
#
#   https://github.com/kakukosaku
#
# © 2019-2022 Kaku Kosaku All Rights Reserved
from typing import List, NoReturn
from utils.sort_utils import swap


def adjust_down(arr: List[int], check_idx: int, array_size: int) -> NoReturn:
    the_ele = arr[check_idx]

    while check_idx * 2 + 1 < array_size:
        l_child_idx = (2 * check_idx) + 1
        r_child_idx = (2 * check_idx) + 2
        bigger_child_idx = l_child_idx

        if r_child_idx < array_size and arr[l_child_idx] < arr[r_child_idx]:
            bigger_child_idx = r_child_idx

        if the_ele > arr[bigger_child_idx]:
            break
        else:
            arr[check_idx] = arr[bigger_child_idx]
            check_idx = bigger_child_idx

    # Find the position & move for place & put right position.
    arr[check_idx] = the_ele


def build_heap(arr: List[int], array_size: int) -> NoReturn:
    i = (array_size // 2) - 1
    while i >= 0:
        adjust_down(arr, i, array_size)
        i -= 1


def heap_sort(arr: List[int], array_size: int) -> NoReturn:
    """Heap sort, c style :)

    Notes:
        1. In Python arguments pass by reference to mutable variables, needn't return arr.
        2. Pass array size to function `even in Python can get array(list) len on runtime`.
        3. Replace for loop with while since Python for loop is not friendly to use index.

        4. build the heap, remote the first element, heapify again.

    """
    build_heap(arr, array_size)

    sorted_num = 1
    while sorted_num < array_size:
        swap(arr, 0, array_size - sorted_num)
        adjust_down(arr, 0, array_size - sorted_num)
        sorted_num += 1

    return arr
