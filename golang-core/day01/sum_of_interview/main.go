package main

func main(){

}

func sum1(array []int64) int64 {
	var res int64 = 0
	for i:= 0 ; i < len(array) ; i++ {
		res += array[i]
	}
	return res
}