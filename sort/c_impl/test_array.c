//
// Created by kaku on 2019/10/6.
//

#include <stdlib.h>
#include "array.h"

#define ARRAY_SIZE 10

int main() {
    int *ptr;

    ptr = random_array(ARRAY_SIZE);
    show_array(ptr, ARRAY_SIZE, "Test Case.");
    free(ptr);

    return 0;
}