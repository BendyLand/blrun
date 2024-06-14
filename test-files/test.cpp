#include <iostream>

using namespace std;

int add(int a, int b);

int main()
{
    cout << "This is a test of my new run tool" << endl;
    int result = add(512, 22);
    cout << "The result of the add function was: " << result << endl;
}

int add(int a, int b)
{
    return a + b;
}
