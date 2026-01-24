#include <stdio.h>

#define N 9

void name_print(int n) {
	printf("test\n");
	if (n == 0) return;

	name_print(n-1);
}

int factorial(int n) {
	if(n==0 || n == 1) return 1;

	return n * factorial(n - 1);
}

int main() {
	printf("======Recursion=======\n");

	// name_print(N);


	printf("Factorial of %d is %d\n", N, factorial(N));
	return 0;
}
