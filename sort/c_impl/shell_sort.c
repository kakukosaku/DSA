//
// Created by kaku on 2019/10/13.
//

#include "array.h"

void shell_sort(int *arr, int array_size) {
    int i, j, temp;
    for (int increment = array_size / 2; increment > 0; increment = increment / 2) {
        for (i = increment; i < array_size; i++) {
            temp = arr[i];
            if (arr[i] < arr[i - increment]) {
                for (j = i - increment; j >= 0 && arr[j] > temp; j -= increment) {
                    arr[j + increment] = arr[j];
                }
                arr[j + increment] = temp;
            }
        }
    }
}
