#include "tree.h"
#include "comm.h"
#include <limits.h>

#define NULL_NODE INT_MAX

static bst_node *new_bst_node(void *key){
    bst_node *n = malloc(sizeof(bst_node));
    if(!n)
        err_exit(MEM_ERR);
    n->left = n->right = NULL;
    n->key = key;
    return n;
}

static void del_bst_node(bst_node **node_ptr){
    free(*node_ptr);
    *node_ptr = NULL;
}

int key_cmp(void *k1, void *k2){
    if((int)k1 < (int)k2){
        return SMALL;
    }else if((int)k1 > (int)k2){
        return BIG;
    }else{
        return EQUAL;
    }
}

void init_bst(bs_tree *t){
    t->root = NULL;
    t->node_num = 0;
    t->cmp = key_cmp;
}

bst_node *insert_bst_helper(bst_node *root, void *key, cmp_ptr cmp,
    int *node_num){
    bst_node *curr = root, 
             *parent = NULL;

    while(curr){
        parent = curr;
        switch(cmp(curr->key, key)){
            case SMALL:
                curr = curr->right;
            break;
            case BIG:
                curr = curr->left;
            break; 
            case EQUAL:
                return root;           
        }
    }

    if(!parent)
        parent = new_bst_node(key);
    else {
        switch(cmp(parent->key, key)){
            case SMALL:
                parent->right = new_bst_node(key);
                break;
            default:
                parent->left = new_bst_node(key);
        }
    }
    (*node_num)++;
    return parent;
}

void insert_bst(bs_tree *t, void *key){
    if(!t->root)
        t->root = insert_bst_helper(t->root, key, t->cmp, &t->node_num);
    else
        insert_bst_helper(t->root, key, t->cmp, &t->node_num);
}

void pre_order(bs_tree *t, vector *v){
    if(!t->root)
        return;

    INIT_VEC(queue);
    push_back_vec(&queue, t->root);

    while(total_vec(&queue) > 0){
        bst_node *n = pop_front_vec(&queue);
        push_back_vec(v, n->key);

        if(n->left)
            push_back_vec(&queue, n->left);
        if(n->right)
            push_back_vec(&queue, n->right);
    }

    FREE_VEC(queue);
}

static void pre_order_recur_helper(bst_node *root, vector *v){
    if(!root)
        return;
    
    push_back_vec(v, root->key);
    pre_order_recur_helper(root->left, v);
    pre_order_recur_helper(root->right, v);
}

void pre_order_recur(bs_tree *t, vector *v){
    pre_order_recur_helper(t->root, v);
}

static bst_node *insert_bst_recur_helper(bst_node *root, void *key, 
    cmp_ptr cmp, int *node_num){
    if(!root){
        (*node_num)++;
        return new_bst_node(key);
    }
        
    switch(cmp(root->key, key)){
        case SMALL:
            root->right = insert_bst_recur_helper(
                root->right, key, cmp, node_num);
        break;
        case BIG:
            root->left = insert_bst_recur_helper(
                root->left, key, cmp, node_num);
        break;
        case EQUAL:
        break;
    }
    return root;
}

void insert_bst_recur(bs_tree *t, void *key){
    t->root = insert_bst_recur_helper(t->root, key, t->cmp, &t->node_num);
}

static bst_node *post_order_recur_helper(bst_node *root, fun_ptr func){
    if(!root)
        return NULL;
    
    root->left = post_order_recur_helper(root->left, free);
    root->right = post_order_recur_helper(root->right, free);
    free(root);
    return NULL;
}

void reset_bst(bs_tree *t){
    t->root = post_order_recur_helper(t->root, free);
    t->node_num = 0;
}

void in_order(bs_tree *t, vector *v){
    if(!t->root)
        return;

    bst_node *curr = t->root;
    INIT_VEC(stack);

    while(curr || total_vec(&stack) > 0){
        while(curr){
            push_back_vec(&stack, curr);
            curr = curr->left;
        }

        curr = pop_back_vec(&stack);
        push_back_vec(v, curr->key);

        curr = curr->right;
    }

    FREE_VEC(stack);
}

static void serialize_helper(bst_node *root, vector *serie){
    if(!root){
        push_back_vec(serie, (void *)(long)NULL_NODE);
        return;
    }

    push_back_vec(serie, root->key);
    serialize_helper(root->left, serie);
    serialize_helper(root->right, serie);
}

void serialize(bs_tree *t, vector *serie){
    serialize_helper(t->root, serie);
}

static void de_serialize_helper(bst_node **root_ptr, vector *serie){
    int val = (int)pop_front_vec(serie);
    if(val == NULL_NODE)
        return;

    *root_ptr = new_bst_node((void *)(long)val);
    de_serialize_helper(&((*root_ptr)->left), serie);
    de_serialize_helper(&((*root_ptr)->right), serie);
}

void de_serialize(bs_tree *t, vector *serie){
    de_serialize_helper(&(t->root), serie);
}

void test_bst(){
    int nums[] = {3, 1, 5, 7, 6, 4, 2};
    int len = ITEMS_OF(nums);
    INIT_TREE(tree);
    INIT_VEC(serie);
    for(int i = 0; i < len; i++){
        insert_bst(&tree, (void *)(long)nums[i]);
        //insert_bst_recur(&tree, (void *)(long)nums[i]);
    }
    printf("node num: %d\n", tree.node_num);
    pre_order(&tree, &serie);
    //pre_order_recur(&tree, &serie);
    print_vec(&serie);

    REINIT_VEC(serie);
    in_order(&tree, &serie);
    print_vec(&serie);

    FREE_TREE(tree);

    INIT_TREE(tree2);
    INIT_TREE(tree3);
    INIT_VEC(serie2);
    for(int i = 0; i < len; i++){
        insert_bst(&tree2, (void *)(long)nums[i]);
        //insert_bst_recur(&tree, (void *)(long)nums[i]);
    }
    serialize(&tree2, &serie2);
    de_serialize(&tree3, &serie2);

    REINIT_VEC(serie);
    in_order(&tree3, &serie);
    print_vec(&serie);

    FREE_VEC(serie2);
    FREE_TREE(tree2);
    FREE_TREE(tree3);
    FREE_VEC(serie);
}