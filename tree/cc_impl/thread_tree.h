//
// Created by kaku on 2019/10/29.
//

#ifndef DSA_THREAD_TREE_H
#define DSA_THREAD_TREE_H

typedef int ElemType;

typedef struct ThreadNode {
    ElemType data;
    struct ThreadNode *lchild, *rchild;
    int ltag, rtag;
} ThreadNode, *TreadTree;

#endif //DSA_THREAD_TREE_H
