#ifndef VECTOR_H
#define VECTOR_H

#include <stdbool.h>
#include <stddef.h>

typedef struct Vector Vector;

Vector *vector_create();
Vector *vector_create_with_size(size_t capacity);

bool vector_push(Vector *vec, int val);
bool vector_delete(Vector *vec, size_t idx);
bool vector_get(Vector *vec, size_t idx, int *target_ptr);


void vector_destroy(Vector *vec);
#endif
