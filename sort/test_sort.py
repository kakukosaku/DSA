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
#   python -m py_impl.test_sort

from typing import List, NoReturn, Callable

from utils.sort_utils import random_array, show_array, is_sorted
from bubble.bubble_sort import bubble_sort
from insert.insert_sort import insert_sort
from select_impl.select_sort import select_sort
from heap_impl.heap_sort import heap_sort
from quick.quick_sort import quick_sort_test_wrap

ARRAY_SIZE = 10


def test_case(arr: List[int], array_size: int, description: str, sort_func: Callable) -> NoReturn:
    _arr = list(arr)
    _arr = sort_func(_arr, array_size)
    show_array(_arr, description)
    is_sorted(_arr, False)


def main():
    arr = random_array(ARRAY_SIZE)
    show_array(arr, "Original Array")
    test_case(arr, ARRAY_SIZE, "After bubble Sort", bubble_sort)
    test_case(arr, ARRAY_SIZE, "After insert Sort", insert_sort)
    test_case(arr, ARRAY_SIZE, "After select Sort", select_sort)
    test_case(arr, ARRAY_SIZE, "After heap Sort", heap_sort)
    test_case(arr, ARRAY_SIZE, "After quick Sort", quick_sort_test_wrap)


if __name__ == '__main__':
    """Usage: 
    $PWD: $BASE_PATH/DSA/sort
        python test_sort.py
    """
    main()
