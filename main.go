package main

import (
	"fmt"
	"time"
)

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

//! task 1

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

// type Target struct {
// 	IP           string
// 	Port         int
// 	IsVulnerable bool
// }

// func main() {
// 	target1 := Target{
// 		IP:           "192.168.1.5",
// 		Port:         22,
// 		IsVulnerable: true,
// 	}

// 	fmt.Printf("IP: %s, Port: %d IsVulnerable: %t", target1.IP, target1.Port, target1.IsVulnerable)
// }

// type Target struct {
// 	IP   string
// 	Port int
// }

// func (t Target) Ping() {
// 	fmt.Printf("Sending packets to %s on port %d....\n", t.IP, t.Port)
// }

// func main() {
// 	server := Target{IP: "10.0.0.1", Port: 80}
// 	server.Ping()
// }

//! task 4

// type Hacker struct {
// 	Codename string
// 	level    int
// 	IsActive bool
// }

// type Target struct {
// 	IP           string
// 	Port         int
// 	IsVulnerable bool
// }

// func (t Target) Exploit() {
// 	if t.IsVulnerable == true {
// 		fmt.Printf("Success! Hacked into IP %s", t.IP)
// 	} else {
// 		fmt.Printf("Failed! IP: %s is secure.", t.IP)
// 	}
// }

// func main() {
// 	target1 := Target{
// 		IP:           "192.168.1.5",
// 		Port:         22,
// 		IsVulnerable: false,
// 	}

// 	target1.Exploit()
// }

// func main() {
// 	var name string = "Arafat"
// 	var age int = 22
// 	var height float64 = 5.7
// 	var isStudent bool = true

// 	// short declaration
// 	country := "Bangladesh"
// 	salary := 45000

// 	fmt.Println("Name:", name)
// 	fmt.Println("Age:", age)
// 	fmt.Println("Height:", height)
// 	fmt.Println("Country:", country)
// 	fmt.Println("Salary:", salary)
// 	fmt.Println("Student:", isStudent)
// }

// type Scanner interface {
// 	Scan()
// }

// type WebScanner struct {
// 	URL string
// }

// type PortScanner struct {
// 	IP string
// }

// func (w WebScanner) Scan() {
// 	fmt.Printf("Scanning website vulnerabilities for: %s\n", w.URL)
// }

// func (p PortScanner) Scan() {
// 	fmt.Printf("Scanning open ports for IP: %s\n", p.IP)
// }

// func startAttack(s Scanner) {
// 	s.Scan()
// }

// func main() {
// 	target1 := WebScanner{URL: "http://vuln-site.com"}
// 	target2 := PortScanner{IP: "192.168.1.100"}

// 	startAttack(target1)
// 	startAttack(target2)
// }

//! task 5

// type Payload interface {
// 	Execute()
// }

// type SQLInjection struct {
// 	TargetDB string
// }

// type Ransomware struct {
// 	TargetFolder string
// }

// func (s SQLInjection) Execute() {
// 	fmt.Printf("[!] Extracting admin passwords from database: %s\n", s.TargetDB)
// }

// func (t Ransomware) Execute() {
// 	fmt.Printf("Encrypting all files in folder: %s\n", t.TargetFolder)
// }

// func deployPayload(P Payload) {
// 	P.Execute()
// }

// func main() {
// 	tDB := SQLInjection{TargetDB: "myShop"}
// 	targetF := Ransomware{TargetFolder: "user"}

// 	deployPayload(tDB)
// 	deployPayload(targetF)
// }

//? error handel
// func connectServer(ip string) (string, error) {
// 	if ip == "192.168.1.1" {
// 		return "Connected successfully", nil
// 	}
// 	return "", errors.New("Connection Failed! Target is offline")
// }

// func main() {
// 	message, err := connectServer("192.168.1.1")

// 	if err != nil {
// 		fmt.Println("[ERROR]:", err)
// 		return
// 	}
// 	fmt.Println("[SUCCESS:]", message)
// }

// func loginUser(user string, pass string) (bool, error) {
// 	if user == "admin" && pass == "admin" {
// 		return true, nil
// 	}
// 	return false, errors.New("Access Denied! Invalid credentials.")
// }

// func main() {
// 	const userName string = "admin"
// 	const password string = "admin1"

// 	message, err := loginUser(userName, password)

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Success:", message)
// }

//? defer

// func stealData() {
// 	fmt.Println("[1] Connecting to target server...")

// 	// defer
// 	defer fmt.Println("[3] Disconnecting and clearing tracks...")

// 	fmt.Println("[2] Downloading secret files...")
// }

// func main() {
// 	stealData()
// }

// func readSecretFile() {
// 	fmt.Println("[+] File opened: passwords.txt")
// 	defer fmt.Println("[-] File closed safely.")
// 	fmt.Println("[!] Extracting hashes: admin:21232f297a57a5a743894a0e4a801fc3")
// }

// func main(){
// 	readSecretFile()
// }

// func attackTarget(target string) {
// 	fmt.Println("[*] Attacking:", target)
// }

// func main() {
// 	go attackTarget("Server 1")
// 	go attackTarget("Server 2")

// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Attack complete!")
// }
