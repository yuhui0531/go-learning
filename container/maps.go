package container

import "fmt"

func main() {

	create()
	access()
	deleteK()

}

// map几乎所有的内奸类型都可以作为key,除了slice,map,function外.
// struct类型不包含slice,map,function等,也可以作为key

func deleteK() {
	m := map[string]string{
		"name":    "yuhui",
		"quality": "not bad",
	}
	name, ok := m["name"]
	fmt.Println("before deleting")
	fmt.Println(name, ok)
	delete(m, "name")
	fmt.Println("after deleting")
	fmt.Println(name, ok)
}

func access() {

	m := map[string]string{
		"name":    "yuhui",
		"quality": "not bad",
	}

	for k, v := range m {
		fmt.Printf("k=%s, v=%s\n", k, v)
	}

	name2, ok := m["abc"]
	fmt.Println("value="+name2, ok)

	if name, ok := m["abc"]; ok {
		fmt.Println(name)
	} else {
		fmt.Printf("the value associated with %d does not exist.", m["name"])
	}

	var m3 map[string]int
	k, ok := m3["aaa"]
	fmt.Println(k, ok)
}

func create() {
	m := map[string]string{
		"name":    "yuhui",
		"quality": "not bad",
	}
	m2 := make(map[string]int)
	// m2  == empty map
	var m3 map[string]int
	fmt.Println(m, m2, m3)
}
