package main

import "fmt"

func main() {
	/*	var (
		name	string
		age		int
	)*/
	/*
		n,_ := fmt.Sscan("polaris 28",&name,&age)
		// 可以将"polaris 28"中的空格换成"\n"试试
		// n, _ := fmt.Sscan("polaris\n28", &name, &age)
		fmt.Println(n,name,age) //2 polaris 28
	*/

	/*	n,_ := fmt.Sscanf("polaris 28","%s%d",&name,&age)
		// 可以将"polaris 28"中的空格换成"\n"试试
		// n, _ := fmt.Sscanf("polaris\n28", "%s%d", &name, &age
		fmt.Println(n,name,age) //2 polaris 28*/

	/*	n, _ := fmt.Sscanln("polaris 28", &name, &age)
		// 可以将"polaris 28"中的空格换成"\n"试试
		// n, _ := fmt.Sscanln("polaris\n28", &name, &age)
		fmt.Println(n, name, age)
	*/
	for i := 0; i < 2; i++ {
		var name string
		fmt.Print("Input Name:")
		n, err := fmt.Scanf("%s", &name)
		fmt.Println(n, err, name)
	}
}
