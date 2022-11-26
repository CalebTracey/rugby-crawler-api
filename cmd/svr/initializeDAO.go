package main

import (
	config "github.com/calebtracey/config-yaml"
	compdao "github.com/calebtracey/rugby-crawler-api/internal/dao/comp"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/psql"
	"github.com/calebtracey/rugby-crawler-api/internal/facade"
	"github.com/calebtracey/rugby-crawler-api/internal/facade/comp"
	log "github.com/sirupsen/logrus"
)

func initializeDAO(config config.Config) (facade.APIFacadeI, error) {
	crawlerConfig, err := config.GetCrawlConfig("COLLY")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	psqlDbConfig, err := config.GetDatabaseConfig("PSQL")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return facade.APIFacade{
		CompService: comp.Facade{
			DbDAO: psql.DAO{
				Db: psqlDbConfig.DB,
			},
			CompDAO: compdao.DAO{
				Collector: crawlerConfig.Collector,
			},
			DbMapper: psql.Mapper{},
		},
	}, nil
}
