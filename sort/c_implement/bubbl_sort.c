//
// Created by kaku on 2019/10/6.
//

#include "array.h"

// Notice changed usage for optimize bubble sort!
// If some circle not move elements position means array is sorted!
void bubble_sort(int *arr, int array_size) {
    int temp;
    Bool moved = True;

    for (int i = 0; i < array_size && moved; i++) {
        moved = False;
        for (int j = 0; j < (array_size - i - 1); j++) {
            // from small to large
            if (arr[j] > arr[j + 1]) {
                temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
                moved = True;
                // show_array(arr, ARRAY_SIZE, "Changing...");
            }
        }
    }
}
