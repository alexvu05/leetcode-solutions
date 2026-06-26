class Solution {
    public int compress(char[] chars) {
        int n = chars.length;
        int write = 0;  // position to write the compressed result
        int read  = 0;  // start of the current group of repeated chars
        while (read < n) {
            char currentChar = chars[read];
            int groupStart = read;
            // Advance read while the same character repeats
            while (read < n && chars[read] == currentChar) {
                read++;
            }
            int groupLength = read - groupStart;
            // Write the character itself
            chars[write++] = currentChar;
            // Write the count only if the group has more than 1 character
            if (groupLength > 1) {
                for (char digit : String.valueOf(groupLength).toCharArray()) {
                    chars[write++] = digit;
                }
            }
        }
        return write;
    }
}