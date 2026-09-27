func reverseParentheses(s string) string {
    stack := [][]byte{{}}

    for i := 0; i < len(s); i++ {
        c := s[i]
        if c == '(' {
            // Push new layer
            stack = append(stack, []byte{})
        } else if c == ')' {
            // Pop top layer, reverse it, append to layer below
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            // Reverse top
            for l, r := 0, len(top)-1; l < r; l, r = l+1, r-1 {
                top[l], top[r] = top[r], top[l]
            }
            stack[len(stack)-1] = append(stack[len(stack)-1], top...)
        } else {
            // Append char to current layer
            stack[len(stack)-1] = append(stack[len(stack)-1], c)
        }
    }

    return string(stack[0])
}