//
// Created by kaku on 2019/10/30.
//

#include "binary_search_tree.h"
#include <iostream>

using namespace std;

int main() {
    ElemType arr[] = {5, 1, 2, 3, 4, 7, 8, 6};
    auto t = new BSNode;
    bst_init(t, arr, 8);
    show_tree_node(t);
    show_tree_node(t->lchild);
    show_tree_node(t->lchild->rchild);
    show_tree_node(t->rchild);

    auto n = bst_search(t, 7);
    cout << "Search Result:\n";
    show_tree_node(n);
}