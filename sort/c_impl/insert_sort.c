//
// Created by kaku on 2019/10/7.
//

void insert_sort(int *arr, int array_size) {
    int i, j, temp;

    for (i = 1; i < array_size; i++) {
        // from small to lager.
        if (arr[i] < arr[i - 1]) {
            temp = arr[i];
            // find the right position, and move others for space.
            for (j = i - 1; arr[j] > temp && j >= 0; j--) {
                arr[j + 1] = arr[j];
            }
            // put in right position.
            arr[j + 1] = temp;
        }
    }
}