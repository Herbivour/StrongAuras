package main

import "github.com/gonutz/d3d9"

var drawTriangle bool = true

func buildVertBuff() {
	var ts []vertex
	for _, ind := range cfg.Indicators {
		ts = append(ts, ind.Triangles...)
	}
	vb, err := device.CreateVertexBuffer(
		uint(uint(len(ts))*ts[0].Stride()),
		0,
		d3d9.FVF_XYZRHW|d3d9.FVF_DIFFUSE,
		d3d9.POOL_MANAGED,
		0,
	)
	check(err)
	defer vb.Release()
	for i, t := range ts {
		t.WriteToVBuff(i*int(t.Stride()), vb)
	}
	check(device.SetStreamSource(0, vb, 0, ts[0].Stride()))
	decl, err := device.CreateVertexDeclaration(ts[0].GetVertexElement())
	check(err)
	defer decl.Release()
	check(device.SetVertexDeclaration(decl))
}

func render() {
	check(device.Clear(nil, d3d9.CLEAR_TARGET, 0, 0, 0))
	check(device.BeginScene())
	for _, ind := range cfg.Indicators {
		if ind.Visable {
			if ind.SpriteSheet != nil && *ind.SpriteSheet != "" {
				icon := texturePool[*ind.SpriteSheet]
				if icon != nil {
					check(device.SetTexture(0, icon))
				}
			}
			check(
				device.DrawPrimitive(d3d9.PT_TRIANGLELIST, uint(ind.Id*6), 2),
			)
		}
	}
	check(device.EndScene())
	check(device.Present(nil, nil, 0, nil))

}
