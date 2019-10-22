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

void show_node(LNode *);

LinkedList head_insert_linked_list(const ElemType [], int);

LinkedList tail_insert_linked_list(const ElemType [], int);

LNode *get_elem(LinkedList, int);

bool insert_elem(LinkedList, LNode *, int);

bool delete_elem(LinkedList, int);

void reverse_linked_list(LinkedList);

#endif //DSA_LINKED_LIST_H
