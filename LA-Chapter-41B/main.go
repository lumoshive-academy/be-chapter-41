package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

// Inisialisasi context untuk Go-Redis
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

	// Menampilkan pesan sukses koneksi
	fmt.Println("success connected to Redis")

	StructureDataString(ctx, rdb)

}

func StructureDataString(ctx context.Context, rdb *redis.Client) {
	// Set key "example" dengan nilai "Hello, Redis!"
	err := rdb.Set(ctx, "example", "Hello, Redis!", 0).Err()
	if err != nil {
		panic(err)
	}

	// Mendapatkan nilai dari key "example"
	val, err := rdb.Get(ctx, "example").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("example:", val)
}

func StructureDataList(ctx context.Context, rdb *redis.Client) {
	// Menambahkan elemen ke list
	rdb.RPush(ctx, "tasks", "task1")
	rdb.RPush(ctx, "tasks", "task2")
	rdb.LPush(ctx, "tasks", "task0")

	// Mengambil semua elemen dalam list
	tasks, err := rdb.LRange(ctx, "tasks", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("tasks:", tasks)

	// Menghapus dan mengambil elemen pertama dari list
	firstTask, err := rdb.LPop(ctx, "tasks").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("First task:", firstTask)

	// Mengambil elemen yang tersisa dalam list
	remainingTasks, err := rdb.LRange(ctx, "tasks", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Remaining tasks:", remainingTasks)
}
