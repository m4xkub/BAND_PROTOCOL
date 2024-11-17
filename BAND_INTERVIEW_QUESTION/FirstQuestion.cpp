#include <iostream>
#include <algorithm>

using namespace std;

int main()
{
    int bb = 0;
    string shooting;
    cin >> shooting;
    int length = shooting.length();
    if (!length)
    {
        cout << "Good boy\n";
        return 0;
    }
    if (shooting[0] == 'R')
    {
        cout << "Bad boy\n";
        return 0;
    }

    for (int i = 0; i < length; i++)
    {
        if (shooting[i] == 'S')
        {
            bb++;
        }
        else if (bb > 0)
        {
            bb--;
        }
    }

    if (bb)
        cout << "Bad boy\n";
    else
        cout << "Good boy\n";

    return 0;
}