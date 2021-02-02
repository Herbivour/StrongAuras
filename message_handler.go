package main

import (
	"time"

	"github.com/gonutz/w32"
)

func runMessageHandler() {
	var msg w32.MSG
	var lastRender time.Time = time.Now()
	var lastReposition time.Time = time.Now()
	w32.PeekMessage(&msg, 0, 0, 0, w32.PM_NOREMOVE)
	for msg.Message != w32.WM_QUIT {
		if w32.PeekMessage(&msg, 0, 0, 0, w32.PM_REMOVE) {
			w32.TranslateMessage(&msg)
			w32.DispatchMessage(&msg)
		} else {
			now := time.Now()
			if now.Sub(lastReposition) > time.Duration(5*time.Second) {
				if !w32.IsWindow(eq) {
					return
				}
				eqSize := w32.GetWindowRect(eq)
				eqWidth = eqSize.Right - eqSize.Left
				eqHeight = eqSize.Bottom - eqSize.Top
				w32.MoveWindow(mainWindow, int(eqSize.Left), int(eqSize.Top), int(eqWidth), int(eqHeight), false)
				lastReposition = now
			}
			if now.Sub(lastRender) > time.Duration(100*time.Millisecond) {
				render()
				lastRender = now
			}
		}
	}
}
