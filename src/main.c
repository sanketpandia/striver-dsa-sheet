#include <stdio.h>

int main(void) {
	printf("===== Striver DSA Series ====\n");

    int n = 0;
    char x = 'a';
    char str[100];
    printf("Enter a number, char and string: ");
    scanf("%d", &n);
    // Clear input buffer
    while (getchar() != '\n');
    scanf("%c", &x);
    scanf("%99s", str);
    printf("You entered string: %s\n", str);
    printf("You entered: %d & char: %c\n", n, x);
    return 0;
}
