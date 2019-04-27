#include "comm.h"

using namespace std;

typedef int key_t;
typedef int val_t;

struct cache_node_t
{
    key_t key;
    val_t val;
    int freq, tick; //freq is the times of visits, tick is the time of recent visit
    //less if is less freq or add before
    bool operator < (const cache_node_t& rhs) const {
        return freq < rhs.freq ||
            (freq == rhs.freq && tick < rhs.tick);
    }
};

class lfu_cache_t{
    int cur_tick, capacity; //currect time
    unordered_map<key_t, cache_node_t> kv;
    set<cache_node_t>                  hist; //visit queue

public:
    lfu_cache_t(int cap):capacity(cap){}
    int get(key_t k){
        if(kv.count(k) > 0){ //if key exist
            touch(kv[k]);
            return kv[k].val;
        }else{ //if key not valid
            return -1;
        }
    }

    void put(key_t k, val_t v){
        if(capacity <= 0) //invalid cache
            return;

        if(kv.count(k) > 0){ //update exist k-v
            kv[k].val = v;
        }else{
            //remove old node
            if(hist.size() == capacity){
                cache_node_t evict_node = *hist.cbegin();
                kv.erase(evict_node.key);
                hist.erase(evict_node);
            }

            //insert new node
            cache_node_t new_node{k, v, 0, cur_tick};
            kv[k] = new_node;
            hist.insert(new_node);
        }
        touch(kv[k]); //update node whatever
    }

private:
    void touch(cache_node_t &node){
        hist.erase(node); //temporary remove node
        node.freq++; //update visit times and last 
        node.tick = ++cur_tick; //visit time
        hist.insert(node); //add node again
    }
};

void lfu_cache_test(){
    lfu_cache_t cache( 2 /* capacity */ );
 
    cache.put(1, 1);
    cache.put(2, 2);
    cout<<cache.get(1);       // returns 1
    cache.put(3, 3);    // evicts key 2
    cout<<cache.get(2);       // returns -1 (not found)
    cout<<cache.get(3);       // returns 3.
    cache.put(4, 4);    // evicts key 1.
    cout<<cache.get(1);       // returns -1 (not found)
    cout<<cache.get(3);       // returns 3
    cout<<cache.get(4);       // returns 4
    cout<<endl;
}

typedef pair<key_t, val_t> kv_pair;

class lru_cache_t{
    int                           capacity;
    list<kv_pair>                 hist;
    unordered_map<key_t, 
    list<kv_pair>::iterator>      kv; 

public:
    lru_cache_t(int cap):capacity(cap){}

    void put(key_t k, val_t v) {
        auto it = kv.find(k);

        if(it != kv.end()){
            it->second->second = v; //update value
            hist.splice(hist.begin(), hist, it->second); //move visit record to front
        }else{
            if(hist.size() == capacity){
                auto evict_node = hist.back();
                kv.erase(evict_node.first); 
                hist.pop_back(); //remove the key in back
            }

            hist.emplace_front(k, v); //insert new key
            kv[k] = hist.begin();
        }
    }

    val_t get(key_t k){
        auto it = kv.find(k);
        if(it == kv.end()){
            return -1;
        }else{
            hist.splice(hist.begin(), hist, it->second);
            return it->second->second;
        }
    }
};

void lru_cache_test(){
    lru_cache_t cache( 2 /* capacity */ );
    
    cache.put(1, 1);
    cache.put(2, 2);
    cout<< cache.get(1);       // returns 1
    cache.put(3, 3);    // evicts key 2
    cout<< cache.get(2);       // returns -1 (not found)
    cache.put(4, 4);    // evicts key 1
    cout<< cache.get(1);       // returns -1 (not found)
    cout<<cache.get(3);       // returns 3
    cout<<cache.get(4);       // returns 4   
}