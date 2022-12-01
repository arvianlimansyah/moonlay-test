package server

import (
	"github.com/arvianlimansyah/moonlay-test/router"
	"github.com/arvianlimansyah/moonlay-test/utils"
)

func Run() {
	utils.InitLogger()
	e := router.New()
	e.Logger.Fatal(e.Start(":8000"))
}
