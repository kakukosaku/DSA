import java.util.HashMap;
import java.util.Map;

public class RomanToInteger {
    public int romanToInt(String s) {
        Map<String, Integer> m = new HashMap<>();
        m.put("I", 1);
        m.put("II", 2);
        m.put("III", 3);
        m.put("IV", 4);
        m.put("V", 5);
        m.put("IX", 9);
        m.put("X", 10);
        m.put("XL", 40);
        m.put("L", 50);
        m.put("XC", 90);
        m.put("C", 100);
        m.put("CD", 400);
        m.put("D", 500);
        m.put("CM", 900);
        m.put("M", 1000);

        int rest = 0;
        for (int i = 0; i < s.length();) {
            for (int j = 3; j > 0; j--) {
                if (i+j > s.length()) {
                    continue;
                }
                if (m.get(s.substring(i, i + j)) != null) {
                    rest += m.get(s.substring(i, i + j));
                    i += j;
                    break;
                }
            }
        }
        return rest;
    }

    public static void main(String[] args) {
        System.out.println(new RomanToInteger().romanToInt("MCMXCIV"));
    }
}
