#include <iostream>
#include <vector>
#include <string.h>
#include <assert.h>
#include <algorithm>

using namespace std;

/**
 * 全排列：n个元素的所有排列组合，如123的全排列有：123、132、213、231、312、321 六种
 * 全排列通常有两种类型：
 * 		1. 不含重复元素的全排列
 * 		2. 含重复元素的全排列，此时需要去重
 */

/****************回溯法**************************/

// 1. 使用递归，每一层递归都将第一个元素与后续元素交换
// 默认不含有重复元素
void FullPermutation(char *str, int start, int end) {
	assert(str);

	if (start == end)
		cout << str << endl;

	for (int i = start; i < end; i++) {
		swap(str[start], str[i]);
		FullPermutation(str, start+1, end);
		swap(str[i], str[start]);
	}
}

// 考虑含有重复元素的情况
// 去重规则：从第一个数字起每个数分别与它后面非重复出现的数字交换
// 在 [start, end) 中检测是否存在 str[end] 相等的元素
bool isSwap(char *str, int start, int end) {
	for (int i = start; i < end; i++) {
		if (str[i] == str[end])
			return false;
	}

	return true;
}

void FullPermutationUnique(char *str, int start, int end) {
	assert(str);

	if (start == end)
		cout << str << endl;

	for (int i = start; i < end; i++) {
		if (isSwap(str, start, i)) {
			swap(str[start], str[i]);
			FullPermutation(str, start+1, end);
			swap(str[i], str[start]);
		}
	}
}

/***********************字典序法*********************/
/**
 * 字典序法的核心是找到下一个序列，下一个序列是指将全排列按照字典序排序，当前序列的下一个序列
 * 以序列 "158476531" 为例，找到它的下一个排列的方法：
 *  1. 从当前排列 "158476531" 的右端开始，找到第一个比其右边数字小的数字 "4"（找最后一个正序）
 *  2. 在 "4" 的右边数字中，找出所有比 "4" 大的数字中最小的数字 "5"
 *  3. 交换 "4" 和 "5" 的位置，此时生成的排列为 "158576431"
 *  4. 再将现在序列中 "5" 的位置之后的数字 "76431" 进行倒转，生成下一个字典序列 "158513467"
 **/

bool nextPermutation(char *str) {
	int len = strlen(str);
	//return std::next_permutation(str, str+len); // C++库
	int start = -1;
	for (int i = len-1; i > 0; i--) {
		if (str[i-1] < str[i]) {
			start = i - 1;
			break;
		}

	}
	//说明已经是最大序列了，即逆序序列
	if (start == -1)
		return false;

	int min = -1;
	for (int j = start; j < len; j++) {
		if (str[j] > str[start] &&  (min == -1 || str[j] < str[min]))
			min = j;
	}

	swap(str[start], str[min]);

	for (int i = start+1, j = len-1; i < j; i++, j--)
		swap(str[i], str[j]);

	return true;
}

void FullPremutationDicSort(char *str) {
	int len = strlen(str);
	sort(str, str+len);
	cout << str << endl;
	while (nextPermutation(str)) {
		cout << str << endl;
	}
}

int main() {
	char str[] = "123";

	FullPermutation(str, 0, 3);

	cout << endl;

	char str2[] = "122";

	FullPermutationUnique(str2, 0, 3);

	cout << endl;

	FullPremutationDicSort(str);
}