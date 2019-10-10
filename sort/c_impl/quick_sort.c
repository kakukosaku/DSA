//
// Created by kaku on 2019/10/7.
//

#include "array.h"

int partition(int *arr, int low, int high) {
    // Always select low as pivot
    int pivot_val = arr[low];

    while (low < high) {
        while (low < high && arr[high] >= pivot_val) high--;
        arr[low] = arr[high];
        while (low < high && arr[low] <= pivot_val) low++;
        arr[high] = arr[low];
    }

    arr[low] = pivot_val;
    return low;
}

void _quick_sort(int *arr, int low, int high) {
    if (low < high) {
        int pivot = partition(arr, low, high);

        _quick_sort(arr, low, pivot - 1);
        _quick_sort(arr, pivot + 1, high);
    }
}

void quick_sort(int *arr, int array_size) {
    _quick_sort(arr, 0, array_size - 1);
}