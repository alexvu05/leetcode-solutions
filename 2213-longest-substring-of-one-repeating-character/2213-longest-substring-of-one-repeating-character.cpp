class Solution {
    struct Node {
        int lmx;  // longest prefix run
        int rmx;  // longest suffix run
        int mx;   // longest run in interval
        int l, r; // interval bounds
    };

    string s;
    vector<Node> tree;

    void build(int node, int l, int r) {
        tree[node].l = l;
        tree[node].r = r;
        if (l == r) {
            tree[node].lmx = tree[node].rmx = tree[node].mx = 1;
            return;
        }
        int mid = (l + r) / 2;
        build(2*node, l, mid);
        build(2*node+1, mid+1, r);
        pushup(node);
    }

    void pushup(int node) {
        auto& L = tree[2*node];
        auto& R = tree[2*node+1];
        auto& cur = tree[node];
        int lLen = L.r - L.l + 1;
        int rLen = R.r - R.l + 1;
        // Prefix: left prefix + possibly extend into right
        cur.lmx = L.lmx;
        if (L.lmx == lLen && s[L.r] == s[R.l])
            cur.lmx = lLen + R.lmx;
        // Suffix: right suffix + possibly extend into left
        cur.rmx = R.rmx;
        if (R.rmx == rLen && s[L.r] == s[R.l])
            cur.rmx = rLen + L.rmx;
        // Max: either side, or merge across boundary
        cur.mx = max(L.mx, R.mx);
        if (s[L.r] == s[R.l])
            cur.mx = max(cur.mx, L.rmx + R.lmx);
    }

    void update(int node, int pos) {
        if (tree[node].l == tree[node].r) {
            tree[node].lmx = tree[node].rmx = tree[node].mx = 1;
            return;
        }
        int mid = (tree[node].l + tree[node].r) / 2;
        if (pos <= mid) update(2*node, pos);
        else update(2*node+1, pos);
        pushup(node);
    }

public:
    vector<int> longestRepeating(string s, string queryCharacters, vector<int>& queryIndices) {
        this->s = s;
        int n = s.size();
        tree.resize(4 * n);
        build(1, 0, n - 1);
        int k = queryCharacters.size();
        vector<int> result(k);
        for (int i = 0; i < k; i++) {
            // Apply update
            this->s[queryIndices[i]] = queryCharacters[i];
            update(1, queryIndices[i]);
            result[i] = tree[1].mx;
        }
        return result;
    }
};