#ifndef VECTOR_H_
#define VECTOR_H_

#define INIT_CAP   10
#define FALSE       0
#define TRUE        1

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

void test_vec();

#endif