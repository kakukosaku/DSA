//
// Created by kaku on 2019/10/9.
//

#include <iostream>
#include "list.h"

using std::cout;
using std::ostream;

// Constructor Function
List::List(ElemType arr_[], int array_size_) : len(array_size_), cap(array_size_ * 2) {
    array = new ElemType[cap];
    for (int i = 0; i < array_size_; i++) {
        array[i] = arr_[i];
    }
}

// Destructor Function
List::~List() {
    delete array;
}

int List::Length() {
    return len;
}

void List::reallocate() {
    cap *= 2;
    ElemType *new_array = new ElemType[cap];

    for (int i = 0; i < len; i++) {
        new_array[i] = array[i];
    }

    delete array;
    array = new_array;
}

bool List::Insert(ElemType ele, int pos) {
    // Check `pos`, `pos` means index will to insert!
    if (pos < 0 || pos > len + 1) {
        return false;
    }

    // Re-allocate memory if necessary.
    if (len + 1 > cap) {
        reallocate();
    }

    // Insert
    for (int i = len; i >= pos; i--) {
        array[i] = array[i - 1];
    }
    array[pos] = ele;
    len += 1;

    return true;
}

// Reload operator <<
ostream &operator<<(ostream &os, const List &arr) {
    // How to use c style char string and cc style string.
    char cstr[100];
    sprintf(cstr, "Show List: len: %d, cap: %d\n", arr.len, arr.cap);
    string ccstr(cstr);

    cout << ccstr;
    cout << "\t[";
    for (int i = 0; i < arr.len; i++) {
        cout << arr.array[i] << " ";
    }
    cout << "]\n";
    return os;
}
