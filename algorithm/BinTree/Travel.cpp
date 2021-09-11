#include <iostream>

using namespace std;

struct Node {
	int val;
	Node *left;
	Node *right;
};

void PreTravel(Node *head) {
	if (head == nullptr)
		return ;

	cout << head->val << endl;
	PreTravel(head->left);
	PreTravel(head->right);
}

void InTravel(Node *head) {
	if (head == nullptr)
		return ;

	InTravel(head->left);
	cout << head->val << endl;
	InTravel(head->right);
}

void PostTravel(Node *head) {
	if (head == nullptr)
		return ;

	PostTravel(head->left);
	PostTravel(head->right);
	cout << head->val << endl;
}

int main() {
	return 0;
}