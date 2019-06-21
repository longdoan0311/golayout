package user

import (
	"{{.ModName}}"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Cfg configuration struct of user package
type Cfg struct {
	Store {{.ProjName}}.Store
}

// InfoImpl concrete implementation of {{.ProjName}}.UserInfoService
type infoImpl struct {
	store {{.ProjName}}.Store
}

// NewInfoImpl initialize user package. must be invoke this func before use this package
func NewInfoImpl(cfg Cfg) {{.ProjName}}.UserInfoService {
	log.Info("New UserInfoService")
	var impl infoImpl
	impl.store = cfg.Store

	return &impl
}

func (s *infoImpl) GetAll() ([]{{.ProjName}}.User, error) {
	log.Info("Get all users")

	r, err := s.store.User().GetAll()
	if err != nil {
		return r, errors.Wrap(err, "fail to GetAll() users")
	}

	return r, nil
}