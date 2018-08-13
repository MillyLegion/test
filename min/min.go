package min

func main(){
  println(Min(1,2,3))
  println(Min(2,3,1))
  println(Min(3,2,1, 0,-1,3))
}

func Min(nums ...int) int {
   min := nums[0]
   for _, num := range nums{
      if num < min{
         min = num
      }
   }
   return min
}
