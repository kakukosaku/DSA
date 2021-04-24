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

public class RemoveNthNode {
    /**
     * 删除链表中倒数第N个节点
     * @param head, 无头节点的单链表的头指针
     * @param n
     * @return
     */
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode fastNode, slowNode, dummyNode;
        dummyNode = new ListNode(0, head);
        fastNode = head;
        slowNode = dummyNode;
        for (int i = 0; i < n; i++) {
            fastNode = fastNode.next;
        }

        while (fastNode != null) {
            fastNode = fastNode.next;
            slowNode = slowNode.next;
        }

        slowNode.next = slowNode.next.next;
        return dummyNode.next;
    }
}
