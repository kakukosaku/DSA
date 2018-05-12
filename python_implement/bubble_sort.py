#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 18/05/08

# GitHub:
#
#   https://github.com/kakuchange
#
# 冒泡排序 Python 实现, 转载请附原链接

from random import randint


def bubble_sort(ls):
    """冒泡排序

    pass

    Args:
        ls: 待排列表

    Returns:
        sorted_ls

    """
    ls = ls[:]
    total_compare = 0
    total_swap = 0
    need_sort = True
    for i in range(len(ls)):
        if need_sort:
            for j in range(len(ls) - 1 - i):
                total_compare += 1
                if ls[j] > ls[j + 1]:  # 从小到大排
                    total_swap += 1
                    ls[j], ls[j + 1] = ls[j + 1], ls[j]
                    need_sort = True
                else:
                    need_sort = False
        else:
            print('一般情况: 外层循环 *未全部* 走完即有序')
            break
    else:
        print('最坏情况: 外层循环 *全部* 走完')
    print('\ntotal compare:', total_compare, end='\n')
    print('\ntotal swap:', total_swap, end='\n\n')
    return ls


if __name__ == '__main__':
    ls = list()
    for _ in range(10):
        ls.append(randint(0, 100))
    # 完全反序用以模拟最差情况
    # ls = [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]
    print('old:')
    print(ls)
    sorted_ls = bubble_sort(ls)
    print('sorted:')
    print(sorted_ls)
