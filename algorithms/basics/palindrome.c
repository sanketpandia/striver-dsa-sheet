#include <stdbool.h>
#include <stdio.h>
#include <string.h>

bool is_palindrome(char str[], int start, int end) {

  if (start > end) {
    return true;
  }

  if (str[start] != str[end]) {
    printf("Palindrome mismatched!\n");
    return false;
  }

  return is_palindrome(str, start + 1, end - 1);
}

int main() {

  char str[] = {'s', 'a', 'm', 'a', 's', '\0'};
  int len = 5;
  printf("String: %s\nIsPalindrome=%d", str, is_palindrome(str, 0, strlen(str) - 1));

  return 1;
}
