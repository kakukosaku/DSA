//
// Created by kaku on 2019/10/9.
//

#ifndef DSA_CPP_LIST_H
#define DSA_CPP_LIST_H

#include <iostream>

using namespace std;
typedef int ElemType;

class List {
private:
    int *array;
    int len;
    int cap;  // Not use

public:
    // Constructor function
    explicit List(ElemType *arr_, int array_size_);

    ~List();

    int Length();

    void reallocate();

    int LocateElem(ElemType);

    int GetElem(int);

    bool Insert(ElemType, int);

    int Delete(int);

    bool Empty();

    friend ostream &operator<<(ostream &os, const List &arr);
};

#endif //DSA_CPP_LIST_H
