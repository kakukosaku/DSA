//
// Created by kaku on 2019/10/7.
//

#include "array.h"

// you can define array parameter like this `int arr[]`.
void adjust_down(int arr[], int idx, int array_size) {
    int temp = arr[idx];

    for (int i = 2 * idx + 1; i < array_size; i = (2 * i + 1)) {
        if (i < array_size - 1 && arr[i] < arr[i + 1]) i++;  // find the bigger between two children.
        if (temp >= arr[i]) break;
        else {
            arr[idx] = arr[i];
            idx = i;
        }
    }

    // Important: find position & move for space & put in right!
    arr[idx] = temp;
}

void build_min_heap(int *arr, int array_size) {
    for (int i = array_size / 2; i >= 0; i--) {
        adjust_down(arr, i, array_size);
    }
}


// Correct Index is critical and it depend on you do what use index!
void heap_sort(int *arr, int array_size) {
    int temp;

    build_min_heap(arr, array_size);
    for (int i = array_size; i > 0; i--) {
        temp = arr[0];
        arr[0] = arr[i - 1];
        arr[i - 1] = temp;
        adjust_down(arr, 0, i - 1);
    }
}