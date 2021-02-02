package main

import (
	"os"

	"github.com/hpcloud/tail"
)

func processLogLine(text string) {
	// fmt.Println("Log Line:", text)
	if len(cfg.Indicators) > 0 {
		for _, i := range cfg.Indicators {
			i.ProcessLogLine(text)
		}
	}
}

func tailLog(filename string) {
	t, _ := tail.TailFile(filename, tail.Config{
		Follow: true,
		Location: &tail.SeekInfo{
			Whence: os.SEEK_END,
		},
	})
	for line := range t.Lines {
		processLogLine(line.Text)
	}
}
