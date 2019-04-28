#ifndef COMM_H_
#define COMM_H_

#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#define FALSE       0
#define TRUE        1

#define SMALL       0
#define BIG         1
#define EQUAL       2
#define ITEMS_OF(arr)  sizeof(arr)/sizeof(arr[0])     

void err_exit(const char *msg);
void swap(void **rhs, void **lhs);

typedef bool (*cmp_func) (void *, void *);
bool less_val_cmp(void *lhs, void *rhs);
bool more_val_cmp(void *lhs, void *rhs);

#define MEM_ERR  "Fail to allocate memory"

#endif