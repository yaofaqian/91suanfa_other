package topic_380

type RandomizedSet struct {
	strackLen int
	strack    []int
	strackMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{0, []int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	this.strack = append(this.strack, val)
	this.strackMap[val] = this.strackLen
	this.strackLen++
}

func (this *RandomizedSet) Remove(val int) bool {
	index := this.strackMap[val]
	this.strack = append(this.strack[:index], this.strack[index+1:]...)
	delete(this.strackMap, val)
	this.strackLen--
}
func (this *RandomizedSet) GetRandom() int {

}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
