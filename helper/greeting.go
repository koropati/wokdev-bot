package helper

import (
	"math/rand"
	"time"
)

func GetGreeting() string {
	hour := time.Now().Hour()

	switch {
	case hour >= 5 && hour < 12:
		return "Selamat pagi!"
	case hour >= 12 && hour < 15:
		return "Selamat siang!"
	case hour >= 15 && hour < 18:
		return "Selamat sore!"
	default:
		return "Selamat malam!"
	}
}

func GetRandomGreetingOpening() string {
	greetings := []string{"Hey", "Halo", "Hello", "Hola"}

	// Mengatur seed acak berdasarkan waktu saat ini
	rand.Seed(time.Now().UnixNano())

	// Memilih ucapan secara acak dari slice greetings
	randomIndex := rand.Intn(len(greetings))
	return greetings[randomIndex]
}
