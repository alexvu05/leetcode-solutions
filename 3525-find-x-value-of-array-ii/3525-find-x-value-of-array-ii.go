func resultArray(nums []int, k int, queries [][]int) []int {
    n := len(nums)

    type Node struct {
        cnt  []int // cnt[v] = số prefix của đoạn có tích%k == v
        prod int   // tích toàn đoạn % k
    }

    newLeaf := func(val int) Node {
        node := Node{cnt: make([]int, k), prod: val % k}
        node.cnt[val%k] = 1
        return node
    }

    newEmpty := func() Node {
        return Node{cnt: make([]int, k), prod: 1}
    }

    // Merge: prefix of [L,R] = prefix of L OR (full L * prefix of R)
    // cnt of merged:
    //   - prefixes entirely in L: L.cnt[v]
    //   - prefixes spanning L into R: L.prod * prefix of R ending at some point
    merge := func(L, R Node) Node {
        res := newEmpty()
        res.prod = (L.prod * R.prod) % k

        // Prefixes entirely in L
        for v := 0; v < k; v++ {
            res.cnt[v] = L.cnt[v]
        }

        // Prefixes spanning: full L (prod=L.prod) then prefix of R with prod u
        // combined prod = L.prod * u % k
        for u := 0; u < k; u++ {
            if R.cnt[u] == 0 {
                continue
            }
            combined := (L.prod * u) % k
            res.cnt[combined] += R.cnt[u]
        }

        return res
    }

    tree := make([]Node, 4*n)
    for i := range tree {
        tree[i] = newEmpty()
    }

    var build func(node, l, r int)
    build = func(node, l, r int) {
        if l == r {
            tree[node] = newLeaf(nums[l])
            return
        }
        mid := (l + r) / 2
        build(2*node, l, mid)
        build(2*node+1, mid+1, r)
        tree[node] = merge(tree[2*node], tree[2*node+1])
    }
    build(1, 0, n-1)

    var update func(node, l, r, pos int)
    update = func(node, l, r, pos int) {
        if l == r {
            tree[node] = newLeaf(nums[pos])
            return
        }
        mid := (l + r) / 2
        if pos <= mid {
            update(2*node, l, mid, pos)
        } else {
            update(2*node+1, mid+1, r, pos)
        }
        tree[node] = merge(tree[2*node], tree[2*node+1])
    }

    var query func(node, l, r, ql, qr int) Node
    query = func(node, l, r, ql, qr int) Node {
        if ql <= l && r <= qr {
            return tree[node]
        }
        mid := (l + r) / 2
        if qr <= mid {
            return query(2*node, l, mid, ql, qr)
        }
        if ql > mid {
            return query(2*node+1, mid+1, r, ql, qr)
        }
        L := query(2*node, l, mid, ql, qr)
        R := query(2*node+1, mid+1, r, ql, qr)
        return merge(L, R)
    }

    result := make([]int, len(queries))
    for i, q := range queries {
        idx, val, start, x := q[0], q[1], q[2], q[3]

        nums[idx] = val
        update(1, 0, n-1, idx)

        node := query(1, 0, n-1, start, n-1)
        result[i] = node.cnt[x]
    }

    return result
}