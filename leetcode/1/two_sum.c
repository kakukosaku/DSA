//
// Created by kaku on 2019/10/24.
//

#include <stdlib.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *twoSum(int *nums, int numsSize, int target, int *returnSize) {
    int *rest = (int *) malloc(*returnSize * (int) sizeof(int));
    for (int i = 0; i < numsSize; i++) {
        for (int j = i + 1; j < numsSize; j++) {
            if (nums[i] + nums[j] == target) {
                rest[0] = i;
                rest[1] = j;
                return rest;
            } else {
                continue;
            }
        }
    }
    return rest;
}
