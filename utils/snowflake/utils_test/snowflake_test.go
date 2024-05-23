package utils_test

import (
	"fmt"
	"testing"
	"time"
	"web/utils/snowflake"
)

func TestXxx(t *testing.T) {
	snowflake.SetUpSnowFlakeWorker(0,0)

	go func ()  {
		time.Sleep(time.Second)
		panic("end")
	}()

	go func ()  {
		for {
			fmt.Println(snowflake.GlobalSnowFlakeWorker.NextID())
		}
	}()
	
	go func ()  {
		for {
			fmt.Println(snowflake.GlobalSnowFlakeWorker.NextID())
		}
	}()
}