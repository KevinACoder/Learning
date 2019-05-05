#include "trie.h"
#include "comm.h"

#define MEM_FAIL "fail to allocate memory"

static trie_node *new_trie_node() {
	trie_node *n = malloc(sizeof(trie_node));
	if (!n)
		err_exit(MEM_FAIL);
	
	n->is_word = false;
	memset(n->children, 0, sizeof(trie_node *)*KEY_SET_NUM);

	return n;
}

static void del_trie_node(trie_node *n) {
	//delet all childrens first
	for (int i = 0; i < KEY_SET_NUM; i++) {
		if (n->children[i]) {
			del_trie_node(n->children[i]);
			n->children[i] = NULL;
		}	
	}
	free(n);
}

void init_trie(trie *t) {
	t->root = new_trie_node();
}

void reset_trie(trie *t) {
	del_trie_node(t->root);
	t->root = NULL;
}

void insert_trie(trie *t, const char *word) {
	trie_node *cur = t->root;
	if (!cur)
		return;

	int word_len = strlen(word);
	for (int i = 0; i < word_len; i++) {
		int ix = word[i] - 'a';
		if (!cur->children[ix])
			cur->children[ix] = new_trie_node();

		cur = cur->children[ix];
	}
	cur->is_word = true;
}

static trie_node * find(trie *t, const char *prefix) {
	trie_node *cur = t->root;

	int word_len = strlen(prefix);
	for (int i = 0; i < word_len; i++) {
		cur = cur->children[prefix[i] - 'a'];
		if (!cur)
			break;
	}
	return cur;
}

bool search_trie(trie *t, const char *word) {
	trie_node *cur = find(t, word);
	return cur && cur->is_word;
}

bool start_with_trie(trie *t, const char *prefix) {
	return find(t, prefix) != NULL;
}

void test_tire() {
	const char *word = "banan";
	INIT_TRIE(t);
	insert_trie(&t, word);
	printf("has prefix %d", start_with_trie(&t, "ban"));
	FREE_TRIE(t);
}