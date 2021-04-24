import java.util.HashMap;
import java.util.Map;

public class IntegerToRoman {
    public String intToRoman(int num) {
        Map<Integer, String> m = new HashMap<>();
        m.put(1, "I");
        m.put(2, "II");
        m.put(3, "III");
        m.put(4, "IV");
        m.put(5, "V");
        m.put(9, "IX");
        m.put(10, "X");
        m.put(40, "XL");
        m.put(50, "L");
        m.put(90, "XC");
        m.put(100, "C");
        m.put(400, "CD");
        m.put(500, "D");
        m.put(900, "CM");
        m.put(1000, "M");

        int[] arr = new int[] { 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 3, 2, 1 };
        StringBuilder sb = new StringBuilder();

        for (int ele : arr) {
            while (num >= ele) {
                sb.append(m.get(ele));
                num -= ele;
            }
        }

        return sb.toString();
    }

    public static void main(String[] args) {
        System.out.println(new IntegerToRoman().intToRoman(58));
    }
}
