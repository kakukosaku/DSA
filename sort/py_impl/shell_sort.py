#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 19/10/12

# GitHub:
#
#   https://github.com/kakukosaku
#
# © 2019-2022 Kaku Kosaku All Rights Reserved

from typing import List, NoReturn


def shell_sort(arr: List[int], array_size: int) -> NoReturn:
    increment = array_size // 2
    while increment >= 1:

        i = increment
        while i < array_size:
            # from small to large
            if arr[i] < arr[i - increment]:
                tmp = arr[i]
                j = i - increment
                while j >= 0 and tmp < arr[j]:
                    arr[j + increment] = arr[j]
                    j -= increment

                arr[j + increment] = tmp

            i += 1

        increment = increment // 2

