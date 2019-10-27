//
// Created by kaku on 2019/10/27.
//

#include "binary_tree.h"

using namespace std;

template<class T>
void print_vector(const vector<T> &v) {
    for (auto x: v) {
        cout << ' ' << x;
    }
    cout << '\n';
}

void show_tree_node(BiTNode *n) {
    cout << "\t\tNode.data(" << n->data << ")\n";

    if (n->lchild) {
        cout << "Node.lchild(" << n->lchild->data << ")\t";
    }

    if (n->rchild) {
        cout << "Node.rchild(" << n->rchild->data << ")\n";
    } else {
        cout << "\n";
    }
}

BiTree init_tree_from_array(vector<int> v) {
    print_vector(v);
    int d;
    vector<BiTNode *> tmp_v;
    BiTNode *tmp_n, *unfull_node;

    // init root
    d = v[v.size() - 1];
    v.pop_back();
    tmp_n = new BiTNode(d);
    BiTree t = tmp_n;
    tmp_v.insert(tmp_v.begin(), tmp_n);

    while (!v.empty()) {
        // Get a non-full Node
        unfull_node = tmp_v[tmp_v.size() - 1];
        if (unfull_node->rchild) {
            tmp_v.pop_back();
            unfull_node = tmp_v[tmp_v.size() - 1];
        }

        // Get an element and Init Node
        d = v[v.size() - 1];
        v.pop_back();
        tmp_n = new BiTNode(d);

        // Find where to insert
        if (!unfull_node->lchild) {
            unfull_node->lchild = tmp_n;
        } else {
            unfull_node->rchild = tmp_n;
        }
        tmp_v.insert(tmp_v.begin(), tmp_n);
    }

    return t;
}
