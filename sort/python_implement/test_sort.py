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
#   python -m python_implement.test_sort

from typing import List, NoReturn, Callable

from .array import random_array, show_array, check_sorted
from . import bubble_sort, insert_sort, select_sort

ARRAY_SIZE = 10


def test_case(arr: List[int], array_size: int, description: str, sort_func: Callable) -> NoReturn:
    _arr = list(arr)
    sort_func(_arr, array_size)
    show_array(_arr, description)
    check_sorted(_arr)


def main():
    arr = random_array(ARRAY_SIZE)
    show_array(arr, "Original Array")

    test_case(arr, ARRAY_SIZE, "After Bubble Sort", bubble_sort)
    test_case(arr, ARRAY_SIZE, "After Insert Sort", insert_sort)
    test_case(arr, ARRAY_SIZE, "After Select Sort", select_sort)


if __name__ == '__main__':
    main()
