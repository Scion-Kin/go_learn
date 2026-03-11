package main

type Membership struct {
	Type string
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType string) User {
	charLimit := 1000
	if (membershipType != "premium") { charLimit = 100 }
	return User{
		Name: name,
		Membership: Membership{
			Type: membershipType,
			MessageCharLimit: charLimit,
		},
	}
}
