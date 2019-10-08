//
// Created by kaku on 2019/10/6.
//

#ifndef DSA_ARRAY_H
#define DSA_ARRAY_H
#define True 1
#define False 0

typedef int Bool;

// Init a pointer to array
int *random_array(int);

// Copy a array, return a new pointer to a new array
int *copy_array(const int *, int);

// Show array by pointer pointed to array
void show_array(int *, int, char *);

// Check array is sorted or not
Bool check_sorted(const int *, int, Bool);

// Bubble Sort
void bubble_sort(int *, int);

// Quick Sort
void quick_sort(int *, int);

// Select Sort
void select_sort(int *, int);

// Heap Sort
void heap_sort(int *, int);

// Insert Sort
void insert_sort(int *, int);

#endif //DSA_ARRAY_H
