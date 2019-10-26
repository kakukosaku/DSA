//
// Created by kaku on 2019/10/25.
//

// Definition for singly-linked list.

#include <iostream>

using namespace std;


struct ListNode {
    int val;
    ListNode *next;

    ListNode(int x) : val(x), next(nullptr) {}
};

void show_linked_list(ListNode *l) {
    ListNode *n;
    cout << "Show Linked List:\n";
    cout << "\t";
    cout << l->val << " -> ";
    n = l->next;
    while (n) {
        cout << n->val << " -> ";
        n = n->next;
    }
    cout << "nullptr\n";
}

class Solution {
public:
    ListNode *addTwoNumbers(ListNode *l1, ListNode *l2) {
        auto _l1n = new ListNode(0);
        _l1n->next = l1;
        l1 = _l1n;
        auto _l2n = new ListNode(0);
        _l2n->next = l2;
        l2 = _l2n;
        int tmp, rest;
        int extra = 0;
        auto *l3 = new ListNode(0);
        auto l3_head = l3;
        auto l1t = l1->next;
        auto l2t = l2->next;
        while (l1t && l2t) {
            tmp = l1t->val + l2t->val + extra;
            if (tmp >= 10) {
                extra = 1;
                rest = tmp - 10;
            } else {
                extra = 0;
                rest = tmp;
            }

            auto new_node = new ListNode(rest);
            l3->next = new_node;
            l3 = l3->next;

            l1t = l1t->next;
            l2t = l2t->next;
        }

        while (l1t) {
            if (extra) {
                l1t->val += extra;
                extra = 0;
            }

            if (l1t->val >= 10) {
                l1t->val = l1t->val - 10;
                extra = 1;
            }

            l3->next = l1t;
            l3 = l3->next;
            l1t = l1t->next;
        }

        while (l2t) {
            if (extra) {
                l2t->val += extra;
                extra = 0;
            }

            if (l2t->val >= 10) {
                l2t->val = l2t->val - 10;
                extra = 1;
            }

            l3->next = l2t;
            l3 = l3->next;
            l2t = l2t->next;
        }

        if (extra != 0) {
            l3->next = new ListNode(1);
        }

        return l3_head->next;
    }
};


int main() {
    auto l1 = new ListNode(1);
//    l1->next = new ListNode(4);
//    l1->next->next = new ListNode(6);
//    l1->next->next->next = new ListNode(3);
//    l1->next->next->next->next = new ListNode(0);
//    l1->next->next->next->next->next = new ListNode(5);
    auto l2 = new ListNode(9);
    l2->next = new ListNode(9);
//    l2->next->next = new ListNode(4);
//    l2->next->next->next = new ListNode(4);
//    l2->next->next->next->next = new ListNode(4);
//    l2->next->next->next->next->next = new ListNode(5);

    show_linked_list(l1);
    show_linked_list(l2);

    Solution s;
    auto _l3 = s.addTwoNumbers(l1, l2);
    show_linked_list(_l3);
}