//
// Created by kaku on 2019/10/30.
//

#ifndef DSA_BINARY_SEARCH_TREE_H
#define DSA_BINARY_SEARCH_TREE_H

typedef int ElemType;

typedef struct BSNode {
    ElemType data;
    struct BSNode *lchild, *rchild;
} BSNode, *BSTree;

void show_tree_node(BSNode *);

void bst_init(BSTree &, ElemType [], int);

BSNode *bst_search(BSTree, ElemType);

bool bst_insert(BSTree &, ElemType);

#endif //DSA_BINARY_SEARCH_TREE_H
