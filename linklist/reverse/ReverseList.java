/*
 * author: kaku
 * date: 2020/08/09
 *
 * GitHub:
 *   https://github.com/kakukosaku
 *
 * Description:
 */
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

public class ReverseList {

    public static ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }

        ListNode currNode = head;
        ListNode prevNode = null;
        while (currNode != null) {
            ListNode tmpNode = currNode.next;
            currNode.next = prevNode;
            prevNode = currNode;
            currNode = tmpNode;
        }

        return prevNode;
    }

}