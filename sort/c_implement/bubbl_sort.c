//
// Created by kaku on 2019/10/6.
//

#include "array.h"

// Notice changed usage for optimize bubble sort!
// if some circle not change elements position means array is sorted!
void bubble_sort(int *ptr, int array_size) {
    int temp;
    Bool changed = True;

    for (int i = 0; i < array_size && changed; i++) {
        changed = False;
        for (int j = 0; j < (array_size - i - 1); j++) {
            // from small to large
            if (ptr[j] > ptr[j + 1]) {
                temp = ptr[j];
                ptr[j] = ptr[j + 1];
                ptr[j + 1] = temp;
                changed = True;
                // show_array(ptr, ARRAY_SIZE, "Changing...");
            }
        }
    }
}
