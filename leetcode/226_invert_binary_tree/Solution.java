/*
 * author: kaku
 * date: 2020/月/日
 *
 * GitHub:
 *   https://github.com/kakukosaku
 *
 * Description:
 */
class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }
}

public class Solution {
    public void reverseNode(TreeNode node) {
        if (node == null) {
            return;
        }
        TreeNode tmp = node.left;
        node.left = node.right;
        node.right = tmp;
        if (node.left != null) {
            reverseNode(node.left);
        }

        if (node.right != null) {
            reverseNode(node.right);
        }
    }

    public TreeNode invertTree(TreeNode root) {
        reverseNode(root);
        return root;
    }
}