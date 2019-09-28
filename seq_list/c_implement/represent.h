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
#define True 1
#define False 0

// Bool Type Define
typedef int Bool;

// Sequence List Type Define
typedef struct {
    int *data;
    int length, capacity;
} SeqList;


void print_array(SeqList *list) {
    int i;
    printf("Array: have \"%d\" elements.\n", list->length);
    for (i = 0; i < list->length; i++) {
        printf("%d ", list->data[i]);
    }
    printf("\n");
}
