//
// Created by kaku on 2019/10/9.
//

#include <cmath>
#include <cstdlib>
#include <ctime>
#include <iostream>
#include "array.h"

using std::cout;
using std::ostream;

// Constructor Function
Array::Array(int array_size): len(array_size), cap(2*array_size){
    cout << "Array Initializing ...\n";
    array = new int[cap];

    // To generate random int each time call.
    srandom(time(nullptr));
    int tmp;
    for (int i = 0; i < array_size; i++) {
        tmp = (int) random() % 100;
        *(array + i) = tmp;
    }
}

// Destructor Function
Array::~Array() {
    cout << "Array Destroying ...\n";
    delete [] array;
}

// Reload operator <<
ostream & operator<<(ostream & os, const Array &arr) {
    cout << "Show Array: [";
    for (int i = 0; i < arr.len; i++) {
        cout << arr.array[i] << " ";
    }
    cout << "]\n";
    return os;
}
