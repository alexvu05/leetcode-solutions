class Solution:
    def findSubtreeSizes(self, parent: List[int], s: str) -> List[int]:
        # Time: O(n) | Space: O(n)
        n = len(parent)
        children = [[] for _ in range(n)]
        for v in range(1, n):
            children[parent[v]].append(v)

        new_parent = parent[:]            # parent in the final tree
        last = [-1] * 26                  # last[c] = closest ancestor on the current path with char c
        saved = [-1] * n                  # value of last[c] to restore when leaving v
        order = []                        # preorder of the ORIGINAL tree

        stack = [0]
        while stack:
            v = stack.pop()
            if v >= 0:                    # entering v
                order.append(v)
                c = ord(s[v]) - 97
                if last[c] != -1:
                    new_parent[v] = last[c]    # closest same-char ancestor (original tree)
                saved[v] = last[c]
                last[c] = v
                stack.append(~v)               # marker: leave v after its subtree
                stack.extend(children[v])
            else:                         # leaving v
                v = ~v
                last[ord(s[v]) - 97] = saved[v]

        # new_parent[v] is an original ancestor of v → appears earlier in preorder.
        # Reverse preorder therefore finishes every node before its new parent.
        size = [1] * n
        for v in reversed(order):
            if v != 0:
                size[new_parent[v]] += size[v]

        return size