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


def ensure_ele_decorator(fun):
    """decorator for reduce redundancy code"""
    def inner(self, data):
        if not isinstance(data, LinkListEle):
            data = LinkListEle(data)
        return fun(self, data)
    return inner


class LinkListEle:
    """单链表 节点 元素类"""

    def __init__(self, data):
        self.data = data
        self.next = None

    def __repr__(self):
        return 'data: %s ' \
               'next=>%s' % (self.data, self.next)


class LinkList:

    def __init__(self):
        self.head = None
        self.current_ele = None

    @ensure_ele_decorator
    def add(self, ele):
        if self.head is None:
            self.head = ele
            self.current_ele = self.head
        else:
            self.current_ele.next = ele
            self.current_ele = ele

    def reverse(self):
        if not self.head or not self.head.next:
            return self.head
        last = None
        head = self.head
        while head:
            tmp = head.next
            head.next = last
            last = head
            head = tmp
        self.head = last

    def get_all_ele(self):
        cursor = self.head
        while True:
            print(cursor.data)
            if not cursor.next:
                break
            else:
                cursor = cursor.next

    def fake_present(self):
        print(self.head)


if __name__ == '__main__':
    ll = LinkList()
    for i in range(0, 5):
        ll.add(i)
    # print(ll.head)
    ll.get_all_ele()
    ll.reverse()
    print('反转链表后')
    ll.get_all_ele()
