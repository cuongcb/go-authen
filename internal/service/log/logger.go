package log

import "fmt"

// Log prints logs to default stdout
func Log(msg ...interface{}) {
	fmt.Println(msg...)
}
