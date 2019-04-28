#ifndef HEAP_H
#define HEAP_H

#include "vector.h"
#include <stdbool.h>
#include "comm.h"

#define MIN_HEAP        0
#define MAX_HEAP        1
#define INIT_HEAP_SIZE 10
#define INIT_HEAP(h)    heap_t h; init_heap(&h, MIN_HEAP)
#define FREE_HEAP(h)    reset_heap(&h);

typedef struct heap_t{
    vector      v;
    int         type; //max heap or min heap
    cmp_func    cmp; //cmp func ptr
}heap_t;

void init_heap(heap_t *h, int type);
void insert_heap(heap_t *h, void *item);
void *extract_end(heap_t *h);
void *get_end(heap_t *h);
void reset_heap(heap_t *h);
int total_heap(heap_t *h);

void heap_test();

typedef struct median_finder_t{
    heap_t  left, right;
}median_finder_t;

void init_median_finder(median_finder_t *mf);
void add_num(median_finder_t *mf, int num);
float find_median(median_finder_t *mf);
void reset_median_finder();

void median_finder_test();

#endif