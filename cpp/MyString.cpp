#include <iostream>
#include <string.h>

using namespace std;

class MyString {
public:
    MyString(const char *str) {
        if (str == nullptr) {
            m_pStr = new char[1];
            m_pStr[0] = '\0';
        } else {
            int len = strlen(str);
            m_pStr = new char[len+1];
            strcpy(m_pStr, str);
        }
    }

    ~MyString() {
        delete [] m_pStr;
        m_pStr = nullptr;
    }

    MyString& operator=(const MyString &str) {
        if (this != &str) {
            MyString strTemp(str);

            char *pTemp = strTemp.m_pStr;
            strTemp = this->m_pStr;
            this->m_pStr = pTemp;
        }

        return *this;
    }

    // 注意，拷贝构造函数内部不可像operator=那样写，否则会导致栈溢出
    MyString(const MyString &str) {
        int len = strlen(str.m_pStr);
        m_pStr = new char[len+1];
        strcpy(m_pStr, str.m_pStr);
    }

    void output() {
        cout << m_pStr << endl;
    }

private:
    char *m_pStr;
};

int main() {
    char cstr[10] = "123456789";
    MyString str(cstr);
    str.output();
    MyString str2(str);
    str2.output();
    return 0;
}
