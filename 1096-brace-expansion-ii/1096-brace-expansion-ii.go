func braceExpansionII(expression string) []string {
    idx := 0

    // parse: returns sorted deduplicated set for current level
    // stops at end of string or at '}'
    var parse func() []string

    // parseGroup: handles concatenation of items at current level
    // items are separated by ',' (handled by caller)
    var parseConcats func() []string

    // parseItem: single letter or {...} block
    var parseItem func() []string

    parseItem = func() []string {
        if idx < len(expression) && expression[idx] == '{' {
            idx++ // consume '{'
            result := parse()
            idx++ // consume '}'
            return result
        }
        // single letter
        ch := string(expression[idx])
        idx++
        return []string{ch}
    }

    parseConcats = func() []string {
        result := []string{""}
        for idx < len(expression) && expression[idx] != ',' && expression[idx] != '}' {
            items := parseItem()
            // Cartesian product: result x items
            var next []string
            for _, a := range result {
                for _, b := range items {
                    next = append(next, a+b)
                }
            }
            result = next
        }
        return sortDedup(result)
    }

    parse = func() []string {
        var result []string
        // Union of concatenation groups separated by ','
        for {
            group := parseConcats()
            result = append(result, group...)
            if idx >= len(expression) || expression[idx] != ',' {
                break
            }
            idx++ // consume ','
        }
        return sortDedup(result)
    }

    return parse()
}

func sortDedup(s []string) []string {
    if len(s) == 0 { return s }
    sort.Strings(s)
    result := s[:1]
    for i := 1; i < len(s); i++ {
        if s[i] != s[i-1] {
            result = append(result, s[i])
        }
    }
    return result
}