class LongestPalindromicSubString {

    public String longestPalindrome(String s) {
        int[][] dp = new int[s.length()][s.length()];
        String result = "";
        if (s.equals("")) {
            return result;
        }

        for (int l = 0; l < s.length(); l++) {
            for (int i = 0; i + l < s.length(); i++) {
                int j = i + l;
                if (l == 0) {
                    dp[i][j] = 1;
                } else if (l == 1) {
                    if (s.charAt(i) == s.charAt(j)) {
                        dp[i][j] = 1;
                    }
                } else {
                    if (s.charAt(i) == s.charAt(j)) {
                        dp[i][j] = dp[i + 1][j - 1];
                    }
                }

                if (dp[i][j] > 0 && l + 1 > result.length()) {
                    result = s.substring(i, j + 1);
                }
            }
        }

        return result;
    }

    public static void main(String[] args) {
        System.out.println(new LongestPalindromicSubString().longestPalindrome("aaaa"));
    }

}