package tree

// #include <iostream>

// using namespace std;

// struct TreeLinkNode {
//     int val;
//     TreeLinkNode *left, *right, *next;
//     TreeLinkNode(int x) : val(x), left(NULL), right(NULL), next(NULL) {}
// };

// class Solution {
// public:
//     void connect(TreeLinkNode *root) {
//         if (NULL == root) {
//             return;
//         }
//         if (root -> left != NULL){
//             root -> left -> next = root -> right;
//             auto next = root -> next;
//             root -> right -> next = next ? next -> left : NULL;
//             connect(root->left);
//             connect(root -> right);
//         }
//     }
// };
