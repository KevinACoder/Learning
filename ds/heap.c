#include "heap.h"

static int parent(int i){
    return (i-1)/2;
}

static int left(int i){
    return 2*i + 1;
}

static int right(int i){
    return 2*i + 2;
}

void init_heap(heap_t *h, int type){
    h->type = type;
    init_vec(&h->v, INIT_HEAP_SIZE);
    if(type == MIN_HEAP)
        h->cmp = less_val_cmp;
    else
        h->cmp = more_val_cmp;
}

static void heapify(heap_t *h, int i){
    int l = left(i),
        r = right(i);

    int smallest = i;
    //check if child element is not in order with i
    if(l < h->v.len && h->cmp(h->v.items[l], h->v.items[i]))
        smallest = l;
    else if(r < h->v.len && h->cmp(h->v.items[r], h->v.items[i]))
        smallest = r;

    //recursively heapify until in order
    if(smallest != i){
        swap(&h->v.items[i], &h->v.items[smallest]);
        heapify(h, smallest);
    }
}

void insert_heap(heap_t *h, void *item){
    push_back_vec(&h->v, item); 
    int i = h->v.len - 1; //append new element in back

    //swap element with its parent if they are not in order
    while(i > 0 && !h->cmp(h->v.items[parent(i)], h->v.items[i])){
        swap(&h->v.items[parent(i)], &h->v.items[i]);
        i = parent(i);
    }
}

void *get_end(heap_t *h){
    if(total_vec(&h->v) <= 0){
        return NULL;
    } else {
        return h->v.items[0];
    }
}

void *extract_end(heap_t *h){
    void *ret = get_end(h);//get root element
    if(!ret){
        return NULL;
    }

    pop_front_vec(&h->v);//remove root
    heapify(h, 0);//re-heapify
    return ret;
}

void reset_heap(heap_t *h){
    reset_vec(&h->v);
}

void heap_test(){
    INIT_HEAP(h);
    insert_heap(&h, (void *)3);
    insert_heap(&h, (void *)2);
    insert_heap(&h, (void *)1);
    printf("min %d\n", (int)extract_end(&h));
    printf("min %d\n", (int)extract_end(&h));
    insert_heap(&h, (void *)1);
    printf("min %d\n", (int)extract_end(&h));
    printf("min %d\n", (int)extract_end(&h));
    FREE_HEAP(h);
}

void init_median_finder(median_finder_t *mf){
    init_heap(&mf->left, MAX_HEAP);
    init_heap(&mf->right, MIN_HEAP);
}

void add_num(median_finder_t *mf, int num){
    heap_t *left = &mf->left,
           *right = &mf->right;
    //add num and make sure left heap is longer
    if(total_heap(left) == 0 || num < (int)get_end(left)){
        insert_heap(left, (void *)(long)num);
    }else{
        insert_heap(right, (void *)(long)num);
    }

    //balance left and right heap, is left heap is shorter
    // move element from right to left
    //if left heap has 2 more elements than right
    // move from left to right
    if(total_heap(left) < total_heap(right)){
        insert_heap(left, get_end(right));
        extract_end(right);
    }else if (total_heap(left) - total_heap(right) >= 2){
        insert_heap(right, get_end(left));
        extract_end(left);
    }
}

float find_median(median_finder_t *mf){
    heap_t *left = &mf->left,
           *right = &mf->right;
    if(get_end(left) == NULL)
        return 0.0;

    //odd
    if(total_heap(left) > total_heap(right)){
        return (float)(long)get_end(left);
    }else{ //even
        return ((float)(long)get_end(left) + 
            (float)(long)get_end(right))/2;
    }
}

int total_heap(heap_t *h){
    return total_vec(&h->v);
}

void median_finder_test(){
    median_finder_t mf;
    init_median_finder(&mf);
    add_num(&mf, 1);
    add_num(&mf, 2);
    printf("median %f\n", find_median(&mf));
    add_num(&mf, 3);
    printf("median %f\n", find_median(&mf));
}