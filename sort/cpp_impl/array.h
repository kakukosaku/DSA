//
// Created by kaku on 2019/10/9.
//

#ifndef DSA_CPP_ARRAY_H
#define DSA_CPP_ARRAY_H

#include <iostream>

using namespace std;

class Array {
private:

    int *array;
    const int len;
    const int cap;  // Not use

public:
    // Constructor function
    ArrayPlus(int array_size = 10);

    ~ArrayPlus();

    friend ostream &operator<<(ostream &os, const ArrayPlus &arr);
};

#endif //DSA_CPP_ARRAY_H
