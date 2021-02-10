package main

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/gonutz/w32"
	"github.com/mrherbivour/StrongAuras/config"
	"github.com/mrherbivour/StrongAuras/gfx"
)

var cfg config.Config

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic: %v\nstack\n---\n%s\n---\n", err, debug.Stack())
			msg := fmt.Sprint("panic: ", err)
			const MB_TOPMOST = 0x00040000
			w32.MessageBox(0, msg, "Error", w32.MB_OK|w32.MB_ICONERROR|MB_TOPMOST)
		}
	}()
	cfg = config.LoadConfig("./config.json")
	go tailLog(
		fmt.Sprintf("%vLogs\\eqlog_%v_%v.txt", cfg.EqFolder, cfg.Character, cfg.Server),
	)
	gfx.Init(&cfg)
}
