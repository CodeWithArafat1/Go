package main

import "fmt"

// func main() {
// 	var a = 10 | 20
// 	fmt.Printf("type %T ", a)
// 	fmt.Println(a)

// }

// func main(){
// 	services := make(map[int]string)
// 	services[80] = "HTTP"
//     services[443] = "HTTPS"
//     services[22] = "SSH"
// 	fmt.Printf("Port 22 is running: %s\n", services[22])
// 	fmt.Println("All services:")
// }

// task 1

// func main() {
// 	credentials := make(map[string]string)

// 	credentials["admin"] = "admin123"
// 	credentials["arafat"] = "hacker99"
// 	credentials["guest"] = "123456"

// 	fmt.Printf("The hacker is: %s\n", credentials["arafat"])
// }

// func main() {
// 	targetHealth := 100

// 	healthPointer := &targetHealth
// 	fmt.Println("Original Health: ", targetHealth)
// 	fmt.Println("Memory Address:", healthPointer)
// 	*healthPointer = 0
// 	fmt.Println("Hacked Health: ", targetHealth)
// }

//! task 2

// func main() {
// 	ammo := 10
// 	ammoPtr := &ammo // Memory Address Access
// 	fmt.Println("Original Ammo: ", ammo)
// 	fmt.Println("Memory Address: ", ammoPtr)
// 	*ammoPtr = 999 // Inter the Memory and change value (hidden)
// 	fmt.Println("Hacked Ammo: ", ammo)
// }

// type Hacker struct {
// 	Codename string
// 	level    int
// 	IsActive bool
// }

// func main() {
// 	user1 :=
// 		Hacker{
// 			Codename: "Nill",
// 			level:    99,
// 			IsActive: true,
// 		}
// 	fmt.Printf("Hacker Name: %s\n", user1.Codename)
// 	fmt.Printf("Level: %d\n", user1.level)
// 	fmt.Printf("Level: %t\n", user1.IsActive)
// }

//! task 3

type Target struct {
	IP           string
	Port         int
	IsVulnerable bool
}

func main() {
	target1 := Target{
		IP:           "192.168.1.5",
		Port:         22,
		IsVulnerable: true,
	}

	fmt.Printf("IP: %s, Port: %d IsVulnerable: %t", target1.IP, target1.Port, target1.IsVulnerable)
}
