package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Membuat client Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // alamat Redis server
		Password: "",               // password Redis, kosongkan jika tidak ada
		DB:       0,                // DB Redis yang akan digunakan
	})

	// Mengecek koneksi ke Redis
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Redis")

	// Nama stream dan grup konsumen
	streamName := "mystream"
	groupName := "mygroup"
	consumerName := "consumer1"

	PublishStream(ctx, rdb, streamName)
	CreateConsumer(ctx, rdb, streamName, groupName, consumerName)
	GetStream(ctx, rdb, streamName, groupName, consumerName)
}

func PublishStream(ctx context.Context, rdb *redis.Client, streamName string) {
	// Membuat stream dengan data awal jika belum ada
	rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: streamName,
		Values: map[string]interface{}{
			"sensor-id":   "1234",
			"temperature": "22.5",
		},
	})

}

func CreateConsumer(ctx context.Context, rdb *redis.Client, streamName, groupName, consumerName string) {
	// Membuat grup konsumen pada stream
	err := rdb.XGroupCreate(ctx, streamName, groupName, "0").Err()
	if err != nil && err != redis.Nil {
		log.Fatal(err)
	}
	fmt.Println("Consumer group created")

	// Membuat konsumen dalam grup
	err = rdb.XGroupCreateConsumer(ctx, streamName, groupName, consumerName).Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Consumer created in group")
}

func GetStream(ctx context.Context, rdb *redis.Client, streamName, groupName, consumerName string) {

	// Konsumen membaca pesan dari grup
	entries, err := rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    groupName,
		Consumer: consumerName,
		Streams:  []string{streamName, ">"}, // Membaca pesan yang belum diklaim
		Count:    1,                         // Baca satu entri
		Block:    0,                         // Blok sampai data tersedia
	}).Result()

	if err != nil {
		log.Fatal(err)
	}

	for _, stream := range entries {
		for _, message := range stream.Messages {
			fmt.Printf("Message ID: %s\n", message.ID)
			fmt.Printf("Sensor ID: %s\n", message.Values["sensor-id"])
			fmt.Printf("Temperature: %s\n", message.Values["temperature"])
		}
	}
}

func PubSubRedis(ctx context.Context, rdb *redis.Client) {
	// Membuat subscriber
	pubsub := rdb.Subscribe(ctx, "mychannel")
	defer pubsub.Close()

	// Memberikan waktu untuk subscriber terhubung
	time.Sleep(1 * time.Second)

	// Mem-publish pesan ke saluran
	err := rdb.Publish(ctx, "mychannel", "Hello, World!").Err()
	if err != nil {
		log.Fatal(err)
	}

	// Menggunakan ReceiveMessage untuk menerima pesan secara eksplisit
	msg, err := pubsub.ReceiveMessage(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Received message: %s from channel: %s\n", msg.Payload, msg.Channel)
}
