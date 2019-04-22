//
//  filename.c
//  clangSkill
//
//  Created by kaku on 2019/xx/xx.
//  Copyright © 2019 kaku. All rights reserved.
//

/*
 * 二分查找的一般形式
 *
 * KeyWrod:
 * 有序顺序表的查找
 */

#include <stdio.h>
#include "represent_array.h"

#define ArraySize 10

int binary_search(int x, SeqList *list);


int main() {
    int a[ArraySize] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
    SeqList list = {a, ArraySize, ArraySize};
    print_array(&list);

    int i, x = 15;
    PRINT_INT(x);
    i = binary_search(x, &list);
    printf("The index of x: %d\n", i);
}


int binary_search(int x, SeqList *list) {
    int mid, low = 0, high = list->length;

    while (low <= high) {
        mid = (low + high) / 2;
        if (list->data[mid] == x) {
            break;
        } else if (list->data[mid] < x) {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }

    if (low < high) {
        return mid;
    } else {
        return -1;
    }
}

