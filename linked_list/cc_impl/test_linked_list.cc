//
// Created by kaku on 2019/10/21.
//

#include "linked_list.h"

int main() {
    ElemType arr[] = {1, 2, 3, 4, 5};
    LinkedList l = head_insert_linked_list(arr, 5);
    show_linked_list(l);

    l = tail_insert_linked_list(arr, 5);
    show_linked_list(l);
}