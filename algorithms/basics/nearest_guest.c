#include <stdio.h>


int find_max_dist_0(int arr[], size_t len){


	int ptr1 = 0;
	int ptr2 = 0;
	int maxDist = 0;

	for(int i = 0; i < len; i++){
		if(arr[i] == 0 ){
			maxDist = 1;
		}

		for
	}

	return maxDist;
}

int main(){
	int arr[10] = {0,1,0,0,0,1,0,1,1,0};

	find_max_dist_0(arr, 10);
	return 1;
}
