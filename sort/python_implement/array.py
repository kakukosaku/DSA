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

from typing import List
from random import randint


def random_array(array_size: int = 10, array_range: tuple = (0, 100)) -> List[int]:
    """Get random array"""
    array = list()
    for _ in range(array_size):
        array.append(randint(*array_range))

    return array


def reversed_array(array_size: int = 10):
    array = list()
    for i in range(array_size):
        array.append(array_size - i)

    return array


def swap(array: list, i: int, j: int):
    """Swap i and j corresponding value

    Notes:
        1. Needn't pass pointer, since array pass by reference

    """
    array[i], array[j] = array[j], array[i]


def show_array(array, description):
    print("Description:", description, sep="\t")
    print("Show Array", array, sep="\t")
