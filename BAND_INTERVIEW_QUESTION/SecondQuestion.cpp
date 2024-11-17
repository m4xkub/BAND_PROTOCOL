#include <iostream>
#include <algorithm>

using namespace std;

int n, k;
vector<int> position;
int ans = 0;

int main()
{
    cout << "Number of chickens are : ";
    cin >> n;
    cout << "Length of roof is : ";
    cin >> k;

    for (int i = 0; i < n; i++)
    {
        int temp;
        cin >> temp;
        position.push_back(temp);
    }

    int start = 0, stop = 0;

    while (stop < n)
    {
        int distance = position[stop] - position[start] + 1;
        if (distance <= k)
        {
            ans = max(ans, stop - start + 1);
            stop++;
        }
        else
        {
            start++;
        }
    }

    cout << ans;
}