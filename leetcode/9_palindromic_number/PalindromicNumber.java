public class PalindromicNumber {

    public boolean isPalindrome(int x) {
        if (x < 0 || (x % 10 == 0 && x != 0)) {
            return false;
        }

        int partOfX = 0;
        while (x > partOfX) {
            int remind = x % 10;
            x = x / 10;
            partOfX = partOfX * 10 + remind;

        }

        return partOfX == x || (x == partOfX / 10);
    }

    public static void main(String[] args) {
        System.out.println(new PalindromicNumber().isPalindrome(121));
    }

}
