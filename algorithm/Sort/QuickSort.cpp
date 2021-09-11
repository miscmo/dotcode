#include <iostream>
#include <assert.h>
#include <vector>

using namespace std;

int partition(vector<int> &arr, int start, int end) {
	assert(arr.size() > 0);

	int i = start+1;
	int j = end;

	// 1. while (true)
	while (true) {
		while (i <= end && arr[i] < arr[start]) 
			i++;
		while (j >= start && arr[j] > arr[start])
			j--;

		// 2. 结束条件
		if (i >= j)
			break;

		swap(arr[i], arr[j]);
	}

	// 3. 最后跟j交换
	swap(arr[j], arr[start]);

	return j;
}

void QuickSort(vector<int> &arr, int start, int end) {
	assert(arr.size() > 0);

	// 4. 结束条件
	if (start >= end)
		return ;

	// 5. [start, end]
	int m = partition(arr, start, end);
	QuickSort(arr, start, m-1);
	QuickSort(arr, m+1, end);
}

int main() {
	int arr[] = {4, 5, 3, 2, 1, 9, 7, 8, 10};
	//int arr[] = {9, 5};

	int len = sizeof(arr) / sizeof(arr[0]);

	vector<int> vec(arr, arr+len);

	QuickSort(vec, 0, vec.size()-1);
}