#include <iostream>
#include <string.h>

using namespace std;

void ReverseStr(char *str) {
	if (str == nullptr)
		return ;

	int len = strlen(str);

	for (int i = 0, j = len-1; i < j; i++, j--) {
		swap(str[i], str[j]);
	}
}

int main() {
	//char str[] = "helloworld";
	//char str[] = "h";
	char str[] = "he";

	ReverseStr(str);

	cout << str << endl;
	return 0;
}