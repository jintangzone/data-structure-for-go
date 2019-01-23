package advance

type TrieNode struct {
	isWord bool
	next map[rune]*TrieNode
}

type Trie struct {
	root *TrieNode
	size int
}

func (t *Trie) GetSize() int {
	return t.size
}

func (t *Trie) Add(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := rune(word[i])
		if _, ok := cur.next[c]; !ok {
			tn := &TrieNode{
				isWord: false,
				next: make(map[rune]*TrieNode),
			}
			cur.next[c] = tn
		}
		cur = cur.next[c]
	}
	if !cur.isWord {
		cur.isWord = true
		t.size++
	}
}

func (t *Trie) Contains(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := rune(word[i])
		if _, ok := cur.next[c]; !ok {
			return false
		}
		cur = cur.next[c]
	}
	return cur.isWord
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{isWord: false, next: make(map[rune]*TrieNode)},
		size: 0,
	}
}


