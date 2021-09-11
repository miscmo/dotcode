#include <iostream>
#include <vector>
#include <string.h>

using namespace std;

/**
 * 求 abc 的组合：a, b, c, ab, ac, bc, abc
 */

// 二进制法
void printStr(int bin, char *str, int len) {
	for (int i = len-1; i >= 0; i--) {
		if (bin & (1 << i))
			cout << str[i];
	}
	cout << endl;
}

void Combinatorial1(char *str) {
	int s = 1;
	int len = strlen(str);
	int e = 1 << len;

	while (s < e) {
		printStr(s, str, len);
		s++;
	}
}

// 递归
// 1. 每次从str中选择第i个元素，然后再从i后面的元素中选择下一个元素

void Combinatorial2(vector<char> str, vector<char> ret) {
	for (auto c : ret) {
		cout << c;
	}
	cout << endl;

	for (int i = 0; i < str.size(); i++) {
		ret.push_back(str[i]);
		Combinatorial2(vector<char>(str.begin()+i+1, str.end()), ret);
		ret.pop_back();
	}
}


int main() {
	char str[] = "abc";

	//Combinatorial1(str);
	vector<char> vec_str(str, str+3);
	vector<char> ret;
	Combinatorial2(vec_str, ret);

	return 0;
}