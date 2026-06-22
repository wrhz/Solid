package solid

import (
	"sync"

	"xorm.io/xorm"
)

var xormEngine *xorm.Engine
var XormSessionsManager *XormSessionsManagerStruct = &XormSessionsManagerStruct{}

type XormSessionsManagerStruct struct {
	store sync.Map
}

func (g *XormSessionsManagerStruct) Set(requestID string) error {
	session := xormEngine.NewSession()

	err := session.Begin()

	if err != nil {
		return err
	}

	g.store.Store(requestID, session)

	return nil
}

func (g *XormSessionsManagerStruct) Get(requestID string) (*xorm.Session, bool) {
	orm, ok := g.store.Load(requestID)

	if ok {
		return orm.(*xorm.Session), true
	}

	return nil, false
}

func (g *XormSessionsManagerStruct) Delete(requestID string) {
	orm, ok := g.store.Load(requestID)

	if ok {
		orm.(*xorm.Session).Close()

		g.store.Delete(requestID)
	}
}

func InitXorm() error {
	databaseConfig := GetDatabaseConfig()

	xormDriverName := databaseConfig.GetXormDriverName()
	xormDataSourceName := databaseConfig.GetXormDataSourceName()

	if xormDriverName != "" && xormDataSourceName != "" {
		var err error

		xormDriverOptions := databaseConfig.GetXormDriverOptions()
		xormEngine, err = xorm.NewEngine(xormDriverName, xormDataSourceName, xormDriverOptions...)

		if err != nil {
			return err
		}

		xormEngine.ShowSQL(databaseConfig.GetXormShowSQL())
	}

	return nil
}

func RemoveXorm() error {
	if IsStartXorm() {
		return xormEngine.Close()
	}

	return nil
}

func IsStartXorm() bool {
	return xormEngine != nil
}