#include "vector.h"
#include "comm.h"
#include "hash_tbl.h"
#include "tree.h"
#include "sort.h"
#include "heap.h"

int main(){
    test_vec();
    test_hash_tbl();
    test_bst();
    test_sort();
    heap_test();
    median_finder_test();
    return 0;
}