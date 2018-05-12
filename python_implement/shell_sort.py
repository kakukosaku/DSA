#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 18/05/08

# GitHub:
#
#   https://github.com/kakuchange
#
# 希尔排序 Python 实现, 转载请附原链接


from random import randint


def shell_sort(ls):
    """希尔排序
    Args:
        ls:

    Returns:

    """
    ls = ls[:]
    total_compare = 0
    total_swap = 0
    increment = len(ls)
    increment = increment // 2
    while increment > 0:
        for j in range(increment, len(ls)):
            for k in range(j, 0, -increment):
                # 此处有小坑, ls[-1] 是会取到值的, 最后一轮的直接插入, 简单粗暴的从结果上
                # 掩盖了问题, 只是后面看每种算法的比较, 交换次数时, 才发现怎么这么多?, if 的判断是必须的
                if k - increment > 0:
                    total_compare += 1
                    if ls[k] < ls[k - increment]:
                        total_swap += 1
                        ls[k], ls[k - increment] = ls[k - increment], ls[k]
                        continue  # 原因同前插入排序
                    else:
                        break
        increment = increment // 2
    print('\ntotal compare:', total_compare, end='\n')
    print('\ntotal swap:', total_swap, end='\n\n')
    return ls


if __name__ == '__main__':
    ls = list()
    for _ in range(10):
        ls.append(randint(0, 100))
    # 完全反序用以模拟最差情况
    ls = [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]
    print('old:')
    print(ls)
    sorted_ls = shell_sort(ls)
    print('sorted:')
    print(sorted_ls)
