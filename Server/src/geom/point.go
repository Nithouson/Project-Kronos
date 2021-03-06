package geom

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	pos Coord
}

func (p Point) GeomType() int {
	return KrPoint
}

func (p Point) GetPos() (float64, float64) {
	return p.pos.x, p.pos.y
}

func (p *Point) SetPos(_x float64, _y float64) {
	p.pos.x = _x
	p.pos.y = _y
}

func (p Point) ExportWKT() string {
	return "POINT (" + fmt.Sprintf("%f", p.pos.x) + " " + fmt.Sprintf("%f", p.pos.y) + ")"
}

func (p Point) ExportMap() map[string]interface{} {
	mj := make(map[string]interface{})
	mj["type"] = "Point"
	mj["coordinates"] = []float64{p.pos.x, p.pos.y}
	return mj
}

func (p Point) ExportGeoJSON() string {
	s, _ := json.Marshal(p.ExportMap())
	return string(s)
}
