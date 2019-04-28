#include "sort.h"

bool less_val_cmp(void *lhs, void *rhs){
    if((int)lhs <= (int)rhs){
        return true;
    }else{
        return false;
    }
}

bool more_val_cmp(void *lhs, void *rhs){
    if((int)lhs >= (int)rhs){
        return true;
    }else{
        return false;
    }
}

void bubble_sort_vec(vector *v, cmp_func cmp){
    bool do_swap = false;
    for(int i = 0; i < v->len - 1; i++){
        for(int j = v->len - 1; j > 0; j--){
            if(!cmp(v->items[j-1], v->items[j])){
                swap(&v->items[j-1], &v->items[j]);
                do_swap = true;
            }
        }

        if(!do_swap)
            break;
    }
}

void insert_sort_vec(vector *v, cmp_func cmp){
    for(int i = 1; i < v->len; i++){
        void *key = v->items[i];
        int j = i - 1;

        while(j >= 0 && cmp(key, v->items[j])){
            v->items[j+1] = v->items[j];
            j--;
        }
        v->items[j+1] = key;
    }
}

void select_sort_vec(vector *v, cmp_func cmp){
    for(int i = 0; i < v->len; i++){
        int min_idx = i; 
        void *min_key = v->items[i];
        //select the min element in 'unorder' region
        for(int j = i; j < v->len; j++){
            if(cmp(v->items[j], min_key)){
                min_idx = j;
                min_key = v->items[j];
            }
        }
        //put min element in the back of 'ordered region'
        swap(&v->items[i], &v->items[min_idx]);
    }
}

static int partition(vector *v, cmp_func cmp, 
    int low, int high){
    void *pivot = v->items[high];
    int i = low;

    for(int j = low; j < high; j++){
        if(!cmp(pivot, v->items[j])){
            swap(&v->items[i], &v->items[j]);
            i++;
        }
    }
    swap(&v->items[i], &v->items[high]);
    return i;
}

static void quick_sort_vec_helper(vector *v, cmp_func cmp, 
    int low, int high){
    if(low < high){
        int pi = partition(v, cmp, low, high);
        quick_sort_vec_helper(v, cmp, low, pi-1);
        quick_sort_vec_helper(v, cmp, pi+1, high);
    }
}

void quick_sort_vec(vector *v, cmp_func cmp){
    quick_sort_vec_helper(v, cmp, 0, v->len-1);
}

static void merge(vector *v, cmp_func cmp, int low, int mid,
    int high){
    int llen = mid - low + 1,
        rlen = high - mid;
    void *L[llen], *R[rlen];

    //copy elements to left and right set
    memmove(L, v->items + low, sizeof(void*)*llen);
    memmove(R, v->items + mid + 1, sizeof(void*)*rlen);

    //merge two set
    int i = 0, j = 0, idx = low;
    while(i < llen && j < rlen){
        if(cmp(L[i], R[j])){
            v->items[idx++] = L[i++];
        }else{
            v->items[idx++] = R[j++];
        }
    }

    //copy the remain set after merge
    if(i < llen){
        memmove(&v->items[idx], &L[i], sizeof(void*)*(llen - i));
    }else if(j < rlen){
        memmove(&v->items[idx], &R[j], sizeof(void*)*(rlen - j));
    }
}

static void merge_sort_vec_helper(vector *v, cmp_func cmp,
    int low, int high){
    if(low < high){
        int mid = low + (high - low)/2;
        merge_sort_vec_helper(v, cmp, low, mid);
        merge_sort_vec_helper(v, cmp, mid+1, high);
        merge(v, cmp, low, mid, high);
    }
}

void merge_sort_vec(vector *v, cmp_func cmp){
    merge_sort_vec_helper(v, cmp, 0, v->len - 1);
}

void shell_sort_vec(vector *v, cmp_func cmp){
    //for each kind of iter gap
    for(int gap = v->len/2; gap > 0; gap /= 2){
        //for each element in the first gap region
        for(int i = gap; i < v->len; i += 1){
            void *key = v->items[i];
            int j = i;
            //exchange with 'gap' far element
            for(; j >= gap && cmp(key, v->items[j - gap]); j -= gap){
                v->items[j] = v->items[j-gap];
            }
            v->items[j] = key;
        }
    }
}

void test_sort(){
    int nums[] = {3, 4, 1, 2, 6, 5};
    INIT_VEC(data);
    set_data_vec(&data, nums, ITEMS_OF(nums));
    printf("sort...\n");
    print_vec(&data);
    //bubble_sort_vec(&data, less_val_cmp);
    //insert_sort_vec(&data, less_val_cmp);
    //select_sort_vec(&data, less_val_cmp);
    //quick_sort_vec(&data, less_val_cmp);
    //merge_sort_vec(&data, less_val_cmp);
    shell_sort_vec(&data, less_val_cmp);
    print_vec(&data);
    FREE_VEC(data);
}