//
// Created by kaku on 2019/10/30.
//

#include <iostream>
#include "binary_search_tree.h"

using namespace std;

void show_tree_node(BSNode *n) {
    cout << "+++++++++++++++++++++++++++++++++++++++++\n";
    cout << "+\t\tNode.data(" << n->data << ")\t\t\t\t\t+\n";

    bool has_lchild = false;
    if (n->lchild) {
        has_lchild = true;
        cout << "+ Node.lchild(" << n->lchild->data << ")\t";
    }

    if (n->rchild) {
        if (not has_lchild) {
            cout << "+\t\t\t\t";
        }
        cout << "Node.rchild (" << n->rchild->data << ")";

        if (not has_lchild) {
            cout << "\t";
        }

        cout << "\t\t+\n";
    } else {
        cout << "\t\t\t\t\t+\n";
    }
    cout << "+++++++++++++++++++++++++++++++++++++++++\n\n";
}

void bst_init(BSTree &t, ElemType arr[], int arr_size) {
    t->data = arr[0];
    for (int i = 1; i < arr_size; i++) {
        bst_insert(t, arr[i]);
    }
}

BSNode *bst_search(BSTree t, ElemType k) {
    auto p = t;
    while (p && k != p->data) {
        if (k < p->data) {
            p = p->lchild;
        } else {
            p = p->rchild;
        };
    }

    return p;
}

bool bst_insert(BSTree &t, ElemType k) {
    if (!t) {
        t = new BSNode;
        t->data = k;
        t->lchild = t->rchild = nullptr;
        return true;
    } else if (k == t->data) {
        return false;
    } else if (k < t->data) {
        return bst_insert(t->lchild, k);
    } else {
        return bst_insert(t->rchild, k);
    }
}
