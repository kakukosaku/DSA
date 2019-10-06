//
// Created by kaku on 2019/10/7.
//

#include "array.h"

void select_sort(int *ptr, int array_size) {
    int max_val;
    int max_idx;
    int temp;

    for (int i = 0; i < array_size; i++) {
        max_val = ptr[0];
        max_idx = 0;
        for (int j = 1; j < array_size - i - 1; j++) {
            if (max_val < ptr[j]) {
                max_idx = j;
                max_val = ptr[j];
            }
        }

        temp = ptr[max_idx];
        ptr[max_idx] = ptr[array_size - i - 1];
        ptr[array_size - i - 1] = temp;
    }
}
