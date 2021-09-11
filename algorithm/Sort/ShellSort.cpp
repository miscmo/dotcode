#include <iostream>
#include <vector>

using namespace std;

void printVec(const vector<int> &arr) {
	for (auto a : arr)
		cout << a << " ";
	cout << endl;
}

void ShellSort(vector<int> &arr) {
	int len = arr.size();

	for (int gap = len / 2; gap > 0; gap /= 2) {
		for (int i = gap; i < len; i++) {
			for (int j = ; j >= gap ; j -= gap)
		}
	}
}

int main() {
	int arr[] = {8, 7, 3, 9, 1, 9, 5, 4, 2};
	//int arr[] = {9, 5};

	int len = sizeof(arr) / sizeof(arr[0]);

	vector<int> vec(arr, arr+len);

	ShellSort(vec);

	for (auto i : vec) {
		cout << i << " ";
	}
	cout << endl;
}