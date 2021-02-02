package main

import (
	"fmt"
	"log"
	"runtime"
	"runtime/debug"

	"github.com/gonutz/d3d9"
	"github.com/gonutz/w32"
)

const (
	vertexFormat = d3d9.FVF_XYZRHW | d3d9.FVF_DIFFUSE | d3d9.FVF_TEX1 //d3d9.FVF_XYZRHW | d3d9.FVF_DIFFUSE | d3d9.FVF_TEX1
	vertexStride = 16
)

var (
	eqWidth  int32
	eqHeight int32
)

func init() {
	runtime.LockOSThread()
}

const PROCESS_ALL_ACCESS = w32.PROCESS_CREATE_PROCESS | w32.PROCESS_CREATE_THREAD | w32.PROCESS_DUP_HANDLE | w32.PROCESS_QUERY_INFORMATION | w32.PROCESS_QUERY_LIMITED_INFORMATION | w32.PROCESS_SET_INFORMATION | w32.PROCESS_SET_QUOTA | w32.PROCESS_SUSPEND_RESUME | w32.PROCESS_TERMINATE | w32.PROCESS_VM_OPERATION | w32.PROCESS_VM_READ | w32.PROCESS_VM_WRITE | w32.SYNCHRONIZE

type messageCallback func(window w32.HWND, msg uint32, w, l uintptr) uintptr

var eq w32.HWND
var mainWindow w32.HWND

var tex *d3d9.Texture

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic: %v\nstack\n---\n%s\n---\n", err, debug.Stack())
			msg := fmt.Sprint("panic: ", err)
			const MB_TOPMOST = 0x00040000
			w32.MessageBox(0, msg, "Error", w32.MB_OK|w32.MB_ICONERROR|MB_TOPMOST)
		}
	}()
	loadConfig()
	go tailLog(
		fmt.Sprintf("%vLogs\\eqlog_%v_%v.txt", cfg.EqFolder, cfg.Character, cfg.Server),
	)

	initMainWindow()
	initD3D9()
	initIndicators()
	if device == nil {
		panic("Failed to initialize Direct3D 9 Device.")
	}
	defer device.Release()

	runMessageHandler()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
