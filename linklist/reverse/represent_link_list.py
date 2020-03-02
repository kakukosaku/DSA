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

from typing import List, Any


class LinkListNode:
    """单链表 节点 元素类"""

    def __init__(self, data):
        self.data = data
        self.next = None
    
    def __repr__(self):
        return "<Node: data=%s>" % self.data


class LinkList:

    def __init__(self):
        self.next = None

    def head_insert(self, datas: List[Any]):
        for data in datas:
            node = LinkListNode(data)
            node.next = self.next
            self.next = node

    def reverse(self):
        if not self.next:
            raise ValueError("link list is null")

        prev_node = None
        curr_node = self.next
        while curr_node:
            next_node = curr_node.next
            curr_node.next = prev_node
            prev_node = curr_node
            curr_node = next_node
        
        self.next = prev_node


    def travel(self):
        curr = self.next
        while curr:
            print(curr, end=" -> ")
            curr = curr.next
        print("end")


if __name__ == '__main__':
    datas = range(10)
    link_list_obj = LinkList()
    link_list_obj.head_insert(datas)
    link_list_obj.travel()
    link_list_obj.reverse()
    link_list_obj.travel()
