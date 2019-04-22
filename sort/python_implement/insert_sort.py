#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 18/05/08

# GitHub:
#
#   https://github.com/kakuchange
#
# 直接插入排序 Python 实现, 转载请附原链接


from random import randint


def insert_sort(ls):
    """直接插入排序

    Notes:
        1. 类比抓扑克时的操作 => 向有序序列中插入新元素
        2. 外层循环 for i in range(len(ls))


    Args:
        ls:

    Returns:

    """
    ls = ls[:]
    total_compare = 0
    total_swap = 0
    for i in range(len(ls)):
        for j in range(i, 0, -1):
            total_compare += 1
            if ls[j - 1] > ls[j]:  # 当前被比元素较小时才交换, 从小到大
                total_swap += 1
                ls[j - 1], ls[j] = ls[j], ls[j - 1]
                continue  # 为了更清楚的表示, 向有序列表中插入的逻辑!
            else:
                break  # 向已有序列表中插入,
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
    sorted_ls = insert_sort(ls)
    print('sorted:')
    print(sorted_ls)
