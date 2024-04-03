package main

import (
	"context"
	"log"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestIncRedis(t *testing.T) {

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Could not stop redis: %s", err)
	}

	defer func() {
		if err := redisC.Terminate(ctx); err != nil {
			log.Fatalf("Could not stop redis: %s", err)
		}
	}()

	host, err := redisC.Host(ctx)
	if err != nil {
		log.Fatalf("Could not stop redis: %s", err)
	}
	port, err := redisC.MappedPort(ctx, "6379")
	if err != nil {
		log.Fatalf("Could not stop redis: %s", err)
	}

	rdb = redis.NewClient(&redis.Options{
		Addr: host + ":" + port.Port(), // Redis server address host:port
		DB:   0,                        // Default DB
	})

	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "TestIncRedis", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IncRedis(rdb); (err != nil) != tt.wantErr {
				t.Errorf("IncRedis() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
