class FindMinInRotatedSortedArray {

    public int findMin(int[] nums) {
        if (nums.length == 1) {
            return nums[0];
        }

        if (nums[0] < nums[nums.length - 1]) {
            return nums[0];
        }

        int left = 0, right = nums.length - 1, mid = 0;
        while (left <= right) {
            mid = left + (right - left) / 2;

            if (nums[mid] > nums[mid + 1]) {
                return nums[mid + 1];
            }

            if (nums[mid] < nums[mid - 1]) {
                return nums[mid];
            }

            if (nums[mid] > nums[right]) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return -1;
    }

    public static void main(String[] args) {
        System.out.println(new FindMinInRotatedSortedArray().findMin(new int[] { 11, 13, 15, 17 }));
    }

}
