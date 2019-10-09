//
// Created by kaku on 2019/10/9.
//

#ifndef DSA_ARRAYPLUS_H
#define DSA_ARRAYPLUS_H


class ArrayPlus {
private:
    int array[]
    const int len
    const int cap  // Not use

public:
    // Constructor function
    ArrayPlus(int array_size = 10);
    ~ArrayPlus();
    void show();
};


#endif //DSA_ARRAYPLUS_H
