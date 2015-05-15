package collection

type StringSet struct {
	data map[string]bool
}

func NewStringSet() *StringSet {
	return &StringSet{make(map[string]bool)}
}

func (set *StringSet) Add(s string) bool {
	_, found := set.data[s]
	set.data[s] = true
	return !found
}

func (set *StringSet) Contains(s string) bool {
	_, found := set.data[s]
	return found
}

func (set *StringSet) Remove(s string) {
	delete(set.data, s)
}

func (set *StringSet) Size() int {
	return len(set.data)
}
