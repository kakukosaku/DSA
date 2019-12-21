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
        int start_index = 0, maxNoRepeatSubStringLen = 0;
        int i;
        for (i = 0; i < s.length(); i++) {
            Character checking_char = s.charAt(i);
            Integer idx = recordMap.get(checking_char);
            // 解决hashMap中存在早于start_index的无用值, 对判断重复子串的影响
            if (idx == null || (idx != null && idx < start_index)) {
                recordMap.put(checking_char, i);
                maxNoRepeatSubStringLen = (i - start_index + 1) > maxNoRepeatSubStringLen ? i - start_index + 1 : maxNoRepeatSubStringLen;
            } else {
                start_index = recordMap.get(checking_char) + 1;
                recordMap.put(checking_char, i);
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
