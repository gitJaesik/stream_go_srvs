package logger

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	fmt.Println(Logger)
	fmt.Printf("Logger :%v\n", Logger)
	Logger.Infow("TestInit", "logger: ", Logger)
}
