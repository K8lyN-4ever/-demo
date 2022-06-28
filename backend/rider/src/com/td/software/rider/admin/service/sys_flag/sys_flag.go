package sys_flag_service

import (
	"rider/src/com/td/software/rider/admin/mapper/sys_flag"
	"sync"
)

type SysFlagService struct {
}

var (
	sysFlagMapper  = sys_flag_mapper.NewSysFlagMapperInstance()
	sysFlagService *SysFlagService
	sysFlagOnce    sync.Once
)

func NewSysFlagServiceInstance() *SysFlagService {
	sysFlagOnce.Do(
		func() {
			sysFlagService = &SysFlagService{}
		})
	return sysFlagService
}

func (SysFlagService) GetDispatchFlag() bool {
	return sysFlagMapper.GetDispatchFlag()
}

func (SysFlagService) GetGenerateFlag() bool {
	return sysFlagMapper.GetGenerateFlag()
}

func (SysFlagService) SetDispatchFlag(flag string) bool {
	if err := sysFlagMapper.SetDispatchFlag(flag); err != nil {
		return false
	}
	return true
}

func (SysFlagService) SetGenerateFlag(flag string) bool {
	if err := sysFlagMapper.SetGenerateFlag(flag); err != nil {
		return false
	}
	return true
}
