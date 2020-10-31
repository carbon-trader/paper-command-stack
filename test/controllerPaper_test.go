package test

import (
	"testing"

	"github.com/carbon-trader/paper-core/repository"
)

var service = repository.PaperService{}

const (
	TESTPAPER = "TESTE1"
)

func init() {
	service.Server = "root:example@localhost"
	service.Database = "carbontrader_db"
	service.Connect()
}

func TestDatabaseConnection(t *testing.T) {
	// init database
	testing.Init()

	//
	if err := service.IsDBAlive(); err != nil {
		t.Errorf("Error %s", err)
	}
}
