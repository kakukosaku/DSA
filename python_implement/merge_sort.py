#!/usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 18/05/08

# GitHub:
#
#   https://github.com/kakuchange
#
# 归并排序 Python 实现, 转载请附原链接


from random import randint


def merge_sort(ls):
    """归并排序

    Args:
        ls:

    Returns:

    """
    def merge(lls, rls):
        new_ls = list()
        l_index = 0
        r_index = 0
        while len(new_ls) < len(lls) + len(rls):
            if l_index < len(lls) and lls[l_index] <= rls[r_index]:
                new_ls.append(lls[l_index])
                l_index += 1
                if l_index >= len(lls):
                    for tmp in rls[r_index:]:
                        new_ls.append(tmp)
                    break
                continue
            if r_index < len(rls) and rls[r_index] < lls[l_index]:
                new_ls.append(rls[r_index])
                r_index += 1
                if r_index >= len(rls):
                    for tmp in lls[l_index:]:
                        new_ls.append(tmp)
                    break
                continue

        return new_ls

    interrupt_point = len(ls) // 2
    if interrupt_point > 0:  # 注意, 并归不对最小组内大小做区分, 划分的子组最小应为[x]
        left_ls = merge_sort(ls[:interrupt_point])
        right_ls = merge_sort(ls[interrupt_point:])
        return merge(left_ls, right_ls)
    return ls


if __name__ == '__main__':
    ls = list()
    for _ in range(10):
        ls.append(randint(0, 100))
    # 完全反序用以模拟最差情况
    # ls = [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]
    print('old:')
    print(ls)
    sorted_ls = merge_sort(ls)
    print('sorted:')
    print(sorted_ls)
