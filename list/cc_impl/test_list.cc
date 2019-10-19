//
// Created by kaku on 2019/10/10.
//

#include <iostream>
#include "list.h"

using namespace std;

int main() {
    ElemType a[] = {0, 1, 2, 3, 4};
    List l{a, 5};
    cout << l << endl;

    l.Insert(9, 1);
    cout << l << endl;
}