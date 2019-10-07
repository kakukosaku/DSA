//
// Created by kaku on 2019/10/7.
//

#include "array.h"

int partition(int *arr, int low, int high) {
    // Always select low as pivot
    int pivot = low;
    int pivot_val = arr[pivot];

    for (int i = low; i < high; i++) {
        if (arr[i] < pivot_val) {

        }
    }
}

void quick_sort(int *arr, int low, int high) {
    if (low < high) {
        int pivot;
        pivot = partition(arr, low, high);

        quick_sort(arr, low, pivot);
        quick_sort(arr, pivot, high);
    }
}