#ifndef HASH_TBL_H
#define HASH_TBL_H

#define INIT_BUCKET_NUM  40

typedef struct hash_node{
    char              *key;
    char              *val;
    struct hash_node  *next;
}hash_node;

typedef struct hash_tbl{
    int              bucket_num;
    int              node_num;
    hash_node        **buckets;
}hash_tbl;

void init_hash_tbl(hash_tbl *tbl);
void reset_hash_tbl(hash_tbl *tbl);
void insert_hash_tbl(hash_tbl *tbl, const char *k,
    const char *v);
char *search_hash_tbl(hash_tbl *tbl, const char *k);
void del_key_hash_tbl(hash_tbl *tbl, const char *k);

void test_hash_tbl();
#endif