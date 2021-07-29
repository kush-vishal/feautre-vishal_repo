package arithmetic

	import "fmt"
   func Factorial(n int) {
    var f int = 1
    for i := 2; i <= n; i++ {
        f *= i
	fmt.Println(f)
    }
    
}