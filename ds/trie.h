#ifndef TRIE_H
#define TRIE_H

#define INIT_TRIE(t) trie t; init_trie(&t)
#define FREE_TRIE(t) reset_trie(&t);
#define KEY_SET_NUM 26

typedef struct trie_node {
bool is_word;
struct trie_node *children[KEY_SET_NUM];
} trie_node;

typedef struct trie {
trie_node *root;
} trie;

void init_trie(trie *t);
void reset_trie(trie *t);
void insert_trie(trie *t, const char *word);
bool search_trie(trie *t, const char *word);
bool start_with_trie(trie *t, const char *prefix);

void test_tire();

#endif // !TRIE_H