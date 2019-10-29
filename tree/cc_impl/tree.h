//
// Created by kaku on 2019/10/29.
//

#ifndef DSA_TREE_H
#define DSA_TREE_H

typedef int ElemType;

// Parent Representation
typedef struct {
    ElemType data;
    int parent;
} PTNode;


typedef struct {
    PTNode *nodes;
    int n;
} PTree;

// Child Representation
typedef struct CTNode {
    ElemType data;
    struct CTNode *next;
} CTNode;

typedef struct {
    CTNode *nodes;
    int n;
} CTree;

// Child Sibling Representation
typedef struct CSNode {
    ElemType data;
    struct CSNode *first_chile, *next_sibling;
} CSNode, *CSTree;

#endif //DSA_TREE_H
