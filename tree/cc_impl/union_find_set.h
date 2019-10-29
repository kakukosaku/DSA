//
// Created by kaku on 2019/10/30.
//

#ifndef DSA_UNION_FIND_SET_H
#define DSA_UNION_FIND_SET_H

typedef int ElemType;

// Tree Parent Represent
typedef struct UFSNode {
    ElemType data;
    int parent_idx;
};
typedef struct UFSet {
    UFSNode *nodes;
    int count;
};

#endif //DSA_UNION_FIND_SET_H
