#include <iostream> 
#include <stack>

using namespace std;

struct Node {
	int val;
	Node *left;
	Node *right;
};

void PreTravelLoop(Node *head) {
	stack< pair<Node *, bool> > s;	
	s.push(make_pair(head, false));

	bool visited;

	while (!s.empty()) {
		head = s.top().first;
		visited = s.top().second;
		s.pop();

		if (head == nullptr)
			continue;

		if (visited) {
			cout << head->val << endl;
		} else {
			s.push(make_pair(head->right, false));
			s.push(make_pair(head->left, false));
			s.push(make_pair(head, true));
		}
	}
}

void InTravelLoop(Node *head) {
	stack< pair<Node *, bool> > s;
	s.push(make_pair(head, false));
	bool visited;

	while (!s.empty()) {
		head = s.top().first;
		visited = s.top().second;
		s.pop();

		if (head == nullptr)
			continue;
		
		if (visited) {
			cout << head->val << endl;
		} else {
			s.push(make_pair(head->right, false));
			s.push(make_pair(head, true));
			s.push(make_pair(head->left, false));
		}
	}
}

void PostTravelLoop(Node *head) {
	stack< pair<Node *, bool> > s;
	s.push(make_pair(head, false));
	bool visited;

	while (!s.empty()) {
		head = s.top().first;
		visited = s.top().second;
		s.pop();

		if (head == nullptr)
			continue;
		
		if (visited) {
			cout << head->val << endl;
		} else {
			s.push(make_pair(head, true));
			s.push(make_pair(head->right, false));
			s.push(make_pair(head->left, false));
		}
	}
}

int main() {
	return 0;
}