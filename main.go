package main

import (
    "errors"
    "fmt"
    bst "github.com/wuno/datastructures/binarysearchtree"
    dict "github.com/wuno/datastructures/dictionary"
    "github.com/wuno/datastructures/graph"
    "github.com/wuno/datastructures/hashtable"
    "github.com/wuno/datastructures/heap"
    "github.com/wuno/datastructures/linkedlist"
    "github.com/wuno/datastructures/queue"
    "github.com/wuno/datastructures/set"
    "github.com/wuno/datastructures/stack"
    "github.com/wuno/datastructures/trie"
    "log"
)

var values = []string{"d", "b", "c", "e", "a"}
var data = []string{"delta", "bravo", "charlie", "echo", "alpha"}

func main() {
    runTrie()
}

func runTrie() {
    trie := trie.NewTrie()
    // Insert a key value pair into the trie
    trie.Insert([]byte("someKey"), []byte("someValue"))
    // Search for some value
    value, ok := trie.Search([]byte("someKey"))
    if ok {
        fmt.Printf("%v", value)
        fmt.Print("\n")
    }
    // Get all keys stored in the trie
    keys := trie.GetAllKeys()
    fmt.Printf("%v", keys)
    fmt.Print("\n")
    // Get all values stored in the trie
    values := trie.GetAllValues()
    fmt.Printf("%v", values)
    fmt.Print("\n")
    // Get all keys stored in the trie that contain a specific prefix
    keyByPrefix := trie.GetPrefixKeys([]byte("someKey"))
    fmt.Printf("%v", keyByPrefix)
    fmt.Print("\n")
    // Get all values stored in the trie who's corresponding keys contain a
    // specific prefix.
    valuesByPrefix := trie.GetPrefixValues([]byte("someValue"))
    fmt.Printf("%v", valuesByPrefix)
    fmt.Print("\n")
}

func runHeap() {
    h := heap.Heap{}
    h.Insert(32)
    h.Insert(22)
    h.Insert(3432)
    h.Insert(1)
    h.Print()
}

func runSet() {
    set := set.ItemSet{}
    set.Add(3)
    set.Add(4)
    set.Add(5)
    set.Add(2)
    set.Add(1)
    set.Add(6)
    fmt.Printf("%v", set.Has(1))
    fmt.Print("\n")
    fmt.Printf("%v", set.Items())
    fmt.Print("\n")
    fmt.Printf("%v", set.Size())
    set.Delete(1)
    fmt.Print("\n")
    fmt.Printf("%v", set.Size())
    fmt.Print("\n")
    fmt.Printf("%v", set.Items())
}

func runQueue() {
    q := queue.ItemQueue{}
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    fmt.Printf("%v", q.Front())
    fmt.Print("\n")
    fmt.Printf("%v", q.Size())
    fmt.Print("\n")
    q.Dequeue()
    q.Dequeue()
    fmt.Printf("%v", q.Size())
    fmt.Print("\n")
    fmt.Printf("%v", q.Front())
}

func runLL() {
    var ll linkedlist.ItemLinkedList
    ll.Append("first Entry")
    ll.Append("Second Entry")
    fmt.Printf("%v", ll.Size())
    fmt.Printf("%v", ll.Head())
}

func runHashTable() {
    dict := hashtable.ValueHashtable{}
    dict.Put(1123, "Yo Mama")
    dict.Put(1432, "Yo Girl")
    dict.Put(1232, "Yo Papa")
    val := dict.Get(1123)
    fmt.Print(val)
}

func runGraph() {
    var g graph.ItemGraph
    nA := graph.Node{"A"}
    nB := graph.Node{"B"}
    nC := graph.Node{"C"}
    nD := graph.Node{"D"}
    nE := graph.Node{"E"}
    nF := graph.Node{"F"}
    g.AddNode(&nA)
    g.AddNode(&nB)
    g.AddNode(&nC)
    g.AddNode(&nD)
    g.AddNode(&nE)
    g.AddNode(&nF)
    g.AddEdge(&nA, &nB)
    g.AddEdge(&nA, &nC)
    g.AddEdge(&nB, &nE)
    g.AddEdge(&nC, &nE)
    g.AddEdge(&nE, &nF)
    g.AddEdge(&nD, &nA)
    g.String()
}

func runDictionary() {
    dict := dict.ValueDictionary{}
    dict.Set("happy", "Friday")
    dict.Set("Say", "I am saying stuff")
    dict.Set("hey", "I am saying hey")
    has := dict.Has("hey")
    if !has {
        errors.New("expected key2 to be there")
    }
    items := dict.Keys()
    fmt.Printf("%v", items)
    values := dict.Values()
    fmt.Printf("%v", values)
}

func runStack() {
    s := stack.NewStack()
    s.Push(1)
    s.Push(2)
    s.Push(3)
    fmt.Println(s.Pop())
    fmt.Println(s.Pop())
    fmt.Println(s.Pop())
}

func runBst() {

	// Create some values to add to the Binary Search Tree

	tree := &bst.Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}

	// Print the sorted values added to the Binary Search Tree
	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *bst.Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	// Find a value in the Binary Search Tree
	s := "d"
	fmt.Print("Find node '", s, "': ")
	d, found := tree.Find(s)
	if !found {
		log.Fatal("Cannot find '" + s + "'")
	}
	fmt.Println("Found " + s + ": '" + d + "'")

	// Delete the value found in the Binary Search Tree
	err := tree.Delete(s)
	if err != nil {
		log.Fatal("Error deleting "+s+": ", err)
	}
	fmt.Print("After deleting '" + s + "': ")

	// Traverse the tree showing the tree after the node was deleted
	tree.Traverse(tree.Root, func(n *bst.Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	// Reset the tree to nil
	fmt.Println("Single-node tree")
	tree = &bst.Tree{}

	// Insert one node to show a single node tree
	tree.Insert("a", "alpha")
	fmt.Println("After insert:")

	// Traverse the tree showing a single node tree
	tree.Traverse(tree.Root, func(n *bst.Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	// Delete the single node from the tree
	tree.Delete("a")
	fmt.Println("After deletesss:")

	// Traverse the tree showing a nil tree
	tree.Traverse(tree.Root, func(n *bst.Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

}
