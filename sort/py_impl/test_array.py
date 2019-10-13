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
#
# Usage:
#   python -m py_impl.test_array

from .array import random_array, show_array, check_sorted

ARRAY_SIZE = 10


def test():
    arr = random_array(ARRAY_SIZE, (0, 100))
    show_array(arr, "Test Random Case")
    check_sorted(arr)
    arr = list(range(ARRAY_SIZE))
    show_array(arr, "Test Sorted Case")
    check_sorted(arr)


if __name__ == '__main__':
    test()
