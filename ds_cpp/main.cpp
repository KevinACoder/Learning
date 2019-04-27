#include "comm.h"

extern void lfu_cache_test();
extern void lru_cache_test();

int main(){
    lfu_cache_test();
    lru_cache_test();

    return 0;
}