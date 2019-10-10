//
// Created by kaku on 2019/10/6.
//

#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "array.h"

int *random_array(int array_size) {
    int tmp;
    int *arr = (int *) malloc(array_size * sizeof(int));

    // To generate random int each time call.
    srandom(time(NULL));
    for (int i = 0; i < array_size; i++) {
        tmp = (int) random() % 100;
        *(arr + i) = tmp;
    }

    return arr;
}

int *copy_array(const int *arr, int array_size) {
    int *new_arr = (int *) malloc(array_size * sizeof(int));
    for (int i = 0; i < array_size; i++) {
        new_arr[i] = arr[i];
    }

    return new_arr;
}


void show_array(int *arr, int array_size, const char *description) {
    printf("\nDescription:\t%s\nShow Array:\t", description);
    printf("[");
    for (int i = 0; i < array_size; i++) {
        if (i == (array_size - 1)) {
            printf("%d]\n", arr[i]);
        } else {
            printf("%d, ", arr[i]);
        }
    }
}

Bool check_sorted(const int arr[], int array_size, Bool reverse) {
    Bool sort;
    for (int i = 0; i < array_size - 1; i++) {
        if (reverse) {
            sort = (arr[i] >= arr[i + 1]);
        } else {
            sort = (arr[i] <= arr[i + 1]);
        }

        if (sort != True) {
            printf("* Result: Not Sorted Array!\n");
            return False;
        }
    }

    printf("* Result: Sorted Array\n");
    return True;
}
