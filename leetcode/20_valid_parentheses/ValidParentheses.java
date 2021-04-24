import java.util.HashMap;
import java.util.Map;
import java.util.Stack;

public class ValidParentheses {
    public boolean isValid(String s) {
        Map<Character, Character> m = new HashMap<>();
        m.put('(', ')');
        m.put('[', ']');
        m.put('{', '}');

        Stack<Character> stack = new Stack<>();
        for (char i : s.toCharArray()) {
            if (m.get(i) != null) {
                stack.push(i);
                continue;
            }

            if (stack.empty()) {
                return false;
            }
            if (i != m.get(stack.pop())) {
                return false;
            }
        }

        return stack.empty();
    }

    public static void main(String[] args) {
        System.out.println(new ValidParentheses().isValid("{"));
    }
}
