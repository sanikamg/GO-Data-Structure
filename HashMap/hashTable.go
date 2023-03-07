package main

import "fmt"

//hashtable size
const size = 7

// hashtable will hold an array
type HashTable struct {
	array [size]*bucketList
}

//Bucketlist is a linked list each slot contain a linked list
type bucketList struct {
	head *bucketNode
}

// bucketNode is linked list node that holds the key
type bucketNode struct {
	key   string
	value int
	next  *bucketNode
}

func (b *bucketList) insertList(k string, v int) {
	newNode := new(bucketNode)
	newNode.key = k
	newNode.value = v

	if b.head == nil {
		b.head = newNode
	} else {
		temp := b.head
		for temp != nil {
			if temp.next == nil {
				temp.next = newNode
				return
			} else {
				temp = temp.next
			}
		}
	}
}

// find a index to add key to hash tabe and insert key to hash table
func (h *HashTable) insertHash(key string, value int) {
	index := hash(key)
	h.array[index].insertList(key, value)
}

//hash function to find index for the key
func hash(key string) int {

	asciiSum := 0

	for _, value := range key {
		asciiSum = asciiSum + int(value)
	}

	index := asciiSum % size
	return index

}

//init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucketList{}
	}
	return result
}

func (h *HashTable) get(key string) {
	index := hash(key)
	temp := h.array[index].head
	for temp != nil {
		if temp.key == key {
			fmt.Println(temp.value)
			return
		} else {
			temp = temp.next
		}
	}

}

func (h *HashTable) delete(key string) {
	index := hash(key)
	if h.array[index].head.key == key {
		h.array[index].head = h.array[index].head.next
		fmt.Println("Element Deleted")
	} else {
		temp := h.array[index].head
		for temp != nil {
			for temp.next != nil {
				if temp.next.key == key {
					temp.next = temp.next.next
					fmt.Println("Element deleted")
					return
				} else {
					temp = temp.next
				}
			}
		}
	}
}

func (h *HashTable) display() {

	for i := range h.array {
		temp := h.array[i].head
		fmt.Println(i, "---------")
		for temp != nil {
			fmt.Println("key", temp.key, "value", temp.value)
			temp = temp.next
		}
	}

}

func main() {
	hashtable := Init()

	hashtable.insertHash("sanika", 25)
	hashtable.insertHash("sanika", 26)
	hashtable.insertHash("aanika", 18)
	hashtable.insertHash("zanika", 26)
	hashtable.insertHash("lanika", 25)
	hashtable.insertHash("hanika", 26)
	hashtable.get("sanika")
	hashtable.display()
	hashtable.delete("hanika")
	fmt.Println("deleted hash table")
	hashtable.display()

}
