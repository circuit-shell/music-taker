package greeting

import "fmt"

func GetGreeting(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Go programming!", name)
}
