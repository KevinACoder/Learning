#ifndef SORT_H
#define SORT_H
#include <stdbool.h>
#include "vector.h"
#include "comm.h"

void bubble_sort_vec(vector *v, cmp_func cmp);
void insert_sort_vec(vector *v, cmp_func cmp);
void select_sort_vec(vector *v, cmp_func cmp);
void quick_sort_vec(vector *v, cmp_func cmp);
void merge_sort_vec(vector *v, cmp_func cmp);
void shell_sort_vec(vector *v, cmp_func cmp);
void bucket_sort_vec(vector *v, cmp_func cmp);
void radix_sort_vec(vector *v, cmp_func cmp);
void test_sort();

#endif