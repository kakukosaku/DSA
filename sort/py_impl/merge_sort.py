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


def merge(arr: List[int], low: int, mid: int, high: int, arr_tmp: List[int]) -> NoReturn:
    for i in range(low, high + 1):
        arr_tmp[i] = arr[i]

    i, j = low, mid + 1
    k = i
    while i <= mid and j <= high:
        # from small to lager
        if arr_tmp[i] <= arr_tmp[j]:
            arr[k] = arr_tmp[i]
            i += 1
        else:
            arr[k] = arr_tmp[j]
            j += 1

        k += 1

    while i <= mid:
        arr[k] = arr_tmp[i]
        i += 1
        k += 1

    while j <= high:
        arr[k] = arr_tmp[j]
        j += 1
        k += 1


def _merge_sort(arr: List[int], low: int, high: int, arr_tmp: List[int]) -> NoReturn:
    """merge sort, c style :)

    Notes:
        1. In Python arguments pass by reference to mutable variables, needn't return arr.
        2. Pass array size to function `even in Python can get array(list) len on runtime`.
        3. Replace for loop with while since Python for loop is not friendly to use index.

    """
    if low < high:
        mid = (low + high) // 2
        _merge_sort(arr, low, mid, arr_tmp)
        _merge_sort(arr, mid + 1, high, arr_tmp)
        merge(arr, low, mid, high, arr_tmp)


def merge_sort(arr: List[int], array_size: int):
    # arr_tmp avoid `new` list object each time
    arr_tmp = list(arr)
    _merge_sort(arr, 0, array_size - 1, arr_tmp)
