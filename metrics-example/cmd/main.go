package main

import (
	"context"
	"github.com/AbnormalReality/gb_gobackend_2/lesson4/metrics-example/infra/server"
	"github.com/AbnormalReality/gb_gobackend_2/lesson4/metrics-example/logic"
	"net/http"
	"os"
	"os/signal"

	"go.uber.org/zap"

	"github.com/AbnormalReality/gb_gobackend_2/lesson4/metrics-example/infra/telemetry"
)

func main() {
	l := zap.L().Sugar()

	// Слушаем сигналы ОС для завершения работы
	ctx, cancel := newOSSignalContext()
	defer cancel()

	// Настраиваем сборщик трейсов
	tp, err := telemetry.RunTracingCollection(ctx)
	if err != nil {
		l.Panic(err)
	}
	defer func() {
		if err = tp.Shutdown(context.Background()); err != nil {
			l.Errorf("failed to stop the traces collector: %v", err)
		}
	}()

	tr := tp.Tracer("server")
	// Запускаем сервер
	s := &server.S{
		Tr: tr,
		Logic: &logic.Logic{
			Tr: tr,
		},
	}
	go func() {
		err := s.Start()
		if err != nil && err != http.ErrServerClosed {
			l.Panic(err)
		}
	}()

	<-ctx.Done()

	err = s.Stop(context.Background())
	if err != nil {
		l.Error(err)
	}
}

func newOSSignalContext() (context.Context, func()) {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	return ctx, func() {
		signal.Stop(c)
		cancel()
	}
}
