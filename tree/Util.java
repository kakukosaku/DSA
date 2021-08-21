import java.util.LinkedList;
import java.util.Queue;

public class Util {

    public static void show(TreeNode root) {
        if (root == null) {
            return;
        }

        StringBuilder sb = new StringBuilder();
        sb.append("Level show tree:\n");
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        int currLevel = 0;
        while (!queue.isEmpty()) {
            currLevel++;
            sb.append("\nLv.").append(currLevel).append("\t");
            for (int levelCnt = queue.size(); levelCnt > 0; levelCnt--) {
                TreeNode node = queue.remove();
                sb.append(node.val).append(" ");

                if (node.left != null) {
                    queue.add(node.left);
                }

                if (node.right != null) {
                    queue.add(node.right);
                }
            }
        }
        System.out.println(sb.toString());
    }

    public static TreeNode newCompleteTree(int[] arr) {
        if (arr.length == 0) {
            return null;
        }

        TreeNode root = new TreeNode(arr[0]);
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        for (int i = 1; i < arr.length; i++) {
            TreeNode currNode = queue.element();
            TreeNode newNode = new TreeNode(arr[i]);
            if (currNode.left == null) {
                currNode.left = newNode;
            } else {
                currNode.right = newNode;
                queue.remove();
            }
            queue.add(newNode);
        }

        return root;
    }

    public static void main(String[] args) {
        int[] arr = new int[] { 1, 2, 3, 4, 5, 6, 7, 8 };
        TreeNode root = newCompleteTree(arr);
        show(root);
    }

}
