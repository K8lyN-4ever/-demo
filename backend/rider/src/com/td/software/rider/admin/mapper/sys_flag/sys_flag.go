package sys_flag_mapper

import (
	"rider/src/com/td/software/rider/common/resources"
	"sync"
)

type SysFlagMapper struct {
}

var (
	sysFlagMapper *SysFlagMapper
	sysFlagOnce   sync.Once
)

var (
	sysFlagKey = "sys_flag"

	dispatchField = "dispatch"
	generateField = "generate"
)

func NewSysFlagMapperInstance() *SysFlagMapper {
	sysFlagOnce.Do(
		func() {
			sysFlagMapper = &SysFlagMapper{}
		})
	return sysFlagMapper
}

func getFlag(name string) bool {
	res, err := resources.RedisDb.HGet(sysFlagKey, name).Result()
	if err != nil {
		return false
	}
	return res == "true"
}

func setFlag(name string, flag string) error {
	_, err := resources.RedisDb.HSet(sysFlagKey, name, flag).Result()
	if err != nil {
		return err
	}
	return nil
}

func (SysFlagMapper) GetDispatchFlag() bool {
	return getFlag(dispatchField)
}

func (SysFlagMapper) GetGenerateFlag() bool {
	return getFlag(generateField)
}

func (SysFlagMapper) SetDispatchFlag(flag string) error {
	return setFlag(dispatchField, flag)
}

func (SysFlagMapper) SetGenerateFlag(flag string) error {
	return setFlag(generateField, flag)
}
