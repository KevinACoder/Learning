#include "comm.h"

void err_exit(const char *msg){
    printf("%s\n", msg);
    exit(1);
}

void swap(void **rhs, void **lhs){
    void *tmp = *rhs;
    *rhs = *lhs;
    *lhs = tmp;
}