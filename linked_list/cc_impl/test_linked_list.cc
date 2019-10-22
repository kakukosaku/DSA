//
// Created by kaku on 2019/10/21.
//

#include "linked_list.h"
#include <iostream>

using namespace std;

int main() {
    ElemType arr[] = {1, 2, 3, 4, 5};
    LinkedList l = head_insert_linked_list(arr, 5);
    show_linked_list(l);

    l = tail_insert_linked_list(arr, 5);
    show_linked_list(l);

    LNode *n = get_elem(l, 3);
    show_node(n);

    LNode n1 = {9, nullptr};
    insert_ele(l, &n1, 3);
    show_linked_list(l);
}