/*
 * author: kaku
 * date: 2019/月/日
 *
 * GitHub:
 *   https://github.com/kakukosaku
 *
 * Description:
 */
import java.util.Arrays;

public class HeapSort {

    public static void heapnify(int[] arr, int check, int size) {
        int elem = arr[check];

        while (check * 2 + 1 < size) {
            int lchild = check * 2 + 1;
            int rchild = check * 2 + 2;
            int bigger = lchild;

            if (rchild < size && arr[rchild] > arr[bigger]) {
                bigger = rchild;
            }

            if (elem > arr[bigger]) {
                break;
            } else {
                arr[check] = arr[bigger];
                check = bigger;
            }
        }

        arr[check] = elem;
    }

    public static void buildHeap(int[] arr) {
        int size = arr.length;
        for (int check = size / 2 - 1; check >= 0; check--) {
            heapnify(arr, check, size);
        }
    }

    public static void heapSort(int[] arr) {
        int size = arr.length;
        buildHeap(arr);
        for (int sortedNum = 1; sortedNum < size; sortedNum++) {
            swap(arr, 0, size - sortedNum);
            heapnify(arr, 0, size - sortedNum);
        }
    }

    public static void swap(int[] arr, int i, int j) {
        int tmp = arr[i];
        arr[i] = arr[j];
        arr[j] = tmp;
    }

    public static void main(String[] args) {
        int[] arr = new int[]{5, 3, 2, 1, 6, 7, 23, 6, 7};
        System.out.println(Arrays.toString(arr));
        heapSort(arr);
        System.out.println(Arrays.toString(arr));
    }

}
