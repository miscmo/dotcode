#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>

char *MyStrcpy(char *dst, const char *src) {
    // 1. 入参校验，src参数为const
    if (dst == NULL || src == NULL)
        return dst;

    if (dst == src)
        return dst;

    char *ret = dst;

    while (*src != '\0') {
        *(dst++) = *(src++);
    }
    *dst = '\0';  //2. 末尾补零

    //3. 如果要考虑内存重叠，实现参考MyMemmove

    return ret;
}

// MyStrcat：将src字符串追加到dst结尾
char *MyStrcat(char *dst, const char *src) {
    if (dst == NULL || src == NULL)
        return dst;

    char *ret = dst;

    //1. 先找到dst的末尾
    while (*dst != '\0') {
        dst++;
    }

    while (*src != '\0') {
        *(dst++) = *(src++);
    }
    *dst = '\0'; //2. 记得末尾补零

    return ret;
}

int MyStrcmp(const char *str1, const char *str2) {
    //1. 入参校验
    assert(str1 != NULL && str2 != NULL);

    //2. 比较
    while (*str1 != '\0' && *str2 != '\0') {
        if (*str1 > *str2)
            return 1;
        else if (*str1 < *str2)
            return -1;

        str1++;
        str2++;
    }

    //3. 跳出循环后比较
    if (*str1 > *str2)
        return 1;
    else if (*str1 < *str2)
        return -1;

    return 0;
}

void *MyMemcpy(void *dst, void *src, size_t len) {
    //1. 判断入参
    if (dst == NULL || src == NULL)
        return NULL;

    if (dst == src) 
        return dst;

    //2. 提前保存返回值
    void *ret = dst;

    char *c_dst = (char *)dst;
    char *c_src = (char *)src;

    //Memcpy不需要考虑内存重叠
    while (len--) {
        *c_dst++ = *c_src++;
    }

    //3. 返回dest起始位置
    return ret;
}

void *MyMemmove(void *dst, void *src, size_t len) {
    if (dst == NULL || src == NULL)
        return NULL;

    if (dst == src) 
        return dst;

    void *ret = dst;

    char *c_dst = (char *)dst;
    char *c_src = (char *)src;

    //1. Memmove要解决内存重叠问题
    if (c_src > c_dst) {
        //源内存地址大于目标内存地址，说明源地址在目标地址后面，要从前往后复制
        while (len--) {
            *(c_dst++) = *(c_src++);
        }
    } else {
        //源内存地址小于目标内存地址，说明源地址在目标地址前面，要从后往前复制
        while (len--) {
            *(c_dst+len) = *(c_src+len);
        }
    }

    return ret;
}

int main() {
    char test[6] = {'1', '2', '3', '4', '5', '6'};

    //MyMemcpy(test, test+3, 3);

    printf("ret = %s\n", test);

    //MyMemmove(test, test+3, 3); //456456
    //MyMemmove(test, test+1, 3); //234456
    //MyMemmove(test, test, 3);  //123456
    MyMemmove(test+2, test, 3); //121236

    const char *hello = "helloworld";

    char *hello_cpy = (char *)malloc(2 * strlen(hello)+1);

    printf("MyMemcpy ret = %s\n", MyStrcpy(hello_cpy, hello));

    //printf("MyStrcat ret = %s\n", MyStrcat(hello_cpy, hello));

    printf("MyStrcmp ret = %d\n", MyStrcmp(hello_cpy, hello));

    printf("ret = %s\n", test);
}