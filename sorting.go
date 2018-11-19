package main 


import("fmt"
		"sort"
)

type byLength []string

func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}


func main() {
	
	strs := []string{"hello","hi","bye"}

	sort.Strings(strs)
	fmt.Println(strs)

	// sort by length
	fruits := []string{"guyabano", "banana", "apple"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}