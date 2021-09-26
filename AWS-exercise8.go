package main

import "fmt"

func main() {
	ar := []int{1, 2, 3}
	arr := []int{4, 5, 6}
	fmt.Println(InsertStringSlice(2, ar, arr))
}

func InsertStringSlice(index int, ar []int, arr []int) (ns []int) {
	ns = append(ns, ar[:index]...)
	ns = append(ns, arr[:index]...)
	ns = append(ns, ar[index:]...)
	return
}

// InsertStringSlice 将切片插入到另一个切片的指定位置
/*package main

import (
	"fmt"
)

func main() {
	src := []string{"a", "b", "c", "d", "e"}
	newstr := []string{"f", "g"}
	fmt.Println(insertSlice(2, newstr, src))
}

func insertSlice(index int, newstr []string, src []string) (ns []string) {
	ns = append(ns, src[:index]...) // 切片后加..., 相当于拆包成单个元素
	ns = append(ns, newstr...)
	ns = append(ns, src[index:]...)
	return
}*/

// remove sub slice
/*package main

import (
    "fmt"
    "time"
)

func main() {
    s := []int{0, 1,2,3,4,5,6,7,8,9}
    fmt.Println(removeSub(s, 30, 70))
}

func removeSub(s []int, start int, end int) (ns []int) {
    for ix, num := range s {
        if start <= ix && ix < end {
            continue
        }
        ns = append(ns, num)
    }
    return
    // ns = append(ns, s[:start]...)
    // ns = append(ns, s[end:]...)
    // return
}*/
