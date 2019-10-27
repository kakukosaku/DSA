//
// Created by kaku on 2019/10/27.
//

#ifndef DSA_BINARY_TREE_H
#define DSA_BINARY_TREE_H

#include <vector>
#include <iostream>

using namespace std;

typedef int ElemType;

typedef struct BiTNode {
    ElemType data;
    struct BiTNode *lchild, *rchild;

    BiTNode(ElemType d) : data(d), lchild(nullptr), rchild(nullptr) {}
} BiTNode, *BiTree;

BiTree init_tree_from_array(vector<int>);

void show_tree_node(BiTNode *);

#endif //DSA_BINARY_TREE_H
