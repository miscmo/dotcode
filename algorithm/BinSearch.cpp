#include <iostream>
#include <vector>

using namespace std;

int bin_search(const vector<int> &arr, int target) {
	int start = 0;
	int mid = 0;
	int end = arr.size() - 1;	// 注意：-1

	// 这里 <= 防止只有一个元素的时候找不到
	while (start <= end) {
		mid = (start + end) / 2;

		if (arr[mid] == target)
			return mid;
		else if (arr[mid] > target)
			end = mid - 1;
		else 
			start = mid + 1;
	}

	// 找不到返回 -1
	return -1;
}

int main() {
	int arr[10] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};

	vector<int> vec(arr, arr+10);

	// 查找 8
	int index = bin_search(vec, arr[8]);
	if (-1 == index)
		cout << "not find target " << arr[8] << endl;
	else
		cout << "find: " << vec[index] << " index: " << index << endl;

	// 查找 2
	index = bin_search(vec, arr[2]);
	if (-1 == index)
		cout << "not find target " << arr[2] << endl;
	else
		cout << "find: " << vec[index] << " index: " << index << endl;

	// 查找不存在的值
	index = bin_search(vec, -1);
	if (-1 == index)
		cout << "not find target " << -1 << endl;
	else
		cout << "find: " << vec[index] << " index: " << index << endl;

	return 0;
}