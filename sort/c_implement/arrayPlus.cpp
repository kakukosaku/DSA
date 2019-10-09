//
// Created by kaku on 2019/10/9.
//

#include <math>
#include <cstdlib>
#include <time.h>
#include "arrayPlus.h"


ArrayPlus::ArrayPlus(int array_size) : len(array_size), cap(array_size * 2) {
    array = new [array_size]int;

    // To generate random int each time call.
    srandom(time(nullptr));
    for (int i = 0; i < array_size; i++) {
        tmp = (int) random() % 100;
        *(arr + i) = tmp;
    }
}


ArrayPlus::~ArrayPlus() {
    delete array;
}