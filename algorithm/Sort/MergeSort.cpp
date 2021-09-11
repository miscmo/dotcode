#include <iostream>
#include <vector>

using namespace std;

/**
 * 归并排序是将两个已经排序的数据合并为一个有序数组
 * 归并排序分为 自顶向下和自底向上，两者都是通过调用 merge 函数实现的
 */

// 将两个有序数组，合并为一个
// 将 [l, m] 和 (m, h] 合并为一个有序数组
void merge(vector<int> &arr, int l, int m, int h) {
	vector<int> temp(arr.begin(), arr.end());

	// 注意边界
	int s1 = l;
	int s2 = m+1;
	for (int i = l; i <= h; i++) {
		// 先判断异常情况
		if (s1 > m) {
			arr[i] = temp[s2];
			s2++;
		} else if (s2 > h) {
			arr[i] = temp[s1];
			s1++;
		} else {
			if (temp[s1] < temp[s2]) {
				arr[i] = temp[s1];
				s1++;
			} else {
				arr[i] = temp[s2];
				s2++;
			}
		}
	}
}

// 1. 自顶向下
void mergeSort1Core(vector<int> &arr, int l, int h) {
	if (l >= h)
		return ;

	// 注意
	int m = (l + h) / 2;

	mergeSort1Core(arr, l, m);
	mergeSort1Core(arr, m+1, h);
	merge(arr, l, m, h);
}
void MergeSort1(vector<int> &arr) {
	if (arr.size() <= 1)
		return ;

	mergeSort1Core(arr, 0, arr.size()-1);
}

// TODO
// 2. 自底向上

void MergeSort2(vector<int> &arr) {

}

int main() {
	int arr[] = {4, 5, 3, 2, 1, 9, 7, 8, 10};
	//int arr[] = {9, 5};

	int len = sizeof(arr) / sizeof(arr[0]);

	vector<int> vec(arr, arr+len);

	MergeSort1(vec);

	for (auto i : vec) {
		cout << i << " ";
	}
	cout << endl;
}