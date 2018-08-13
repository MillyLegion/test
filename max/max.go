package max

func main(){
  println(Max(1,2,3))
  println(Max(2,3,1))
  println(Max(3,2,1))
}

func Max(nums ...int) int {
   max := nums[0]
   for _, num := range nums{
      if num > max{
         max = num
      }
   }
   return max
}

