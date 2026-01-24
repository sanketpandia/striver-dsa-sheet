// Resizable array
// Operations: push, pop, get, set, size
// Learn: Dynamic memory, realloc, amortized complexity
#include "vector.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

struct Vector {
  int *data;
  size_t size;
  size_t capacity;
};

#define INITIAL_CAPACITY 2
#define GROWTH_FACTOR 2

static bool vector_resize_required(Vector *vec) {
  return vec->size >= vec->capacity;
}

static bool vector_grow(Vector *vec) {
  size_t new_cap = GROWTH_FACTOR * vec->capacity;
  printf("Resizing from %zu to %zu\n", vec->capacity, new_cap);
  int *new_data = realloc(vec->data, new_cap * sizeof(int));
  if (new_data == NULL) {
    return false;
  }
  vec->data = new_data;
  vec->capacity = new_cap;
  return true;
}

static bool vector_shrink_possible(Vector *vec) {
  return vec->size < (vec->capacity / GROWTH_FACTOR);
}

static bool vector_shrink(Vector *vec) {
  // If Size is half of capacity we shrink
  size_t new_size = vec->capacity / GROWTH_FACTOR;

  int *old_ptr = vec->data;
  int *new_data = realloc(vec->data, new_size * sizeof(int));
  if (new_data == NULL) {
    return false;
  }
  vec->data = new_data;
  vec->capacity = new_size;
  return true;
}

bool vector_push(Vector *vec, int val) {
  if (vector_resize_required(vec)) {
    bool gr = vector_grow(vec);
    if (!gr) {
      printf("Unable to grow vector\n");
      return false;
    }
  }

  int *data = vec->data;
  data[vec->size] = val;
  vec->size++;
  return true;
}

Vector *vector_create_with_size(size_t initial_capacity) {
  Vector *vec = malloc(sizeof(Vector));

  if (vec == NULL) {
    return NULL;
  }

  vec->data = malloc(initial_capacity * sizeof(int));
  if (vec->data == NULL) {
    free(vec);
    return NULL;
  }
  vec->size = 0;
  vec->capacity = initial_capacity;
  return vec;
}

Vector *vector_create() { return vector_create_with_size(INITIAL_CAPACITY); }

bool vector_get(Vector *vec, size_t idx, int *target_ptr) {
  if (idx >= vec->size) {
    printf("Index out of bounds\n");
    return false;
  }
  *target_ptr = vec->data[idx];
  return true;
}

bool vector_delete(Vector *vec, size_t idx) {

  if (idx >= vec->size) {
    return false;
  }
  for (int i = idx; i < vec->size - 1; i++) {
    vec->data[i] = vec->data[i + 1];
  }
  vec->size--;

  if (vector_shrink_possible(vec)) {
    bool dn = vector_shrink(vec);
  }
  return true;
}

void vector_destroy(Vector *vec) {

  free(vec->data);
  free(vec);
}
