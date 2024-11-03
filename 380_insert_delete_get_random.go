package main

import "fmt"

type RandomizedSet struct {
	val_arr []int
	indices map[int]int
}

func Constructor() RandomizedSet {

	return RandomizedSet{
		val_arr: []int{},
		indices: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {

	if _, exists := this.indices[val]; exists {
		return false
	}
	this.indices[val] = len(this.val_arr)
	this.val_arr = append(this.val_arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {

}

func main() {
	obj := Constructor()
	param_1 := obj.Insert(5)

	fmt.Println(param_1)
}
