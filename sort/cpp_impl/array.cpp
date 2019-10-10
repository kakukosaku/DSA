//
// Created by kaku on 2019/10/9.
//

#include <cmath>
#include <cstdlib>
#include <ctime>
#include "../c_impl/array.h"


// Constructor Function
Array::ArrayPlus(int array_size): len(array_size), cap(2*array_size){
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
Array::~ArrayPlus() {
    delete [] array;
}

