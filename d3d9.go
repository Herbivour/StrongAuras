package main

import (
	"fmt"
	"log"

	"github.com/gonutz/d3d9"
)

var device *d3d9.Device

func initD3D9() {
	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	if err != nil {
		panic(fmt.Sprint("unable to create Direct3D9 object: ", err))
	}
	defer d3d.Release()

	dev, _, err := d3d.CreateDevice(
		d3d9.ADAPTER_DEFAULT,
		d3d9.DEVTYPE_HAL,
		d3d9.HWND(mainWindow),
		d3d9.CREATE_HARDWARE_VERTEXPROCESSING,
		d3d9.PRESENT_PARAMETERS{
			Windowed:           1,
			SwapEffect:         d3d9.SWAPEFFECT_DISCARD,
			HDeviceWindow:      d3d9.HWND(mainWindow),
			MultiSampleQuality: d3d9.MULTISAMPLE_NONE,
			BackBufferFormat:   d3d9.FMT_A8R8G8B8,
			BackBufferWidth:    uint32(eqWidth),
			BackBufferHeight:   uint32(eqHeight),
		},
	)
	if err != nil {
		log.Fatal("Failed to create Direct3D 9 Device: ", err)
	}

	check(dev.SetRenderState(d3d9.RS_CULLMODE, uint32(d3d9.CULL_NONE)))
	dev.SetFVF(vertexFormat)
	dev.SetRenderState(d3d9.RS_ZENABLE, d3d9.ZB_FALSE)
	dev.SetRenderState(d3d9.RS_CULLMODE, d3d9.CULL_NONE)
	dev.SetRenderState(d3d9.RS_LIGHTING, 0)
	dev.SetRenderState(d3d9.RS_SRCBLEND, d3d9.BLEND_SRCALPHA)
	dev.SetRenderState(d3d9.RS_DESTBLEND, d3d9.BLEND_INVSRCALPHA)
	dev.SetRenderState(d3d9.RS_ALPHABLENDENABLE, 1)

	// texture filter for when zooming
	dev.SetSamplerState(0, d3d9.SAMP_MINFILTER, d3d9.TEXF_LINEAR)
	dev.SetSamplerState(0, d3d9.SAMP_MAGFILTER, d3d9.TEXF_LINEAR)

	dev.SetTextureStageState(0, d3d9.TSS_COLOROP, d3d9.TOP_MODULATE)
	dev.SetTextureStageState(0, d3d9.TSS_COLORARG1, d3d9.TA_TEXTURE)
	dev.SetTextureStageState(0, d3d9.TSS_COLORARG2, d3d9.TA_DIFFUSE)

	dev.SetTextureStageState(0, d3d9.TSS_ALPHAOP, d3d9.TOP_MODULATE)
	dev.SetTextureStageState(0, d3d9.TSS_ALPHAARG1, d3d9.TA_TEXTURE)
	dev.SetTextureStageState(0, d3d9.TSS_ALPHAARG2, d3d9.TA_DIFFUSE)

	dev.SetTextureStageState(1, d3d9.TSS_COLOROP, d3d9.TOP_DISABLE)
	dev.SetTextureStageState(1, d3d9.TSS_ALPHAOP, d3d9.TOP_DISABLE)

	device = dev
}
