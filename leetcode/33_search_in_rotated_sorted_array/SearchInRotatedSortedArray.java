public class SearchInRotatedSortedArray {
    public int search(int[] nums, int target) {
        if ((nums.length == 0) || (nums.length == 1 && nums[0] != target)) {
            return -1;
        }

        int left = 0, right = nums.length - 1, mid = 0;
        while (left <= right) {
            mid = left + (right - left) / 2;
            if (nums[mid] == target) {
                return mid;
            }

            if (nums[mid] < nums[right]) {
                if (nums[mid] < target && target <= nums[right]) {
                    left = mid + 1;
                } else {
                    right = mid - 1;
                }
            } else {
                if (nums[mid] > target && target >= nums[left]) {
                    right = mid - 1;
                } else {
                    left = mid + 1;
                }
            }
        }

        return -1;
    }

    public static void main(String[] args) {
        int a = Integer.MAX_VALUE;
        int b = Integer.MIN_VALUE;
        System.out.println(new SearchInRotatedSortedArray().search(new int[] { 4, 1, 2, 3 }, 2));
    }
}
