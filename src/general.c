#include <stdio.h>

int square(int x) {
    return x * x;
}

int main(void) {
    int a = 5;
    int b = square(a);
    printf("%d\n", b);
    return 0;
}