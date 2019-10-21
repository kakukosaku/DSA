//
// Created by kaku on 2019/10/21.
//

#ifndef DSA_LINKED_LIST_H
#define DSA_LINKED_LIST_H

typedef int ElemType;

typedef struct LNode {
    ElemType data;
    struct LNode *next;
} LNode, *LinkedList;

void show_linked_list(LinkedList);

LinkedList head_insert_linked_list(const ElemType [], int);

LinkedList tail_insert_linked_list(const ElemType [], int);

#endif //DSA_LINKED_LIST_H
