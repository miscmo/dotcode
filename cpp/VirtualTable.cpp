#include <iostream>

using namespace std;

class A {
public:
    virtual ~A() {}

    virtual void v_name() { cout << "vname(): I am A" << endl; }

    void name() { cout << "name(): I am A" << endl; }
};

class B : public A {
public:
    virtual ~B() {}

    virtual void v_name() { cout << "vname(): I am B" << endl; }

    void name() { cout << "name(): I am B" << endl; }
};

int main() {
    A a;
    a.name();

    cout << "class A size: " << sizeof(A) << endl;

    B b;
    b.name();

    cout << "class B size: " << sizeof(B) << endl;

    A *ap = &b;
    ap->v_name(); // vname(): I am B
    ap->name();   // name(): I am A

    A ao = b;
    ao.v_name();  // vname(): I am A
    ao.name();  // name(): I am A

}
