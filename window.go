package main

import (
	"fmt"
	"syscall"
	"time"

	"github.com/gonutz/w32"
)

func msgHandler(hWnd w32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case w32.WM_DESTROY:
		w32.PostQuitMessage(0)
	default:
		return w32.DefWindowProc(hWnd, msg, wParam, lParam)
	}
	return 0
}

func initMainWindow() {
	// Find the first EQ window to attach our overlay to
	for ; eq == 0; eq = w32.FindWindow("_EverQuestwndclass", "EverQuest") {
		fmt.Println("Lookking for eqgame.exe")
		time.Sleep(1 * time.Second)
	}

	eqSize := w32.GetWindowRect(eq)
	eqWidth = eqSize.Right - eqSize.Left
	eqHeight = eqSize.Bottom - eqSize.Top

	windowProc := syscall.NewCallback(msgHandler)

	class := w32.WNDCLASSEX{
		WndProc:    windowProc,
		ClassName:  syscall.StringToUTF16Ptr("StrongAuras"),
		Background: w32.CreateSolidBrush(0),
		Style:      w32.CS_VREDRAW | w32.CS_HREDRAW,
	}
	atom := w32.RegisterClassEx(&class)
	if atom == 0 {
		panic("Failed to open main window")
	}

	mainWindow = w32.CreateWindowEx(
		w32.WS_EX_TOPMOST|w32.WS_EX_TRANSPARENT|w32.WS_EX_LAYERED,
		syscall.StringToUTF16Ptr("StrongAuras"),
		syscall.StringToUTF16Ptr("StrongAuras"),
		w32.WS_POPUP, //w32.WS_OVERLAPPEDWINDOW|w32.WS_VISIBLE,
		int(eqSize.Left), int(eqSize.Top),
		int(eqWidth), int(eqHeight),
		0, 0, 0, nil,
	)
	if mainWindow == 0 {
		panic("Failed to open main window")
	}
	w32.SetLayeredWindowAttributes(mainWindow, 0, 0, w32.LWA_ALPHA)
	w32.SetLayeredWindowAttributes(mainWindow, 0, 0, w32.LWA_COLORKEY)

	w32.SetForegroundWindow(eq)
	w32.ShowWindow(mainWindow, w32.SW_SHOW)
	w32.SetWindowText(mainWindow, "Strong Auras")
}
