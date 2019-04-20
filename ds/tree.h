#ifndef TREE_H_
#define TREE_H_

#include "vector.h"

#define INIT_TREE(tree) bs_tree tree;init_bst(&tree)
#define FREE_TREE(tree) reset_bst(&tree)

typedef int (*cmp_ptr)(void *, void *);
typedef void (*fun_ptr)(void *);

typedef struct bst_node{
    struct bst_node *left;
    struct bst_node *right;
    void            *key;
}bst_node;

typedef struct bs_tree
{
    bst_node        *root;
    int             node_num;
    cmp_ptr         cmp;
}bs_tree;

void init_bst(bs_tree *t);
void reset_bst(bs_tree *t);
void insert_bst(bs_tree *t, void *key);
void insert_bst_recur(bs_tree *t, void *key);
void pre_order(bs_tree *t, vector *v);
void pre_order_recur(bs_tree *t, vector *v);
void in_order(bs_tree *t, vector *v);
void serialize(bs_tree *t, vector *serie);
void de_serialize(bs_tree *t, vector *serie);

int key_cmp(void *k1, void *k2);

void test_bst();

#endif