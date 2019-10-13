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

from typing import List, NoReturn
from random import randint


def random_array(array_size: int = 10, array_range: tuple = (0, 100)) -> List[int]:
    """Get random array"""
    array = list()
    for _ in range(array_size):
        array.append(randint(*array_range))

    return array


def reversed_array(array_size: int = 10):
    """Get reversed array for test reason, like [5,4,3,2,1]"""
    array = list()
    for i in range(array_size):
        array.append(array_size - i)

    return array


def swap(array: List[int], i: int, j: int):
    """Swap i and j corresponding value

    Notes:
        1. Needn't pass pointer, since array pass by reference

    """
    array[i], array[j] = array[j], array[i]


def show_array(array: List[int], description: str) -> NoReturn:
    """Display Function"""
    print("\nDescription:", description, sep="\t")
    print("Show Array", array, sep="\t")


def check_sorted(array: List[int], reverse=False) -> bool:
    """Check array is sorted or not, default is from small to big"""
    for i in range(len(array) - 1):
        if reverse:
            sort = array[i] >= array[i + 1]
        else:
            sort = array[i] <= array[i + 1]

        if not sort:
            print("* Result:\tNot Sorted Array!")
            return False

    print("* Result:\tSorted Array!")
    return True
