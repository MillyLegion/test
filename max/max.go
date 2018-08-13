package max

//import "fmt"

func Max(a, b, c int ) int {
   // fmt.Printf("Arrayvalues: %v, %v, %v\n", a, b, c)
   if a > b && a > c {
     return a
   } else if b>a && b > c {
     return b
   }
   return c
}

func main() {
   print(Max(10,20,30))
}
