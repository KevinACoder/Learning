#ifndef TREE_H_
#define TREE_H_

typedef struct bst_node{
    struct bst_node *left;
    struct bst_node *right;
    void            *key;
}bst_node;

typedef struct bs_tree
{
    bst_node        *root;
    int             node_num;
}bs_tree;


#endif