//
//  filename.c
//  clangSkill
//
//  Created by kaku on 2019/04/22.
//  Copyright © 2019 kaku. All rights reserved.
//

/*
 * Description:
 * 对长度为n的顺序表, 编写一个时间复杂度为O(n), 空间复杂度为O(1)的算法,该算法
 * 删除线性表中所有值为x的数据元素.
 *
 * Key Word:
 * 移动删除
 */

#include "represent_array.h"

#define ArraySize 10


void del_x_1(int x, SeqList *list) {
    int i, k = 0;
    for (i = 0; i < list->length; i++) {
        if (list->data[i] != x) {
            list->data[k] = list->data[i];
            k++;
        } else {
            list->length--;
        }
    }
}


void del_x_2(int x, SeqList *list) {
    int i = 0, k = 0;
    while (i < list->length) {
        if (list->data[i] == x)
            // 为了保持写书风格, 实际上我认为大括号风格更友好些 :)
            k++, list->length--;
        else
            list->data[i-k] = list->data[i];
        i++;
    }
}


int main() {
    int a[ArraySize] = {1, 2, 3, 4, 5, 4, 3, 2, 1, 0};
    SeqList list = {a, ArraySize, ArraySize};
    print_array(&list);

    int x = 1;
    PRINT_INT(x);
    del_x_1(x, &list);
    print_array(&list);
}
