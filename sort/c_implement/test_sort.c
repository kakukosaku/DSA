//
// Created by kaku on 2019/10/7.
//

#include <stdlib.h>
#include "array.h"

#define ARRAY_SIZE 10

void test_case(int * ptr, int array_size, char *description, void (*func)(int*, int)) {
    int * temp_ptr;
    temp_ptr = copy_array(ptr, ARRAY_SIZE);
    func(temp_ptr, array_size);
    show_array(temp_ptr, array_size, description);
    free(temp_ptr);
}

int main() {
    int *ptr;
    ptr = random_array(ARRAY_SIZE);
    show_array(ptr, ARRAY_SIZE, "Original Array.");

    test_case(ptr, ARRAY_SIZE, "After Bubble sort.", bubble_sort);
    test_case(ptr, ARRAY_SIZE, "After Select sort.", select_sort);
    test_case(ptr, ARRAY_SIZE, "After Insert sort.", insert_sort);
    test_case(ptr, ARRAY_SIZE, "After Quick sort.", quick_sort);
    test_case(ptr, ARRAY_SIZE, "After Heap sort.", heap_sort);
}

