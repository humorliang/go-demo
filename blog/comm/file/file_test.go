package file

import (
	"testing"
	"fmt"
	"time"
	"os"
)

func TestFile(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02"))
	fmt.Println(os.Getwd())
	fmt.Println(TodayLogFileName("info"))
}
