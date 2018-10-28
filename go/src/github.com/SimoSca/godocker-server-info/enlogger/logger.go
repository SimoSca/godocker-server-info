package logger

import (
	"fmt"
	"time"
)

// Print Display log
func Print(text string) {
	var t = time.Now().Format(time.RFC3339)
	fmt.Printf("[%s] %s \n",t, text)
}