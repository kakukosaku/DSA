/*
 * author: kaku
 * date: 2019/月/日
 *
 * GitHub:
 *   https://github.com/kakukosaku
 *
 * Description:
 */

import java.util.HashMap;
import java.util.Map;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        Map<Character, Integer> recordMap = new HashMap<>();
        int startIndex = 0, maxNoRepeatSubStringLen = 0;
        int i;
        for (i = 0; i < s.length(); i++) {
            Character checkingChar = s.charAt(i);
            Integer idx = recordMap.get(checkingChar);
            // 解决hashMap中存在早于start_index的无用值, 对判断重复子串的影响
            if (idx == null || (idx != null && idx < startIndex)) {
                recordMap.put(checkingChar, i);
                maxNoRepeatSubStringLen = (i - startIndex + 1) > maxNoRepeatSubStringLen ? i - startIndex + 1 : maxNoRepeatSubStringLen;
            } else {
                startIndex = recordMap.get(checkingChar) + 1;
                recordMap.put(checkingChar, i);
            }
        }

        return maxNoRepeatSubStringLen;
    }


    public static void main(String[] args) {
        Solution s = new Solution();
        int rest = s.lengthOfLongestSubstring("tmmzuxt");
        System.out.println(rest);
    }
}
