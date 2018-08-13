package max

func Max(i, ... int ) int {
   
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
