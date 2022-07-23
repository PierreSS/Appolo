package main

import (
	"Strategies/ichimoku"
	"Strategies/stochastic"
	"Strategies/supertrend"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/robfig/cron"
)

func runCmd(cmd *exec.Cmd) {
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}

func main() {
	c := cron.New()

	c.AddFunc("10 0 * * * *", func() {
		go supertrend.Worker()
		go stochastic.Worker()
		go ichimoku.Worker()

		cmd := exec.Command("python3.10", "PPO_supertrend/main.py")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		go runCmd(cmd)
	})

	c.Start()
	for {
		time.Sleep(time.Second)
	}
}
