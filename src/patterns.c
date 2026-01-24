#include <stdio.h>

void pattern1(int n) { for (int i = 1; i <= n; i++) { printf("* * * * *\n"); }
}

void pattern2(int n) { for (int i = 1; i <= n; i++) { for (int j = 1; j <= i;
	j++) { printf("* "); } printf("\n"); } }


void pattern3(int n) { for (int i = 1; i <= n; i++) {
		for (int j = 1; j <= i; j++) {
			printf("%d ", j);
		}
		printf("\n");
	}
}


void pattern4(int n) {
	for (int i = 1; i <= n; i++) {
		for (int j = 1; j <= i; j++) {
			printf("%d ", i);
		}
		printf("\n");
	}
}


void pattern5(int n) {
	for (int i = n; i >= 1; i--) {
		for (int j = 1; j <= i; j++) {
			printf("* ");
		}
		printf("\n");
	}
}

void pattern6(int n) {
	for (int i = n; i >= 1; i--) {
		for (int j = 1; j <= i; j++) {
			printf("%d ", i);
		}
		printf("\n");
	}
}

void pattern7(int n) {
	for (int i = 1; i <= n; i++) {
		int spaces = n - i;
		int stars = 2 * i - 1;
		for (int s = 1; s <= spaces; s++) {
			printf(" ");
		}
		for (int j = 1; j<=stars; j++) {
			printf("*");
		}
			printf("\n");
	}
}



void pattern8(int n) {
	for (int i = n; i >= 1; i--) {
		int spaces = n - i;
		int stars = 2 * i - 1;
		for (int s = 1; s <= spaces; s++) {
			printf(" ");
		}
		for (int j = 1; j<=stars; j++) {
			printf("*");
		}
			printf("\n");
	}
}




int main(){
	int n = 5;
	printf("Pattern 1:\n");
	pattern1(n);

	printf("Pattern 2:\n");
	pattern2(n);
	printf("Pattern 3:\n");
	pattern3(n);
	printf("Pattern 4:\n");
	pattern4(n);

	printf("Pattern 5:\n");
	pattern5(n);
	printf("Pattern 6:\n");
	pattern6(n);

	printf("Pattern 7:\n");
	pattern7(n);

	printf("Pattern 8:\n");
	pattern8(n);
}
