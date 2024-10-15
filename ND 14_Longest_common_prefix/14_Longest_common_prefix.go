package main

import "strings"

type Node struct {
	Char      string
	childNode [26]*Node
}

func NewNode(char string) *Node {
	node := &Node{Char: char}
	for i := 0; i < 26; i++ {
		node.childNode[i] = nil
	}
	return node
}

type Trie struct {
	RootNode *Node
}

func NewTrie() *Trie {
	root := NewNode("_")
	return &Trie{RootNode: root}

}

// / Insert inserts a word to the trie
func (t *Trie) Insert(word string) error {
	current := t.RootNode                    // Начинаем с корня
	strippedWord := strings.ToLower(word)    // Приводим к нижнему регистру
	for i := 0; i < len(strippedWord); i++ { // Посимвольно проходимся по слову
		index := strippedWord[i] - 'a'       // получаем позицию буквы от 0 до 25
		if current.childNode[index] == nil { // проверяем есть ли у нас уже дочерняя нода
			current.childNode[index] = NewNode(string(strippedWord[i])) // Если нет, то добавляем
		}
		current = current.childNode[index] // переводим указатель на дочернюю ноду
	}
	return nil
}

func longestCommonPrefix(strs []string) string {

}

func main() {

}
