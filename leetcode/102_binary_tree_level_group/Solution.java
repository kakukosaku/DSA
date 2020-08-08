/*
 * author: kaku
 * date: 2019/月/日
 *
 * GitHub:
 *   https://github.com/kakukosaku
 *
 * Description:
 */
import java.util.*;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int val) {
        this.val = val;
    }

}

public class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> lists = new ArrayList<>();
        if (root == null) {
            return lists;
        }

        List<TreeNode> queueNodes = new ArrayList<>();
        queueNodes.add(root);
        while (!queueNodes.isEmpty()) {
            int levelSize = queueNodes.size();
            ArrayList<Integer> list = new ArrayList<>();

            for (int i = 0; i < levelSize; i++) {
                TreeNode node = queueNodes.remove(0);
                list.add(node.val);

                if (node.left != null) {
                    queueNodes.add(node.left);
                }

                if (node.right != null) {
                    queueNodes.add(node.right);
                }
            }
            lists.add(list);
        }
        return lists;
    }

    public static void main(String[] args) {
        System.out.println("Please run in leetcode");
    }
}

