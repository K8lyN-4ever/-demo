package application

import (
	"rider/src/com/td/software/rider/common/resources"
	"rider/src/com/td/software/rider/common/router"
	"rider/src/com/td/software/rider/common/util"
)

func Init() error {

	if err := resources.Db(); err != nil {
		return err
	}
	if err := resources.InitClient(); err != nil {
		return err
	}
	if err := router.Router(); err != nil {
		return err
	}
	if err := util.InitLogger(); err != nil {
		return err
	}
	return nil
}
