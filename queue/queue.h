//
// Created by scugj on 10/23/2019.
//

#ifndef DSA_QUEUE_H
#define DSA_QUEUE_H

ElemType int

// Sequence List Representation
typedef struct {
    ElemType *data;
    int front, rear;
} SeqQueue;

// Linked List Representation
typedef struct QNode{
    ElemType data;
    struct QNode *next;
};

typedef struct {
    QNode *front, *rear;
} LinkedQueue;

#endif //DSA_QUEUE_H
