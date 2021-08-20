public class ReverseLinkedList {

    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    public static ListNode reverseList(ListNode head) {
        if (head == null) {
            return head;
        }

        ListNode currNode = head;
        ListNode prevNode = null;
        ListNode nextNode = null;
        while (currNode != null) {
            nextNode = currNode.next;
            currNode.next = prevNode;
            prevNode = currNode;
            currNode = nextNode;
        }

        return prevNode;
    }

    public static void showLinkedList(ListNode head) {
        StringBuilder sb = new StringBuilder();
        ListNode currNode = head;
        while (currNode != null) {
            sb.append(currNode.val + "-> ");
            currNode = currNode.next;
        }
        System.out.println(sb);
    }

    public static void main(String[] args) {
        ListNode node1 = new ListNode(1);
        ListNode node2 = new ListNode(2);
        ListNode node3 = new ListNode(3);
        node1.next = node2;
        node2.next = node3;
        showLinkedList(node1);

        ListNode reversedNode = reverseList(node1);
        showLinkedList(reversedNode);
    }

}
