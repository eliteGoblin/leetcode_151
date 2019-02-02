
#include <stdio.h>

 // Definition for singly-linked list with a random pointer.
struct RandomListNode {
    int label;
    struct RandomListNode *next;
    struct RandomListNode *random;
};

RandomListNode *makeNode(int label){
    auto p = new(RandomListNode);
    p->label = label;
    return p;
}

struct RandomListNode *copyRandomList(struct RandomListNode *head) {
    for (auto cur = head;cur != NULL;) {
        auto cloneNode = makeNode(cur->label);
        cloneNode->next = cur->next;
        cur->next = cloneNode;
        cur = cloneNode->next;
    }
    // copy random pointer
    for (auto cur = head; cur != NULL;) {
        auto cloneNode = cur->next;
        if (cur->random == NULL){
            cloneNode->random = NULL;
        }else {
            cloneNode->random = cur->random->next;
        }
        cur = cloneNode->next;
    }
    // split list
    auto pseudoHead = makeNode(-1);
    auto curClonePre = pseudoHead;
    for (auto cur = head; cur != NULL;) {
        auto curClone = cur->next;
        auto nextCur = curClone->next;
        curClonePre->next = curClone;
        curClonePre = curClone;
        cur->next = nextCur;
        cur = nextCur;
    }
    curClonePre->next = NULL;
    return pseudoHead->next;
}

int main() {
    return 0;
}