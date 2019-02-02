package tree

// #include <iostream>

// using namespace std;

// struct TreeLinkNode {
//  int val;
//  TreeLinkNode *left, *right, *next;
//  TreeLinkNode(int x) : val(x), left(NULL), right(NULL), next(NULL) {}
// };

// class Solution {
// 	public:
// 		void connect(TreeLinkNode *root) {
// 			if (NULL == root) {
// 				return;
// 			}
// 			// 递归假设：当前层next已经设置好
// 			// 取得next node的最左孩子
// 			TreeLinkNode *next = root -> next;
// 			while (NULL != next) {
// 				if (NULL != next -> left) {
// 					next = next -> left;
// 					break;
// 				}
// 				if (NULL != next -> right) {
// 					next = next -> right;
// 					break;
// 				}
// 				next = next -> next;
// 			}
// 			if (root -> right) {
// 				root -> right -> next = next;
// 			}
// 			if (root -> left){
// 					root -> left -> next = (root -> right) ? root -> right : next;
// 			}
// 			connect(root -> right);
// 			connect(root -> left);
// 		}
// };
