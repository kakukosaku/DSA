//
// Created by kaku on 2019/10/7.
//

#include <stdlib.h>
#include "array.h"

#define ARRAY_SIZE 10

void test_case(int *arr, int array_size, const char description[], void (*func)(int *, int)) {
    int *temp_arr;
    temp_arr = copy_array(arr, ARRAY_SIZE);
    func(temp_arr, array_size);
    show_array(temp_arr, array_size, description);
    check_sorted(temp_arr, array_size, False);
    free(temp_arr);
}

int main() {
    int *arr;
    arr = random_array(ARRAY_SIZE);
    show_array(arr, ARRAY_SIZE, "Original Array");

    test_case(arr, ARRAY_SIZE, "After Bubble Sort", bubble_sort);
    test_case(arr, ARRAY_SIZE, "After Select Sort", select_sort);
    test_case(arr, ARRAY_SIZE, "After Insert Sort", insert_sort);
    test_case(arr, ARRAY_SIZE, "After Quick Sort", quick_sort);
    test_case(arr, ARRAY_SIZE, "After Heap Sort", heap_sort);
}

