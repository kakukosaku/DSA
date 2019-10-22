//
// Created by kaku on 2019/10/21.
//

#include <iostream>
#include "linked_list.h"

using namespace std;

void show_linked_list(LinkedList l) {
    LNode *n;
    cout << "Show Linked List:\n";
    cout << "\t";
    n = l->next;
    while (n) {
        cout << n->data << " -> ";
        n = n->next;
    }
    cout << "nullptr\n";
}

void show_node(LNode *n) {
    cout << "Show LNode:\n";
    cout << "\tdata: " << n->data;
    cout << "\tnext: -> " << n->next->data << endl;
}

LinkedList head_insert_linked_list(const ElemType arr[], int arr_size) {
    LNode *n;
    LinkedList l = new LNode;
    l->next = nullptr;

    for (int i = 0; i < arr_size; i++) {
        n = new LNode;
        n->data = arr[i];
        n->next = l->next;
        l->next = n;
    }

    return l;
}

LinkedList tail_insert_linked_list(const ElemType arr[], int arr_size) {
    LNode *n, *r;
    LinkedList l = new LNode;
    l->next = nullptr;
    r = l;

    for (int i = 0; i < arr_size; i++) {
        n = new LNode;
        n->data = arr[i];
        r->next = n;
        r = n;
    }

    r->next = nullptr;
    return l;
}

LNode *get_elem(LinkedList l, int pos) {
    int i = 1;
    if (pos == 0) {
        return l;
    }
    if (pos < 1) {
        return nullptr;
    }
    LNode *n = l->next;
    while (n && i < pos) {
        n = n->next;
        i++;
    }

    return n;
}

bool insert_elem(LinkedList l, LNode *n, int pos) {
    LNode *prior = get_elem(l, pos - 1);
    if (!prior) {
        return false;
    }

    n->next = prior->next;
    prior->next = n;

    return true;
}

bool delete_elem(LinkedList l, int pos) {
    // valid position
    LNode *p, *q;
    if (pos < 1) {
        return false;
    }
    p = get_elem(l, pos - 1);
    if (!p) {
        return false;
    }

    // delete operation
    q = p->next;
    p->next = q->next;
    delete q;

    return true;
}

void reverse_linked_list(LinkedList l) {
    // Notice: this implement has head node :)
    LNode *prior = nullptr, *curr = l->next, *next = nullptr;

    while (curr) {
        next = curr->next;
        curr->next = prior;

        prior = curr;
        curr = next;
    }
    l->next = prior;
}