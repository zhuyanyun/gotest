package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

var username = flag.String("username", "kb", "Alpenhorn username")
var debug = flag.Bool("debug", false, "Turn on debug mode")
var latency = flag.Duration("latency", 150*time.Millisecond, "latency to coordinator")

var persistPath = flag.String("persist", "E:\\ff", "persistent data directory")

func main() {
	fmt.Println(filepath.Join("c:", "aa", "bb", "cc.txt"))

	fmt.Println(filepath.Join(*persistPath, "cc.txt"))

	flag.Parse()

	if *username == "" {
		fmt.Println("no username specified")
		os.Exit(1)
	}

	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.HomeDir)
	confHome := filepath.Join(u.HomeDir, ".vuvuzela")
	fmt.Println(confHome)
}
