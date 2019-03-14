package cron

import (
	"log"
	"time"

	"urlooker/alarm/backend"
	"urlooker/alarm/cache"
	"urlooker/alarm/g"
	"urlooker/web/api"
	"urlooker/web/model"
)

func SyncStrategies() {
	t1 := time.NewTicker(time.Duration(g.Config.Web.Interval) * time.Second)
	for {
		syncStrategies()

		<-t1.C
	}

}

func syncStrategies() {

	var strategiesResponse api.StrategyResponse
	err := backend.CallRpc("Web.GetStrategies", "", &strategiesResponse)
	if err != nil {
		log.Println("[ERROR] Web.GetStrategies:", strategiesResponse.Data, strategiesResponse.Message, err)
		return
	}

	rebuildStrategyMap(strategiesResponse.Data)
}

func rebuildStrategyMap(strategiesResponse []*model.Strategy) {

	m := make(map[int64]model.Strategy)
	for _, strategy := range strategiesResponse {
		m[strategy.Id] = *strategy
	}

	cache.StrategyMap.ReInit(m)
}
