package main
import "fmt"

func main(){
	var a, b, hasil int
	fmt.Scan(&a, &b)
	hasil = pangkat(a, b)
	fmt.Print(hasil)
}
func pangkat(a, b int)int{
	if b == 1{
		return a
	}else{
		return a*pangkat(a, b-1)
	}
}