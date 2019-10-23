//
// Created by scugj on 10/23/2019.
//

#ifndef DSA_STACK_H
#define DSA_STACK_H

typedef int ElemType;

// Sequence List Representation
typedef struct {
    ElemType *data;
    int bottom;
    int top;
} SeqStack;

// Linked List Representation
typedef struct SNode {
    ElemType data;
    struct SNode *next, *bottom, *top;
}LinkedStack;

#endif //DSA_STACK_H
