package main

func main(){
	 nums := []int{1, 2, 3}
	 tens := []int{12, 13}
	 nums = append(nums, 4, 9)
	 nums = append(nums, tens...)
}