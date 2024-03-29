// Package trie contains a primitive implementation of the Trie data structure.
//
// Copyright 2017 Aleksandr Bezobchuk.

// Trie implmeneted from this github repo
// Use this link to read
// https://github.com/alexanderbez/go-trie

package trie

import (
    "container/list"
    "sync"
)

// Bytes reflects a type alias for a byte slice
type Bytes []byte

// trieNode implements a node that the Trie is composed of. Each node contains
// a symbol that a key can be composed of unless the node is the root. The node
// has a collection of children that is represented as a hashmap, although,
// traditionally an array is used to represent each symbol in the given
// alphabet. The node may also contain a value that indicates a possible query
// result.
//
// TODO: Handle the case where the value given is a dummy value which can be
// nil. Perhaps it's best to not store values at all.
type trieNode struct {
    children map[byte]*trieNode
    symbol   byte
    value    []byte
    root     bool
}

// Trie implements a thread-safe search tree that stores byte key value pairs
// and allows for efficient queries.
type Trie struct {
    rw   sync.RWMutex
    root *trieNode
    size int
}

// NewTrie returns a new initialized empty Trie.
func NewTrie() *Trie {
    return &Trie{
        root: &trieNode{root: true, children: make(map[byte]*trieNode)},
        size: 1,
    }
}

func newNode(symbol byte) *trieNode {
    return &trieNode{children: make(map[byte]*trieNode), symbol: symbol}
}

// Size returns the total number of nodes in the trie. The size includes the
// root node.
func (t *Trie) Size() int {
    t.rw.RLock()
    defer t.rw.RUnlock()
    return t.size
}

// Insert inserts a key value pair into the trie. If the key already exists,
// the value is updated. Insertion is performed by starting at the root
// and traversing the nodes all the way down until the key is exhausted. Once
// exhausted, the currNode pointer should be a pointer to the last symbol in
// the key and reflect the terminating node for that key value pair.
func (t *Trie) Insert(key, value Bytes) {
    t.rw.Lock()
    defer t.rw.Unlock()

    currNode := t.root

    for _, symbol := range key {
        if currNode.children[symbol] == nil {
            currNode.children[symbol] = newNode(symbol)
        }

        currNode = currNode.children[symbol]
    }

    // Only increment size if the key value pair is new, otherwise we consider
    // the operation as an update.
    if currNode.value == nil {
        t.size++
    }

    currNode.value = value
}

// Search attempts to search for a value in the trie given a key. If such a key
// exists, it's value is returned along with a boolean to reflect that the key
// exists. Otherwise, an empty value and false is returned.
func (t *Trie) Search(key Bytes) (Bytes, bool) {
    t.rw.RLock()
    defer t.rw.RUnlock()

    currNode := t.root

    for _, symbol := range key {
        if currNode.children[symbol] == nil {
            return nil, false
        }

        currNode = currNode.children[symbol]
    }

    return currNode.value, true
}

// GetAllKeys returns all the keys that exist in the trie. Keys are retrieved
// by performing a DFS on the trie where at each node we keep track of the
// current path (key) traversed thusfar and if that node has a value. If so,
// the full path (key) is appended to a list. After the trie search is
// exhausted, the final list is returned.
func (t *Trie) GetAllKeys() []Bytes {
    visited := make(map[*trieNode]bool)
    keys := []Bytes{}

    var dfsGetKeys func(n *trieNode, key Bytes)
    dfsGetKeys = func(n *trieNode, key Bytes) {
        if n != nil {
            pathKey := append(key, n.symbol)
            visited[n] = true

            if n.value != nil {
                fullKey := make(Bytes, len(pathKey))

                // Copy the contents of the current path (key) to a new key so
                // future recursive calls will contain the correct bytes.
                copy(fullKey, pathKey)

                // Append the path (key) to the key list ignoring the first
                // byte which is the root symbol.
                keys = append(keys, fullKey[1:])
            }

            for _, child := range n.children {
                if _, ok := visited[child]; !ok {
                    dfsGetKeys(child, pathKey)
                }
            }
        }
    }

    dfsGetKeys(t.root, Bytes{})
    return keys
}

// GetAllValues returns all the values that exist in the trie. Values are
// retrieved by performing a BFS on the trie where at each node we examine if
// that node has a value. If so, that value is appended to a list. After the
// trie search is exhausted, the final list is returned.
func (t *Trie) GetAllValues() []Bytes {
    queue := list.New()
    visited := make(map[*trieNode]bool)
    values := []Bytes{}

    queue.PushBack(t.root)

    for queue.Len() > 0 {
        element := queue.Front()
        queue.Remove(element)

        node := element.Value.(*trieNode)
        visited[node] = true

        if node.value != nil {
            values = append(values, node.value)
        }

        for _, child := range node.children {
            _, ok := visited[child]
            if !ok {
                queue.PushBack(child)
            }
        }
    }

    return values
}

// GetPrefixKeys returns all the keys that exist in the trie such that each key
// contains a specified prefix. Keys are retrieved by performing a DFS on the
// trie where at each node we keep track of the current path (key) and prefix
// traversed thusfar. If a node has a value the full path (key) is appended to
// a list. After the trie search is exhausted, the final list is returned.
func (t *Trie) GetPrefixKeys(prefix Bytes) []Bytes {
    visited := make(map[*trieNode]bool)
    keys := []Bytes{}

    if len(prefix) == 0 {
        return keys
    }

    var dfsGetPrefixKeys func(n *trieNode, prefixIdx int, key Bytes)
    dfsGetPrefixKeys = func(n *trieNode, prefixIdx int, key Bytes) {
        if n != nil {
            pathKey := append(key, n.symbol)

            if prefixIdx == len(prefix) || n.symbol == prefix[prefixIdx] {
                visited[n] = true

                if n.value != nil {
                    fullKey := make(Bytes, len(pathKey))

                    // Copy the contents of the current path (key) to a new key
                    // so future recursive calls will contain the correct
                    // bytes.
                    copy(fullKey, pathKey)

                    keys = append(keys, fullKey)
                }

                if prefixIdx < len(prefix) {
                    prefixIdx++
                }

                for _, child := range n.children {
                    if _, ok := visited[child]; !ok {
                        dfsGetPrefixKeys(child, prefixIdx, pathKey)
                    }
                }
            }
        }
    }

    // Find starting node from the root's children
    if n, ok := t.root.children[prefix[0]]; ok {
        dfsGetPrefixKeys(n, 0, Bytes{})
    }

    return keys
}

// GetPrefixValues returns all the values that exist in the trie such that each
// key that corresponds to that value contains a specified prefix. Values are
// retrieved by performing a DFS on the trie where at each node we check if the
// prefix is exhausted or matches thusfar and the current node has a value. If
// the current node has a value, it is appended to a list. After the trie
// search is exhausted, the final list is returned.
func (t *Trie) GetPrefixValues(prefix Bytes) []Bytes {
    visited := make(map[*trieNode]bool)
    values := []Bytes{}

    if len(prefix) == 0 {
        return values
    }

    var dfsGetPrefixValues func(n *trieNode, prefixIdx int)
    dfsGetPrefixValues = func(n *trieNode, prefixIdx int) {
        if n != nil {
            if prefixIdx == len(prefix) || n.symbol == prefix[prefixIdx] {
                visited[n] = true

                if n.value != nil {
                    values = append(values, n.value)
                }

                if prefixIdx < len(prefix) {
                    prefixIdx++
                }

                for _, child := range n.children {
                    if _, ok := visited[child]; !ok {
                        dfsGetPrefixValues(child, prefixIdx)
                    }
                }
            }
        }
    }

    // Find starting node from the root's children
    if n, ok := t.root.children[prefix[0]]; ok {
        dfsGetPrefixValues(n, 0)
    }

    return values
}
