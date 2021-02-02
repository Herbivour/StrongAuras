package main

import "github.com/gonutz/d3d9"

type vertex struct {
	X  float32
	Y  float32
	TX float32
	TY float32
}

func (t vertex) WriteToVBuff(start int, v *d3d9.VertexBuffer) {
	data, err := v.Lock(0, 0, d3d9.LOCK_DISCARD)
	check(err)
	data.SetFloat32s(start, []float32{t.X, t.Y, t.TX, t.TY})
	check(v.Unlock())
}

// Stride - the size of the vertex in bytes
func (t vertex) Stride() uint {
	return 16
}

func (t vertex) GetVertexElement() []d3d9.VERTEXELEMENT {
	// Describe to D3D how our vertex layout works
	return []d3d9.VERTEXELEMENT{
		{
			Stream:     0,
			Offset:     0,
			Type:       d3d9.DECLTYPE_FLOAT2,
			Method:     d3d9.DECLMETHOD_DEFAULT,
			Usage:      d3d9.DECLUSAGE_POSITION,
			UsageIndex: 0,
		},
		{
			Stream:     0,
			Offset:     8,
			Type:       d3d9.DECLTYPE_FLOAT2,
			Method:     d3d9.DECLMETHOD_DEFAULT,
			Usage:      d3d9.DECLUSAGE_TEXCOORD,
			UsageIndex: 0,
		},
		d3d9.DeclEnd(),
	}
}
