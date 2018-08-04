#! /usr/bin/env python3
# coding: utf-8
#
# author: kaku
# date: 18/08/02
#
# GitHub:
#
#   https://github.com/kakuchange
#
# Description:
#
#   How to representation a Tree, and its iteration.


class Node:

    def __init__(self, data=None):
        self.data = data
        self.lchild = None
        self.rchild = None


class Tree:

    def __init__(self):
        self.root_node = Node()
        self.leaves = list()
