/*
 * author: kaku
 * date: 2020/08/09
 *
 * GitHub:
 *   https://github.com/kakukosaku
 *
 * Description:
 */
import java.util.Arrays;

public class QuickSort {

    public static int partition(int[] arr, int low, int high) {
        int elem = arr[low];

        while (low < high) {
            for (; low < high && arr[high] <= elem; high--)
                ;
            arr[low] = arr[high];
            for (; low < high && arr[low] >= elem; low++)
                ;
            arr[high] = arr[low];
        }
        arr[low] = elem;
        return low;
    }

    public static void quickSort(int[] arr, int low, int high) {
        if (low < high) {
            int pivot = partition(arr, low, high);
            quickSort(arr, low, pivot-1);
            quickSort(arr, pivot+1, high);
        }
    }

    public static void main(String[] args) {
        int[] arr = new int[]{5,3,2,1,6,7,23,6,7};
        System.out.println(Arrays.toString(arr));
        quickSort(arr, 0, arr.length-1);
        System.out.println(Arrays.toString(arr));
    }

}
