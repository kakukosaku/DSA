//
// Created by kaku on 2019/10/6.
//

#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "array.h"


int *init_array(int array_size) {
    int tmp;
    int *ptr = malloc(array_size * sizeof(int));

    // To generate random int each time call.
    srandom(time(NULL));
    for (int i = 0; i < array_size; i++) {
        tmp = (int)random() % 100;
        *(ptr + i) = tmp;
    }

    return ptr;
}


int *copy_array(int *ptr, int array_size) {
    int *new_ptr = malloc(array_size * sizeof(int));
    for (int i = 0; i < array_size; i++) {
        new_ptr[i] = ptr[i];
    }

    return new_ptr;
}


void show_array(int *ptr, int array_size, char * description) {
    printf("Description:\t%s\nShow Array:\t", description);
    printf("[");
    for (int i = 0; i < array_size; i++) {
        if (i == (array_size - 1)) {
            printf("%d]\n", ptr[i]);
        } else {
            printf("%d, ", ptr[i]);
        }
    }
}
