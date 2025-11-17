package main

import (
	"net/http"
	"user-management/config"
)

func main() {
	config.NewEnv()
	logger := config.NewLogger()
	validator := config.NewValidator()
	db := config.NewDatabase()
	cache := config.NewCache()
	r := config.NewRouter()
	conn, ch := config.NewAmqp()
	defer conn.Close()
	defer ch.Close()
	lis, s := config.NewGrpc()
	config.Initialize(&config.Bootstrap{
		DB:        db,
		Cache:     cache,
		Logger:    logger,
		Router:    r,
		Validator: validator,
		Ch:        ch,
		S:         s,
	})
	go func() {
		err := s.Serve(lis)
		if err != nil {
			logger.Error().Err(err).Msg("failed start server for grpc")
		}
	}()
	err := http.ListenAndServe("0.0.0.0:3001", r)
	if err != nil {
		panic(err)
	}
}
