package lessons

import "fmt"

var Users = []string{"orket", "shalom", "Nelly", "Ian", "Sam"}

func DisplayUsers(users []string) {
	// for i := 0; i < len(users); i++ {
	// 	fmt.Println(users[i])
	// }
	for _, user := range users {
		fmt.Println(user)
	}

	println("done")
}
