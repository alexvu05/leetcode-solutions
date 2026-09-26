func evaluate(s string, knowledge [][]string) string {
    // Build lookup map
    lookup := make(map[string]string, len(knowledge))
    for _, kv := range knowledge {
        lookup[kv[0]] = kv[1]
    }

    var sb strings.Builder
    i := 0
    for i < len(s) {
        if s[i] == '(' {
            // Find closing ')'
            j := i + 1
            for j < len(s) && s[j] != ')' {
                j++
            }
            key := s[i+1 : j]
            if val, ok := lookup[key]; ok {
                sb.WriteString(val)
            } else {
                sb.WriteByte('?')
            }
            i = j + 1
        } else {
            sb.WriteByte(s[i])
            i++
        }
    }

    return sb.String()
}