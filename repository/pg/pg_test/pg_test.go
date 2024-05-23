package pgtest_test

import (
	"web/config"
	repository "web/repository/pg"
	"testing"
)

func init(){
	config.InitConfig("")
}

func TestInitDB(t *testing.T) {
	repository.InitPg(config.Configure)
}
