package main

import (
	"fmt"
	logrus "github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

func runAnfWait() int {
	time.Sleep(time.Second * 1)

	return 10
}

func main() {
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error while opening", err)
	}

	logrus.Info("run")
	log.SetOutput(file)

	for i := 0; i < 5; i++ {
		a := runAnfWait()
		log.Println("runAndWaitFinished...")
		log.Println("result", a)
	}

	fmt.Println("done")
}
