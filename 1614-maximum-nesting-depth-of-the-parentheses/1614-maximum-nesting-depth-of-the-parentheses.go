func maxDepth(s string) int {
    depth, best := 0, 0
    for _, c := range s {
        if c == '(' {
            depth++
            if depth > best {
                best = depth
            }
        } else if c == ')' {
            depth--
        }
    }
    return best
}