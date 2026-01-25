#include <stdio.h>
#include <stdlib.h>

void reverse_array(int n[], int len) {

	int start = 0;
	int end = len - 1;

	for(; start < end;){
		int temp = n[start];
		n[start] = n[end];
		n[end] = temp;
		start++;
		end--;
	}

}

void reverse_array_recursion(int arr[], int start, int end){
	if (end < start) {
		return;
	}

	int temp = arr[start];
	arr[start] = end;
	arr[end] = temp;

	reverse_array_recursion(arr, start + 1, end - 1);
}

int main() {
	int arr [] = {1,2,3,4,5,6,7,8,9,10};

	size_t len = sizeof(arr) / sizeof(int);

	reverse_array(arr, len);

	for(int i = 1; i < len; i++) {
		printf("%d ", arr[i]);
	}

		printf("\n " );
	reverse_array_recursion(arr, 0, len-1);

	for(int i = 1; i < len; i++) {
		printf("%d ", arr[i]);
	}

	return 1;
}
