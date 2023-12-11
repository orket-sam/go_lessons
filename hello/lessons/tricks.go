package lessons

func DisplayUsers(users ...string) {
	for _, user := range users {
		println(user)
	}
}
