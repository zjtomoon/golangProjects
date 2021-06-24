package main
/*
	defer的副作用：
	1、对返回值的影响
	2、对性能的影响
 */
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f4() int {
	r := 0
	defer func() {
		r++
	}()
	return r
}

func f5() int {
	r := 0
	defer func(i int) {
		i++
	}(r)
	return 0
}

func main() {
	println("f1=",f1()) //1
	println("f2=",f2()) //5
	println("f3=",f3()) //1
	println("f3=",f4()) //0
	println("f4=",f5()) //0
}
