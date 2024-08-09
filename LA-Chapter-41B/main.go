package main

import (
	"context"

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

	// Set key "example" dengan nilai "Hello, Redis!"
	err := rdb.Set(ctx, "example", "Hello, Redis!", 0).Err()
	if err != nil {
		panic(err)
	}
}

// // Mendapatkan nilai dari key "example"
// val, err := rdb.Get(ctx, "example").Result()
// if err != nil {
// 	panic(err)
// }
// fmt.Println("example:", val)

// // Coba mendapatkan nilai dari key yang tidak ada
// val2, err := rdb.Get(ctx, "nonexistent").Result()
// if err == redis.Nil {
// 	fmt.Println("Key 'nonexistent' tidak ditemukan")
// } else if err != nil {
// 	panic(err)
// } else {
// 	fmt.Println("nonexistent:", val2)
// }
// }
