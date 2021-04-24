class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
        this.val = val;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}

public class MergeTwoSortedLists {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if (l1 == null || l2 == null) {
            return l1 != null ? l1 : l2;
        }
        ListNode newList, newCurrNode, l1CurrNode = l1, l2CurrNode = l2;
        if (l1.val > l2.val) {
            newList = l2;
            l2CurrNode = l2CurrNode.next;
        } else {
            newList = l1;
            l1CurrNode = l1CurrNode.next;
        }
        newCurrNode = newList;

        while (l1CurrNode != null && l2CurrNode != null) {
            if (l1CurrNode.val > l2CurrNode.val) {
                newCurrNode.next = l2CurrNode;
                l2CurrNode = l2CurrNode.next;
                newCurrNode = newCurrNode.next;
                continue;
            }

            newCurrNode.next = l1CurrNode;
            l1CurrNode = l1CurrNode.next;
            newCurrNode = newCurrNode.next;
        }

        if (l1CurrNode != null) {
            newCurrNode.next = l1CurrNode;
        }

        if (l2CurrNode != null) {
            newCurrNode.next = l2CurrNode;
        }

        return newList;
    }
}
