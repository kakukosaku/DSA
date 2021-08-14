import java.util.Arrays;

public class RotateArray {

    public void reverse(int[] nums, int start, int end) {
        int temp;
        while (start < end) {
            temp = nums[start];
            nums[start] = nums[end];
            nums[end] = temp;
            start++;
            end--;
        }
    }

    public void rotate(int[] nums, int k) {
        k %= nums.length;  // 很重要, 对于输入k值的思考!
        reverse(nums, 0, nums.length - 1);
        reverse(nums, 0, k - 1);
        reverse(nums, k, nums.length - 1);
    }

    public static void main(String[] args) {
        int[] nums = new int[] { 1, 2 };
        int k = 4;
        // new RotateArray().rotate(nums, k);
        System.out.println(Arrays.toString(nums));
        System.out.println(k%nums.length);
    }
}
