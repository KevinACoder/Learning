#ifndef VECTOR_H_
#define VECTOR_H_

#define INIT_CAP   10
#define INIT_VEC(vec) vector vec; init_vec(&vec, INIT_CAP)
#define FREE_VEC(vec) reset_vec(&(vec))
#define REINIT_VEC(vec) reset_vec(&(vec)); init_vec(&vec, INIT_CAP)

typedef struct vector{
    int    cap;
    int    len;
    void   **items; 
}vector;

void init_vec(vector *v, int cap);
void reset_vec(vector *v);
void push_back_vec(vector *v, void *item);
void push_front_vec(vector *v, void *item);
void print_vec(vector *v);
void *pop_back_vec(vector *v);
void *pop_front_vec(vector *v);
void insert_vec(vector *v, void *item, int ix);
void *del_at_vec(vector *v, int ix);
int total_vec(vector *v);

void test_vec();

#endif