package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/EgMeln/client/internal/config"
	"github.com/EgMeln/client/internal/model"
	"github.com/EgMeln/client/internal/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Warnf("Config error %v", err)
	}

	mute := new(sync.RWMutex)
	priceMap := map[string]*model.GeneratedPrice{
		"Aeroflot": {},
		"ALROSA":   {},
		"Akron":    {},
	}
	ctx := context.Background()
	priceClient := server.ConnectPriceServer(cfg.PriceServicePort)
	log.Infof("start")
	go server.SubscribePrices(ctx, "Aeroflot", priceClient, mute, priceMap)
	go server.SubscribePrices(ctx, "ALROSA", priceClient, mute, priceMap)
	go server.SubscribePrices(ctx, "Akron", priceClient, mute, priceMap)
	posClient := server.NewPositionServer(priceMap, mute, cfg.PositionServicePort)

	log.Infof("Start open")
	t := time.Now()
	var array []string
	for i := 0; i < 100; i++ {
		id := posClient.OpenPositionAsk(ctx, "Aeroflot")
		array = append(array, id)
	}
	time.Sleep(1 * time.Minute)
	for _, id := range array {
		posClient.ClosePositionAsk(ctx, id, "Aeroflot")
	}
	log.Info(time.Since(t))
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-c
	log.Info("END")
}
