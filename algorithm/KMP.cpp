#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int *getNext(const char *pattern, int *next) {
    next[0] = -1;

    int j = 0;
    int k = -1;

    while (j < strlen(pattern) - 1) {
        if (k == -1 || pattern[j] == pattern[k]) {
            if (pattern[++j] == pattern[++k])
                next[j] = next[k];
            else
                next[j] = k;
        } else {
            k = next[k];
        }
    }

    return next;
}

int KMP(const char *str, const char *pattern, int *next) {
    getNext(pattern, next);

    int i = 0;
    int j = 0;

    int str_len = strlen(str);
    int patter_len = strlen(pattern);

    while (i < str_len && j < patter_len) {
        if (j == -1 || str[i] == pattern[j]) {
            i++;
            j++;
        } else {
            j = next[j];
        }
    }

    if (j == patter_len) {
        return i - j;
    } else {
        return -1;
    }
}


int main() {

	int i;
	int next[20] = {0};

	const char *text = "ababxbababababcdababcabddcadfdsss";
	const char *pattern = "abcabd";

	int idx = KMP(text, pattern, next);

	printf("match pattern : %d\n", idx);

	for (i = 0;i < strlen(pattern);i ++) {
		printf("%4d", next[i]);
	}
	printf("\n");

	return 0;
}