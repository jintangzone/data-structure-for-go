package main

import (
	"data_structure/Array"
	"fmt"
)

func main() {

	arr := Array.NewArray(20)
	for i := 0; i < 10; i++ {
		arr.Append(i)
	}
	fmt.Println(arr.ToString())

	arr.Insert(1, 100)
	fmt.Println(arr.ToString())

	arr.Prepend(99)
	fmt.Println(arr.ToString())

	arr.Append(200)
	fmt.Println(arr.ToString())

	arr.Set(4, 101)
	fmt.Println(arr.ToString())

	fmt.Println("Get Index = 8:", arr.Get(8))

	fmt.Println("Array contains 101 ?", arr.Contains(101))
	fmt.Println("Array contains 201 ?", arr.Contains(201))
	fmt.Println("index of 9 =", arr.Index(9))

	fmt.Println("Delete index 5 =", arr.Delete(5))
	fmt.Println(arr.ToString())

	fmt.Println("Pop: ", arr.Pop())
	fmt.Println(arr.ToString())

	fmt.Println("UnShift: ", arr.UnShift())
	fmt.Println(arr.ToString())

	fmt.Println("Delete Element 1:", arr.DeleteElement(1))
	fmt.Println(arr.ToString())

	fmt.Println()
	fmt.Println("New String Array:")
	arrStr := Array.NewArray(10)

	arrStr.Append("php")
	arrStr.Append("java")
	arrStr.Append("golang")
	arrStr.Append("objectc")
	arrStr.Append("python")
	arrStr.Append("ruby")
	arrStr.Append("cpp")
	arrStr.Append("clang")
	arrStr.Append("erlang")
	fmt.Println(arrStr.ToString())

	arrStr.Set(8, "ios")
	fmt.Println(arrStr.ToString())

	fmt.Println("Get Index = 8:", arrStr.Get(8))

	fmt.Println("Array contains php ?", arrStr.Contains("php"))
	fmt.Println("Array contains javascript ?", arrStr.Contains("javascript"))
	fmt.Println("index of python =", arrStr.Index("python"))

	fmt.Println("Delete index 5 =", arrStr.Delete(5))
	fmt.Println(arrStr.ToString())

	fmt.Println("Pop: ", arrStr.Pop())
	fmt.Println(arrStr.ToString())

	fmt.Println("UnShift: ", arrStr.UnShift())
	fmt.Println(arrStr.ToString())

	fmt.Println("Delete Element objectc:", arrStr.DeleteElement("objectc"))
	fmt.Println(arrStr.ToString())
}