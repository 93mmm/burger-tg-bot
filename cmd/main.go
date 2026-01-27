package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"go.uber.org/zap"

	"github.com/93mmm/burger-tg-bot.git/cmd/config"
	"github.com/93mmm/burger-tg-bot.git/internal/app"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("can't parse app config: %v", err)
	}

	if cfg.MaxCpu > 0 {
		runtime.GOMAXPROCS(cfg.MaxCpu)
	}

	globalLogger := logger.NewLogger(zap.DebugLevel)
	logger.SetLogger(globalLogger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go signalHandler(ctx, cancel)

	application := app.NewApp(cfg)
	if err = application.Run(ctx); err != nil {
		logger.ErrorKV(ctx, "can't run application", err)
		return
	}
	logger.InfoKV(ctx, "application is shutdown normally")
}

// обработчик сигналов системы
func signalHandler(ctx context.Context, cancelFunc context.CancelFunc) {
	osSigCh := make(chan os.Signal, 1)

	signal.Notify(
		osSigCh,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)

	defer signal.Stop(osSigCh)
	s := <-osSigCh
	logger.InfoKV(ctx, "получен signal",
		"signal", s,
	)
	cancelFunc()
}
