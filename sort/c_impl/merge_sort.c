//
// Created by kaku on 2019/10/10.
//

#include <stdlib.h>
#include "array.h"

void merge(int *arr, int low, int mid, int high, int *arr_tmp) {
    int i, j, k;
    for (i = low; i <= high; i++) {
        arr_tmp[i] = arr[i];
    }

    for (i = low, j = mid + 1, k = i; i <= mid && j <= high; k++) {
        if (arr_tmp[i] < arr_tmp[j]) arr[k] = arr_tmp[i++];
        else arr[k] = arr_tmp[j++];
    }

    while (i <= mid) arr[k++] = arr_tmp[i++];
    while (j <= high) arr[k++] = arr_tmp[j++];
}

void _merge_sort(int *arr, int low, int high, int *arr_tmp) {
    if (low < high) {
        int mid = (low + high) / 2;
        _merge_sort(arr, low, mid, arr_tmp);
        _merge_sort(arr, mid + 1, high, arr_tmp);
        merge(arr, low, mid, high, arr_tmp);
    }
}

void merge_sort(int *arr, int array_size) {
    int *arr_tmp = random_array(array_size);
    _merge_sort(arr, 0, array_size - 1, arr_tmp);
    free(arr_tmp);
}