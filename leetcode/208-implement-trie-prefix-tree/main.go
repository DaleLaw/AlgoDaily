package main

import (
	"fmt"
)

// type but, in my opinion, it is a class
type Trie struct {
	isWord   bool
	children map[string]Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := Trie{false, make(map[string]Trie)}
	return root
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	current := this
	for i := 0; i < len(word); i++ {

		charactor := string(word[i])

		var temp Trie
		if _, existed := current.children[charactor]; existed {
			temp = current.children[charactor]
		} else {
			temp = Constructor()
		}

		if i == len(word)-1 {
			temp.isWord = true
		}
		current.children[charactor] = temp
		current = &temp
	}
}

// /** Returns if the word is in the trie. */
// func (this *Trie) Search(word string) bool {

// }

// /** Returns if there is any word in the trie that starts with the given prefix. */
// func (this *Trie) StartsWith(prefix string) bool {

// }

func main() {
	var hashtable = make(map[string]string)
	// fmt.Println(hashtable)
	hashtable["a"] = "calvi"
	hashtable["b"] = "calvin"
	// fmt.Println(hashtable["c"])
	// fmt.Println(hashtable)

	// a_node := make(map[string]Trie)
	// a_node["b"] = nil
	// b := Trie{false, a_node}
	a := Trie{false, make(map[string]Trie)}
	b := Trie{false, make(map[string]Trie)}
	a.children["a"] = b
	// fmt.Println(a)
	if _, ok := a.children["a"]; ok {
		// fmt.Println(a.children["a"])
	}

	t := Constructor()
	t.Insert("abc")
	t.Insert("app")
	fmt.Println(t)
}
