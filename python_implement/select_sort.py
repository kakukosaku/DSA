#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 18/05/08

# GitHub:
#
#   https://github.com/kakuchange
#
# 简单选择排序 Python 实现, 转载请附原链接


from random import randint


def select_sort(ls):
    """简单选择排序

    Args:
        ls:

    Returns:

    """
    ls = ls[:]
    total_compare = 0
    total_swap = 0
    for i in range(len(ls)):
        smaller_index = i
        # 内层循环确定的依据 *1.欲将有序元素放置哪里, 2.欲以何种方式使用索引比较*
        for j in range(i + 1, len(ls)):
            total_compare += 1
            if ls[smaller_index] > ls[j]:  # 比 *目前记录最小* 元素更小时, 记录
                smaller_index = j
        if smaller_index != i:
            ls[i], ls[smaller_index] = ls[smaller_index], ls[i]
            total_swap += 1
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
    sorted_ls = select_sort(ls)
    print('sorted:')
    print(sorted_ls)
