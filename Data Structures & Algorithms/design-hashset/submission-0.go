type MyHashSet struct {
	 hset [] int
}

func Constructor() MyHashSet {
    return MyHashSet{hset: []int{}}
}

func (this *MyHashSet) Add(key int) {
    if !this.Contains(key){
		this.hset=append(this.hset,key)
	}
}

func (this *MyHashSet) Remove(key int) {
    for i,v:= range this.hset{
		if v==key{
			this.hset=append(this.hset[:i],this.hset[i+1:]...)
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
    for _,v:=range this.hset{
		if v==key{
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 