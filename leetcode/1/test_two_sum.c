//
// Created by kaku on 2019/10/24.
//

#include <stdio.h>
#include "two_sum.c"

int main() {
    int returnSize = 2;
    int arr[] = {2,7,11,15};
    int *rest;

    rest = twoSum(arr, 4, 18, &returnSize);
    printf("[%d, %d]", rest[0], rest[1]);
}