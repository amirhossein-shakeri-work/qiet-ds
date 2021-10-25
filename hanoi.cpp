#include <iostream>
#include <stdexcept>

using namespace std;

void hanoi(int, int, int, int);
void printDisk(int, int, int);

int main () {
    cout << "Hanoi Recursive\n";
    hanoi(4, 1, 3, 2);
    return 0;
}

void hanoi (int n, int from, int to, int helper) {
    /* Throw exception if invalid n */
    if (n < 0 || from < 0 || to < 0 || helper < 0)
        return throw invalid_argument("Negative arg not allowed!");
    
    if (n == 1) {
        printDisk(n, from, to);
        return;
    }

    hanoi(n - 1, from, helper, to);
    printDisk(n, from, helper);
    hanoi(n - 1, helper, to, from);
}

void printDisk(int name, int from, int to){
    cout << "(disk " << name << ") {" << from << " --> " << to << "}\n";
}
