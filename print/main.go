package main
import ("fmt")
func print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	print([]string{"Hello, ", "overstarry\n"})
	print([]int64{1, 2, 3})
	print([]float64{1, 2, 3})
}