//
// Created by kaku on 2019/10/7.
//

#include "array.h"

void select_sort(int *arr, int array_size) {
    int max_val;
    int max_idx;
    int temp;

    for (int i = 0; i < array_size; i++) {
        max_val = arr[0];
        max_idx = 0;
        for (int j = 1; j < array_size - i; j++) {
            if (max_val < arr[j]) {
                max_idx = j;
                max_val = arr[j];
            }
        }

        temp = arr[max_idx];
        arr[max_idx] = arr[array_size - i - 1];
        arr[array_size - i - 1] = temp;
    }
}
