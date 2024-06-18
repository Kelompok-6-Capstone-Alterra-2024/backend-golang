package utilities

import "fmt"

func AddContentComplaintUserNotification(name, message string) string {
	content := fmt.Sprintf("User %s has added a new complaint with message: %s", name, message)
	return content
}
