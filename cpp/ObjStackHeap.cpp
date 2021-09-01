#include <iostream>

using namespace std;

// 只能在堆上创建对象的类
class HeapObj {
private:
    ~HeapObj() {}
};


// 只能在栈上创建对象的类
class StackObj {
private:
    void *operator new(size_t val) { return malloc(val); }
};


int main() {
    //HeapObj obj1;    //报错
    HeapObj *obj1 = new HeapObj; // 成功

    //StackObj *obj2 = new StackObj;  // 报错
    StackObj obj2;
    return 0;
}
