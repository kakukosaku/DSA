import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Stack;

public class Travel {

    public static int[] preOrder(TreeNode root) {
        List<Integer> travelRest = new ArrayList<>();
        if (root == null) {
            return convertListToArray(travelRest);
        }

        Stack<TreeNode> stack = new Stack<>();
        stack.push(root);
        TreeNode currNode;
        while (!stack.isEmpty()) {
            currNode = stack.pop();
            travelRest.add(currNode.val);

            if (currNode.right != null) {
                stack.push(currNode.right);
            }

            if (currNode.left != null) {
                stack.push(currNode.left);
            }
        }

        return convertListToArray(travelRest);
    }

    public static int[] inOrder(TreeNode root) {
        List<Integer> travelRest = new ArrayList<>();
        if (root == null) {
            return convertListToArray(travelRest);
        }

        Stack<TreeNode> stack = new Stack<>();
        TreeNode currNode = root;
        while (currNode != null) {
            stack.push(currNode);
            currNode = currNode.left;
        }

        while (!stack.isEmpty()) {
            currNode = stack.pop();
            travelRest.add(currNode.val);

            if (currNode.right != null) {
                currNode = currNode.right;
                while (currNode != null) {
                    stack.push(currNode);
                    currNode = currNode.left;
                }
            }
        }

        return convertListToArray(travelRest);
    }

    public static int[] postOrder(TreeNode root) {
        List<Integer> travelRest = new ArrayList<>();
        if (root == null) {
            return convertListToArray(travelRest);
        }

        Stack<TreeNode> stack1 = new Stack<>();
        Stack<TreeNode> stack2 = new Stack<>();
        stack1.push(root);
        TreeNode currNode;
        while (!stack1.isEmpty()) {
            currNode = stack1.pop();
            stack2.push(currNode);
            if (currNode.left != null) {
                stack1.push(currNode.left);
            }

            if (currNode.right != null) {
                stack1.push(currNode.right);
            }
        }

        while (!stack2.isEmpty()) {
            currNode = stack2.pop();
            travelRest.add(currNode.val);
        }

        return convertListToArray(travelRest);
    }

    public static int[] convertListToArray(List<Integer> input) {
        int[] arr = new int[input.size()];
        for (int i = 0; i < input.size(); i++) {
            arr[i] = input.get(i);
        }
        return arr;
    }

    public static void main(String[] args) {
        int[] arr = new int[] { 1, 2, 3, 4, 5, 6, 7, 8 };
        TreeNode root = Util.newCompleteTree(arr);
        Util.show(root);

        int[] travelRest;
        travelRest = preOrder(root);
        System.out.println("Pre order:\t" + Arrays.toString(travelRest));

        travelRest = inOrder(root);
        System.out.println("In order:\t" + Arrays.toString(travelRest));

        travelRest = postOrder(root);
        System.out.println("Post order:\t" + Arrays.toString(travelRest));
    }
}
