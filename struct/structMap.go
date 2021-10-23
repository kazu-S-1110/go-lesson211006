package main

import "fmt"

// structをmap（object）に格納する

type User struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User{
		1: {"user1", 10},
		2: {"user2", 20},
		3: {"user3", 30},
		4: {"user4", 40},
	}

	fmt.Println(m) //map[1:{user1 10} 2:{user2 20} 3:{user3 30} 4:{user4 40}]

	// keyにUser型を指定してもおk
	m2 := map[User]string{
		{"user1", 10}: "golang",
		{"user2", 20}: "React",
		{"user3", 30}: "Python",
	}
	fmt.Println(m2) //map[{user1 10}:golang {user2 20}:React {user3 30}:Python]

	// makeを使うことも可能
	m3 := make(map[int]User)
	fmt.Println(m3) //map[]
	m3[1] = User{"hoge", 39}
	fmt.Println(m3) //map[1:{hoge 39}]

}
