package main 


import("fmt"
		c "github.com/rollyn/conversion"
)

func main() {
	
	i23 := int64(23)
	fmt.Println(  c.DecimalToBinary(i23) )
	
}