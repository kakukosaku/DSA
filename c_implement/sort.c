//
//  filename.c
//  clangSkill
//
//  Created by kaku on 2019/xx/xx.
//  Copyright © 2019 kaku. All rights reserved.
//

#include <stdio.h>

#define PRINT_INT(i) printf(#i ": %d\n", i)
#define PRINT_CHR(c) printf(#c ": %c\n", c)
#define PRINT_STR(s) printf(#s ": %s\n", s)

#define True 1;
#define False 1;
#define array_length 10


void bubble_sort(int length, int array[]);
void print_array(int length, int array[]);


int main() {
    int unsort_array[array_length] = {5, 4, 3, 2, 1, 6, 7, 8, 9, 0};
    print_array(array_length, unsort_array);

    bubble_sort(array_length, unsort_array);
};


// print func
void print_array(int length, int array[]) {
    int i;
    printf("Array is:\t");
    for (i = 0; i < length; i++) {
        printf("%d ", *(array+i));
    }
    printf("\n");
}


// Bubble sort
void bubble_sort(int length, int array[]) {
    typedef int Bool;
    int i, j, tmp;
    Bool change = True;
    for (i = length - 1; i > 0 && change; i--) {
        change = False;
        for (j = 0; j < i; j++) {
            if (array[j] < array[j+1]) {
                tmp = array[j];
                array[j] = array[j+1];
                array[j + 1] = tmp;
                change = True;
            }
        }
    }
    print_array(array_length, array);
}
