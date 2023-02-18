package lrucache

import (
	"errors"
	"log"
)

// TODO(xenowits): Figure out how make for maps work
//  Also, what happens a lock is held by a function and the function calls another function which
//  too needs mutual exclusion access to the same data structure?

// LLNode represents a (doubly) linked list node.
type LLNode struct {
	left  *LLNode // Can rename this to previous (idiomatic)
	right *LLNode // Can rename this to next (idiomatic)
	key   int
	val   int
}

func NewCache(maxSize int) cache {
	return cache{
		maxSize: maxSize,
		vals:    make(map[int]*LLNode, maxSize), // map[key]pointer to linked list node
	}
}

// cache consists of a doubly linked list with the least recently used at the left end
// and the most recently used at the rightmost end. It keeps a map that contains pointer to a
// linked list node given the key.
type cache struct {
	// mu      sync.Mutex // TODO(xenowits): Use RWMutex here.
	maxSize int
	head    *LLNode
	tail    *LLNode
	vals    map[int]*LLNode
}

func (c cache) Get(key int) (int, error) {
	if len(c.vals) == 0 {
		return 0, errors.New("empty cache")
	}

	node, ok := c.vals[key]
	if !ok {
		return 0, errors.New("key not found")
	}

	c.moveToEndofLL(node)

	return node.val, nil
}

func (c *cache) Set(key, val int) {
	if len(c.vals) == 0 {
		// Create the first LL node
		node := &LLNode{
			key: key,
			val: val,
		}
		c.head = node
		c.tail = node

		// Set first map
		c.vals[key] = node

		return
	}

	for k, _ := range c.vals {
		if k == key { // If such a key is already present, override its value
			c.vals[key].val = val
			c.moveToEndofLL(c.vals[key])
			return
		}
	}

	// This key is new, it doesn't already exist in the cache.
	// Do we have sufficient space? In other words, is the cache full?
	if len(c.vals) == c.maxSize { // OMG, cache's full 😱
		// Let's evict a key based on LRU policy
		c.evict()
	}

	c.vals[key].val = val
	c.moveToEndofLL(c.vals[key])
}

// moveToEndofLL moves the given node to the end of the cache's linked list.
func (c *cache) moveToEndofLL(node *LLNode) {
	if len(c.vals) == 0 || len(c.vals) == 1 {
		// No LL. Or a single node LL. Simply return, nothing much to do here.
		return
	}

	if len(c.vals) == 2 { // Two node LL.
		if node == c.tail {
			// Node is already at the end of the LL.
			return
		}

		// So node is at the front. First node. Let's swap these two nodes & we're done.
		// ie, currently it is head -> tail. We need to do tail -> head.
		tmphead := c.head
		c.head = c.tail
		c.tail = tmphead

		// Adjust left right pointers for both
		c.head.right = c.tail
		c.head.left = nil
		c.tail.right = nil
		c.tail.left = c.head
	}

	// 3-node LL onwards.
	if node == c.head {
		// Case 1: Node is at front. First node. Head node.
		// head/node -> nextToHead -> ........ -> tail
		// nextToHead -> ...... -> tail -> node
		// Advance head to the next node in the LL.
		c.head = c.head.right // New head is current head's right
		c.head.left = nil     // Detach new head from current head

		// Point tail to the old head (now the tail)
		c.tail.right = node
		node.right = nil
		node.left = c.tail

		c.tail = node // Set this node as the tail
	} else if node == c.tail {
		// Case 2: Node is at the end. Last node. Tail node.
		// No need to do anything. It's already most recently used.
	} else {
		// Case 3: Node is somewhere in between two nodes. Consider this node to be `B` in the figure below.
		// head -> ..... -> A -> B -> C -> .... -> tail
		// head -> ..... -> A -> C -> .... -> tail -> B. We need to get to this configuration.
		nodesLeft := node.left       // A
		nodesRight := node.right     // C
		nodesLeft.right = nodesRight // A -> C

		c.tail.right = node
		node.right = nil
		node.left = c.tail

		c.tail = nodesLeft
	}
}

// evict evicts an element based on LRU policy.
func (c *cache) evict() {
	leastRecentlyUsedNode := c.head // As first node in the LL
	// head -> nextToHead -> ...... -> tail
	//         head       -> ...... -> tail. Make current head vanish
	log.Println("Evicting least recently used node", leastRecentlyUsedNode.key, leastRecentlyUsedNode.val)
	tmphead := c.head
	tmphead.right = nil

	c.head = c.head.right

	// Also erase key from hashmap
	delete(c.vals, tmphead.key)
}
