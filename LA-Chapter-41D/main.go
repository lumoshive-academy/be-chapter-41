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

	// implementasi struktur data hash
	StructureDataHash(ctx, rdb)

	// implementasi struktur data geo point
	StructurDataGeoPoit(ctx, rdb)

}

func StructureDataHash(ctx context.Context, rdb *redis.Client) {
	// Menambahkan field ke hash
	rdb.HSet(ctx, "user:1000", "name", "John Doe")
	rdb.HSet(ctx, "user:1000", "age", "30")
	rdb.HSet(ctx, "user:1000", "email", "johndoe@example.com")

	// Mengambil field dari hash
	name, err := rdb.HGet(ctx, "user:1000", "name").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name:", name)

	// Mengambil semua field dan nilai dari hash
	user, err := rdb.HGetAll(ctx, "user:1000").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User details:", user)

	// Memeriksa apakah field "email" ada dalam hash
	exists, err := rdb.HExists(ctx, "user:1000", "email").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email exists:", exists)

	// Menghapus field "age" dari hash
	rdb.HDel(ctx, "user:1000", "age")

	// Mengambil kembali semua field dan nilai setelah penghapusan
	updatedUser, err := rdb.HGetAll(ctx, "user:1000").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated user details:", updatedUser)
}

func StructurDataGeoPoit(ctx context.Context, rdb *redis.Client) {
	// Menambahkan titik-titik geografis ke Redis
	rdb.GeoAdd(ctx, "locations", &redis.GeoLocation{
		Name:      "Jakarta",
		Longitude: 106.8272,
		Latitude:  -6.1751,
	}, &redis.GeoLocation{
		Name:      "Bandung",
		Longitude: 107.6191,
		Latitude:  -6.9175,
	})

	// Mengambil posisi geografis dari suatu lokasi
	pos, err := rdb.GeoPos(ctx, "locations", "Jakarta").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Jakarta Position:", pos)

	// Menghitung jarak antara dua titik geografis
	dist, err := rdb.GeoDist(ctx, "locations", "Jakarta", "Bandung", "km").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Distance between Jakarta and Bandung: %.2f km\n", dist)

	// Mencari lokasi dalam radius 150 km dari Bandung
	locations, err := rdb.GeoRadius(ctx, "locations", 107.6191, -6.9175, &redis.GeoRadiusQuery{
		Radius:    150,
		Unit:      "km",
		WithCoord: true,
		WithDist:  true,
		Sort:      "ASC",
		Count:     10,
	}).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Locations within 150 km of Bandung:")
	for _, loc := range locations {
		fmt.Printf("%s: %.2f km (Lat: %.4f, Lon: %.4f)\n", loc.Name, loc.Dist, loc.Latitude, loc.Longitude)
	}
}

func StructureDataHyperLogLog(ctx context.Context, rdb *redis.Client) {
	// Menambahkan elemen ke HyperLogLog
	rdb.PFAdd(ctx, "myHyperLogLog", "user1", "user2", "user3", "user4", "user5")

	// Menambahkan lebih banyak elemen, termasuk elemen yang sudah ada
	rdb.PFAdd(ctx, "myHyperLogLog", "user2", "user3", "user6", "user7")

	// Menghitung elemen unik
	count, err := rdb.PFCount(ctx, "myHyperLogLog").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Estimated number of unique elements: %d\n", count)
}
