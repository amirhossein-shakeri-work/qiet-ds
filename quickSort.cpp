#include <iostream>

using namespace std;

void quickSort(int[], int, int);
int split(int[], int, int);
void print(int[]);

int main () {
    int a[] = {10, 9, 1, 21, 11, 6, 5};

    quickSort(a, 0, sizeof(a) / sizeof(*a) - 1);

    print(a);

    return 0;
}

void quickSort (int a[], int l, int h) {
    if (l >= h)
        return;
    
    int p = split(a, l, h);
    quickSort(a, l, p - 1);
    quickSort(a, p + 1, h);
}

int split(int a[], int l, int h) {
    int p = a[h];
    int i = (l - 1);
  
    for (int j = l; j <= h - 1; j++) 
        if (a[j] < p) 
            swap(a[++i], a[j]);

    swap(a[i + 1], a[h]); 
    return (i + 1); 
}

void print (int a []) {
    int i;
    cout << "[";
    for (i = 0; a[i + 1]; i++)
        cout << a[i] << ", ";
    cout << a[i] << "]" << endl;
}
