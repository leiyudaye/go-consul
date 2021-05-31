/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2019-04-15 12:46:38
 * @LastEditors: lly
 * @LastEditTime: 2021-05-31 18:28:26
 */
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	discover "napodate/discover"
	endpoint "napodate/endpoint"
	service "napodate/service"
)

func main() {
	var (
		httpAddr = flag.String("addr", ":8811", "http address")
	)
	flag.Parse()

	ctx := context.Background()
	errChan := make(chan error)
	srv := service.NewService()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// 服务注册
	kidDiscover, err := discover.NewKitDiscoverClient("127.0.0.1", 8500)
	if err != nil {
		os.Exit(1)
	}
	if !kidDiscover.Register("napodate", "1", "/healthCheck", "127.0.0.1", 8811, nil, &log.Logger{}) {
		os.Exit(1)
	}

	// mapping endpoints
	endpoints := endpoint.Endpoints{
		GetEndpoint:         endpoint.MakeGetEndpoint(srv),
		StatusEndpoint:      endpoint.MakeStatusEndpoint(srv),
		ValidateEndpoint:    endpoint.MakeValidateEndpoint(srv),
		HealthCheckEndpoint: endpoint.MakeHealthCheckEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("napodate is listening on port:", *httpAddr)
		handler := endpoint.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
