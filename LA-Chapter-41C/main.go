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

	// implementasi struktur data set
	StructureDataSet(ctx, rdb)

	// implementasi struktur data sortedset
	StructurDataSortedSet(ctx, rdb)

}

func StructureDataSet(ctx context.Context, rdb *redis.Client) {
	// Menambahkan elemen ke set
	rdb.SAdd(ctx, "myset", "apple", "banana", "cherry")

	// Menambahkan elemen yang sama, tidak akan ditambahkan lagi karena elemen harus unik
	rdb.SAdd(ctx, "myset", "banana")

	// Mengambil semua elemen dalam set
	members, err := rdb.SMembers(ctx, "myset").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("myset members:", members)

	// Mengecek apakah elemen ada dalam set
	isMember, err := rdb.SIsMember(ctx, "myset", "apple").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Is 'apple' a member of myset?", isMember)
}

func StructurDataSortedSet(ctx context.Context, rdb *redis.Client) {
	// Menambahkan elemen ke sorted set
	rdb.ZAdd(ctx, "leaderboard", &redis.Z{Score: 100, Member: "Alice"})
	rdb.ZAdd(ctx, "leaderboard", &redis.Z{Score: 200, Member: "Bob"})
	rdb.ZAdd(ctx, "leaderboard", &redis.Z{Score: 150, Member: "Charlie"})

	// Mengambil elemen-elemen dari sorted set
	leaderboard, err := rdb.ZRangeWithScores(ctx, "leaderboard", 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Leaderboard:")
	for _, z := range leaderboard {
		fmt.Printf("%s: %.0f\n", z.Member, z.Score)
	}

	// Mendapatkan peringkat dari elemen tertentu
	rank, err := rdb.ZRank(ctx, "leaderboard", "Charlie").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Charlie's rank: %d\n", rank+1)
}
