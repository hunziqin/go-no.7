package main

import "fmt"

func main() {
	var arrAge = [5]int{18, 20, 15, 22, 16}
	//var arrLazy = [...]int{5,6,7,8,22}
	//var arrLazy = []int{5,6,7,8,22}
	//var arrKeyValue = [5]string{3: "hun",4: "zi"}
	//var arrKeyValue = []string{3: "hun",4: "zi"}

	for i := 0; i < len(arrAge); i++ {
		fmt.Printf("Person at %d is %d\n", i, arrAge[i])
	}
	/*for i:=0;i<len(arrLazy);i++{
		fmt.Printf("Person at %d is %d\n",i,arrLazy[i])
	}
	for i:=0;i<len(arrKeyValue);i++{
		fmt.Printf("Person at %d is %s\n",i,arrKeyValue[i])
	}*/
}
