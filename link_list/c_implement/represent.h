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
#define Bool int
#define True 1
#define False 0


typedef struct LNode{
    int data;
    struct LNode *next;
} LNode, *LinkList;


void print_link_list(&LinkList) {
    printf("Link List show as:");
    int p;
    LNode *current_node;
    current_node = LinkList->LNode;
    while (current_node) {
        p = current_node->data;
        current_node = current_node->next;
        if (current_node)
            printf("%d -> ", current_node->data);
        else
            printf("%d\n", current_node->data);
    }
}
