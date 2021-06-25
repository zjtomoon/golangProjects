package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := []int{0,1,2,3,4,5,6}
	b := a[0:4] //[0,1,2,3]
	//fmt.Println(a,b)
	as := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	//a、b共享底层数组
	fmt.Printf("a = %v,len = %d,cap = %d ,type = %d\n",a,len(a),cap(a),as.Data)
	fmt.Printf("b = %v,len = %d,cap = %d ,type = %d\n",b,len(b),cap(b),bs.Data)

	b = append(b,10,11,12)

	//a、b继续共享底层数组，修改b会影响共享的底层数组，间接影响a
	fmt.Printf("a = %v,len = %d,cap = %d\n",a,len(a),cap(a))
	fmt.Printf("b = %v,len = %d,cap = %d\n",b,len(b),cap(b))

	//len(b)=7,底层数组容量是7，此时需要重新分配数组，并将原来数组值复制到新数组
	b = append(b,13,14)

	as = (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs = (*reflect.SliceHeader)(unsafe.Pointer(&b))


	//可以看到a和b指向底层数组的指针已经不同了
	fmt.Printf("a = %v,len = %d,cap = %d,type = %d\n",a,len(a),cap(a),as.Data)
	fmt.Printf("b = %v,len = %d,cap = %d,type = %d\n",b,len(b),cap(b),bs.Data)

}

/*

	多个切片共享一个底层数组，其中一个切片的append操作可能引发如下两种情况:
	(1) append追加的元素没有超过底层数组的容量，此种append操作会直接操作共享的底层数组，如果其他切片有引用数组被覆盖的元素，则会导致其他
	切片的值也隐式地发生变化
	(2) append追加地元素加上原来地元素如果超过底层数组地容量，则此种append操作会重新申请数组，并将原来数组复制到新数组
 */