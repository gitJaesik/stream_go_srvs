package config

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	fmt.Println(SglConfig)
}

func TestTTLConfig(t *testing.T) {
	fmt.Printf("30 Hours :%v\n", time.Hour*time.Duration(30))
	fmt.Printf("30 Hours :%d\n", time.Second*time.Duration(30*60*60))

	jwtTTL, err := strconv.ParseInt(SglConfig.GetJwtTtl(), 10, 64)
	if err != nil {
		jwtTTL = 1800
		fmt.Printf("error occurred: %v\n", err)
	}
	fmt.Printf("jwt : %v\n", time.Second*time.Duration(jwtTTL))
	fmt.Printf("jwt : %d\n", time.Second*time.Duration(jwtTTL))
}
