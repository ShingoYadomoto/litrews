package job

import "fmt"

type ReminderEmails struct {
}

func (e ReminderEmails) Run() {
	// Queries the DB
	// Sends some email
	fmt.Printf("Every 5 sec send reminder emails \n")
}
