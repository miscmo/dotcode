#include <iostream>

using namespace std;

void test_endian1() {
    int i = 1;
    char *a = (char *)&i;

    if ((*a) == 1) {
        cout << "小端" << endl;
    } else {
        cout << "大端" << endl;
    }
}

void test_endian2() {
    union A {
        char c;
        int i;
    };

    A a;
    a.i = 1;

    if (a.c == 1) {
        cout << "小端" << endl;
    } else {
        cout << "大端" << endl;
    }
}

int main() {
    test_endian1();
    test_endian2();
}
