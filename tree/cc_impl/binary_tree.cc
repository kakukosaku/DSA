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
    cout << "Row vector:\n\t";
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

void pre_order_recursive(BiTree t) {
    if (t) {
        cout << t->data << " ";
        pre_order_recursive(t->lchild);
        pre_order_recursive(t->rchild);
    }
}

void pre_order(BiTree t) {
    vector<BiTNode *> stack;
    auto p = t;

    while (p || !stack.empty()) {
        if (p) {
            stack.insert(stack.end(), p);
            cout << p->data << " ";
            p = p->lchild;
        } else {
            p = stack[stack.size() - 1];
            stack.pop_back();
            p = p->rchild;
        }
    }
}

void in_order_recursive(BiTree t) {
    if (t) {
        in_order_recursive(t->lchild);
        cout << t->data << " ";
        in_order_recursive(t->rchild);
    }
}

void in_order(BiTree t) {
    vector<BiTNode *> stack;
    auto p = t;

    while (p || !stack.empty()) {
        if (p) {
            stack.insert(stack.end(), p);
            p = p->lchild;
        } else {
            p = stack[stack.size() - 1];
            stack.pop_back();
            cout << p->data << " ";
            p = p->rchild;
        }
    }
}

void post_order_recursive(BiTree t) {
    if (t) {
        post_order_recursive(t->lchild);
        post_order_recursive(t->rchild);
        cout << t->data << " ";
    }
}

void post_order(BiTree t) {
    vector<BiTNode *> first_queue, second_queue;
    BiTNode * n;

    first_queue.insert(first_queue.begin(), t);

    while (!first_queue.empty()) {
        n = first_queue[first_queue.size() - 1];
        first_queue.pop_back();
        second_queue.push_back(n);

        if (n->lchild) {
            first_queue.push_back(n->lchild);
        }

        if (n->rchild) {
            first_queue.push_back(n->rchild);
        }
    }

    while (!second_queue.empty()) {
        n = second_queue[second_queue.size() - 1];
        second_queue.pop_back();
        cout << n->data << " ";
    }
}

void level_order(BiTree t) {
    vector<BiTNode *> queue;
    BiTNode * n;

    n = t;
    queue.insert(queue.begin(), n);
    while (!queue.empty()) {
        n = queue[queue.size() - 1];
        cout << n->data << " ";
        queue.pop_back();

        if (n->lchild) {
            queue.insert(queue.begin(), n->lchild);
        }

        if (n->rchild) {
            queue.insert(queue.begin(), n->rchild);
        }
    }
}
