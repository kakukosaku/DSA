#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 19/10/10

# GitHub:
#
#   https://github.com/kakukosaku
#
# © 2019-2022 Kaku Kosaku All Rights Reserved

from typing import List, NoReturn
from .array import swap


def adjust_down(arr: List[int], idx: int, array_size: int) -> NoReturn:
    the_ele = arr[idx]
    i = 2 * idx + 1  # left child
    while i < array_size:
        if i < array_size - 1 and arr[i] < arr[i + 1]:
            i += 1

        if the_ele > arr[i]:
            break
        else:
            arr[idx] = arr[i]
            idx = i

        i = 2 * i + 1

    # Find the position & move for place & put right position.
    arr[idx] = the_ele


def build_heap(arr: List[int], array_size: int) -> NoReturn:
    i = array_size // 2
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

    i = 1
    while i < array_size:
        swap(arr, 0, array_size - i)
        adjust_down(arr, 0, array_size - i)
        i += 1
