#include "vector.h"
#include "comm.h"

void init_vec(vector *v, int cap){
    v->cap = cap;
    v->len = 0;
    v->items = calloc(v->cap, sizeof(void *));
    if(!v->items)
        err_exit(MEM_ERR);
}

static void resize_vec(vector *v, int new_cap){
    v->cap = new_cap;
    v->items = realloc(v->items, sizeof(void *)*v->cap);
    if(!v->items)
        err_exit(MEM_ERR);
}

void push_back_vec(vector *v, void *item){
    if(v->len == v->cap)
        resize_vec(v, v->cap*2);
    v->items[v->len++] = item;
}

void print_vec(vector *v){
    for(int i = 0; i < v->len; i++){
        printf("%d ", (int)v->items[i]);
    }
    printf("\n");
}

void push_front_vec(vector *v, void *item){
    if(v->len == v->cap)
        resize_vec(v, v->cap*2);
    memmove(v->items + 1, v->items, v->len*sizeof(void *));
    v->items[0] = item;
    v->len++;
}

void *pop_back_vec(vector *v){
    if(v->len < 1)
        return NULL;
    int last = v->len-1;
    void *ret = v->items[last];
    v->items[last] = NULL;
    --v->len;
    if(v->len < v->cap/3)
        resize_vec(v, v->cap/2);
    return ret;
}

void *pop_front_vec(vector *v){
    if(v->len < 1)
        return NULL;
    void *ret = v->items[0];
    v->len--;
    memmove(v->items, v->items + 1, v->len*sizeof(void *));
    if(v->len < v->cap/3)
        resize_vec(v, v->cap/2);
    return ret;    
}

void insert_vec(vector *v, void *item, int ix){
    if(ix <= 0){
        push_front_vec(v, item);
    }else if(ix >= v->len){
        push_back_vec(v, item);
    }else{
        if(v->len == v->cap)
            resize_vec(v, v->cap*2);
        memmove(v->items + ix + 1, v->items + ix, (v->len - ix)*sizeof(void *));
        v->len++;
        v->items[ix] = item;
    }
}

void *del_at_vec(vector *v, int ix){
    if(v->len == 0 || v->len <= ix || ix < 0)
        return NULL;
    else if(v->len - 1 == ix){
        return pop_back_vec(v);
    }else if(ix == 0){
        return pop_front_vec(v);
    }

    void *ret = v->items[ix];
    memmove(v->items + ix, v->items + ix + 1, (v->len - ix - 1)*sizeof(void *));
    v->len--;
    if(v->len < v->cap/3)
        resize_vec(v, v->cap/2);
    return ret;
}

void reset_vec(vector *v){
    free(v->items);
    v->cap = 0;
    v->len = 0;
}

void test_vec(){
    vector v;
    init_vec(&v, INIT_CAP);
    for(long i = 0; i < 20; i++){
        push_back_vec(&v, (void *)(i));
    }
    for(int i = 0; i < 3; i++){
        pop_front_vec(&v);
        pop_back_vec(&v);
    }
    insert_vec(&v, (void *)3, 3);
    del_at_vec(&v, 3);
    print_vec(&v);

    reset_vec(&v);
}