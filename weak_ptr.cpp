#include <memory>
#include <iostream>

using namespace std;

#define WEAK_PTR_TEST

#ifdef SHARED_PTR_TEST

//测试 shared_ptr 的循环引用问题
class B;
class A {
public:
    A() { cout << "A()" << endl;; }
    ~A() { cout << "~A()" << endl; }

    void name() { cout << "I AM A." << endl; }

    shared_ptr<B> m_bPtr;
};

class B {
public:
    B() { cout << "B()" << endl;; }
    ~B() { cout << "~B()" << endl; }

    void name() { cout << "I AM B." << endl; }

    shared_ptr<A> m_aPtr;
};

int main() {
    A *obj_a = new A();
    B *obj_b = new B();

	shared_ptr<A> a_ptr(obj_a);
	shared_ptr<B> b_ptr(obj_b);

    a_ptr->m_bPtr = b_ptr;
    b_ptr->m_aPtr = a_ptr;
}

/** 结果：
  * A()
  * B()
  * 没有执行 A 和 B 的析构函数
  */
#endif

#ifdef WEAK_PTR_TEST

//测试使用 weak_ptr 解决 shared_ptr 的循环引用问题
class B;
class A {
public:
    A() { cout << "A()" << endl;; }
    ~A() { cout << "~A()" << endl; }

    void name() { cout << "A." << endl; }

    weak_ptr<B> m_bPtr;
};

class B {
public:
    B() { cout << "B()" << endl;; }
    ~B() { cout << "~B()" << endl; }

    void name() { cout << "B." << endl; }

    weak_ptr<A> m_aPtr;
};

int main() {
    A *obj_a = new A();
    B *obj_b = new B();

	shared_ptr<A> a_ptr(obj_a);
	shared_ptr<B> b_ptr(obj_b);

    a_ptr->m_bPtr = b_ptr;
    b_ptr->m_aPtr = a_ptr;

    if (a_ptr->m_bPtr.lock()) {
        a_ptr->m_bPtr.lock()->name();
    }
}

/** 结果：
  * A()
  * B()
  * 没有执行 A 和 B 的析构函数
  */

#endif