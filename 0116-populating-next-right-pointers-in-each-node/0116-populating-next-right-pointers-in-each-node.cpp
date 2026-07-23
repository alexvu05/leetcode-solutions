/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/

class Solution {
public:
    Node* connect(Node* root) {
        if (root == nullptr) return nullptr;
        queue<Node*> bfsQueue;
        bfsQueue.push(root);
        while (!bfsQueue.empty()) {
            int layerSize = bfsQueue.size();
            for (int i = 0; i < layerSize; i++) {
                Node* node = bfsQueue.front();
                bfsQueue.pop();
                // Connect to the next node in the same level
                if (i < layerSize - 1) {
                    node->next = bfsQueue.front();
                }
                if (node->left)  bfsQueue.push(node->left);
                if (node->right) bfsQueue.push(node->right);
            }
        }
        return root;
    }
};