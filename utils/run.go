package utils

import "web/utils/snowflake"

func InitUtils() {
	snowflake.SetUpSnowFlakeWorker(0, 0)
}
