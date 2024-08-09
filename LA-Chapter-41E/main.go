package main

import (
	"context"
	"fmt"
	"log"

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

}

func PipelineRedis(ctx context.Context, rdb *redis.Client) {
	// Menggunakan pipeline untuk mengirim banyak perintah
	pipe := rdb.Pipeline()

	// Mengirim beberapa perintah secara bersamaan
	incr := pipe.Incr(ctx, "counter")
	incrBy := pipe.IncrBy(ctx, "counter", 10)
	set := pipe.Set(ctx, "key", "value", 0)

	// Eksekusi semua perintah dalam pipeline
	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Mendapatkan hasil dari setiap perintah
	fmt.Println("Counter after INCR:", incr.Val())
	fmt.Println("Counter after INCRBY:", incrBy.Val())
	fmt.Println("Set result:", set.Val())
}

func TransctionRedis(ctx context.Context, rdb *redis.Client) {
	// Menjalankan transaksi dengan TxPipelined
	err := rdb.Watch(ctx, func(tx *redis.Tx) error {
		// Dapatkan nilai saat ini dari kunci "balance"
		balance, err := tx.Get(ctx, "balance").Int()
		if err != nil && err != redis.Nil {
			return err
		}

		// Jalankan transaksi jika balance kurang dari 100
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, "balance", balance+50, 0)
			pipe.Set(ctx, "status", "updated", 0)
			return nil
		})

		return err
	}, "balance")

	if err != nil {
		log.Fatalf("Transaction failed: %v", err)
	} else {
		fmt.Println("Transaction completed successfully")
	}

	// Verifikasi hasil transaksi
	balance, err := rdb.Get(ctx, "balance").Int()
	if err != nil {
		log.Fatal(err)
	}
	status, err := rdb.Get(ctx, "status").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Final balance: %d\n", balance)
	fmt.Printf("Status: %s\n", status)
}
