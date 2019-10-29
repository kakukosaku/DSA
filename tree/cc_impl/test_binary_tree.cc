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

    cout << "\nPre Order:\n\t";
    pre_order_recursive(t);
    cout << "\nPre Order non-recursive:\n\t";
    pre_order(t);

    cout << "\nIn Order:\n\t";
    in_order_recursive(t);
    cout << "\nIn Order non-recursive:\n\t";
    in_order(t);

    cout << "\nPost Order:\n\t";
    post_order_recursive(t);
    cout << "\nPost Order non-recursive:\n\t";
    post_order(t);

    cout << "\nLevel Order:\n\t";
    level_order(t);
}