//
// Created by kaku on 2019/10/27.
//

#include "binary_tree.h"


int main() {
    vector<int> v = {8, 7, 6, 5, 4, 3, 2, 1};
    auto t = init_tree_from_array(v);
    show_tree_node(t);

    show_tree_node(t->lchild);
    show_tree_node(t->rchild);
    show_tree_node(t->lchild->lchild);
}