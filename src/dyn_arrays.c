#include <stdio.h>
#include <stdlib.h>

#include "../data_structures/vector.h"

int main() {

  Vector *vec = vector_create();
  printf("Vector created \n");

  int runs = rand() % 1000;
  printf("Iterating for %d times\n", runs);
  for (int i = 1; i <= runs; i++) {
    vector_push(vec, rand() % 10000);
  }

  for (int j = 1; j <= runs; j++) {
    int val = -1;

    size_t idx = (size_t)(rand() % runs - 1);
    vector_get(vec, idx, &val);
    printf("Value fetched for index %zu, value: %d\n", idx, val);
  }

  vector_destroy(vec);

  return 1;
}
