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
        self.root_node = None
        self.waiting_to_full_leave = list()

    def complete_tree_add(self, node):
        """create a complete binary tree(完全二叉树)"""
        if not isinstance(node, Node):
            node = Node(node)
        if not self.root_node:
            self.root_node = node
            self.waiting_to_full_leave.append(node)
        else:
            not_full_node = self.waiting_to_full_leave[0]
            if not not_full_node.lchild:
                not_full_node.lchild = node
                self.waiting_to_full_leave.append(node)
            else:
                not_full_node.rchild = node
                self.waiting_to_full_leave.append(node)
                self.waiting_to_full_leave.pop(0)

    def pre_order_iteration_recursive(self, node):
        if not node:
            print('\t Node arrive end!')
            return
        print('\t', node.data)
        self.pre_order_iteration_recursive(node.lchild)
        self.pre_order_iteration_recursive(node.rchild)

    def pre_order_iteration_stack(self):
        if self.root_node is None:
            return
        stack = list()
        node = self.root_node
        while node or stack:
            while node:  # 从根节点开始，一直找它的左子树
                print('\t', node.data)
                stack.append(node)
                node = node.lchild
            print('\t Node arrive end!')
            node = stack.pop()  # while结束表示当前节点node为空，即前一个节点没有左子树了
            node = node.rchild

    def in_order_iteration_recursive(self, node):
        if not node:
            print('\t Node arrive end!')
            return
        self.in_order_iteration_recursive(node.lchild)
        print('\t', node.data)
        self.in_order_iteration_recursive(node.rchild)

    def in_order_iteration_stack(self):
        if self.root_node is None:
            return
        stack = list()
        node = self.root_node
        while node or stack:
            while node:  # 先将所有左子树压栈
                stack.append(node)
                node = node.lchild
                print('\t Current Node arrive end!')
            node = stack.pop()
            print('\t', node.data)
            node = node.rchild

    def post_order_iteration_recursive(self, node):
        if not node:
            print('\t Node arrive end!')
            return
        self.post_order_iteration_recursive(node.lchild)
        self.post_order_iteration_recursive(node.rchild)
        print('\t', node.data)

    def post_order_iteration_stack(self):
        if self.root_node is None:
            return
        node = self.root_node
        stack1 = list()
        stack2 = list()
        stack1.append(node)
        while stack1:
            node = stack1.pop()
            if node.lchild:
                stack1.append(node.lchild)
            if node.rchild:
                stack1.append(node.rchild)
            stack2.append(node)
        while stack2:
            print(stack2.pop().data)


if __name__ == "__main__":
    tree = Tree()
    for i in range(10):
        tree.complete_tree_add(i)
    print("""完全二叉树:
                0
        1               2
    3       4       5       6
 7     8 9     ...
          """)
    print('前序遍历:')
    print('0 1 3 7 8 4 9 2 5 6')
    print('recursive solution:')
    tree.pre_order_iteration_recursive(tree.root_node)
    print('stack solution:')
    tree.pre_order_iteration_stack()
    print('中序遍历:')
    print('7 3 8 1 9 4 0 5 2 6')
    print('recursive solution:')
    tree.in_order_iteration_recursive(tree.root_node)
    print('stack solution:')
    tree.in_order_iteration_stack()
    print('后序遍历:')
    print('7 8 3 9 4 1 5 6 2 0')
    print('recursive solution:')
    tree.post_order_iteration_recursive(tree.root_node)
    print('stack solution:')
    tree.post_order_iteration_stack()
