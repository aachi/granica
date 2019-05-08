/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package models

import (
	"encoding/json"
)

const (
	// PointGeoType is type name for Point Geometry.
	PointGeoType = "Point"
	// LineGeoType is type name for Line Geometry.
	LineGeoType = "Line"
	// PolygonGeoType is type name for Poligon Geometry.
	PolygonGeoType = "Polygon"
)

// Point describes a geospatial point.
type Point struct {
	Coordinates []float64
}

// Line describes a geospatial line.
type Line struct {
	Points [][]float64
}

// Polygon describes a geospatial polygons.
type Polygon struct {
	Lines [][][]float64
}

// Geometry describes a geospatial type.
type Geometry struct {
	Type        string
	Coordinates json.RawMessage
	Point       Point
	Line        Line
	Polygon     Polygon
}

// Unmarshal defines an unmarshaller for Geometry type.
func (g *Geometry) Unmarshal(b []byte) error {

	type Alias Geometry
	aux := (*Alias)(g)

	err := json.Unmarshal(b, &aux)

	if err != nil {
		return err
	}

	switch g.Type {
	case PointGeoType:
		err = json.Unmarshal(g.Coordinates, &g.Point.Coordinates)
	case LineGeoType:
		err = json.Unmarshal(g.Coordinates, &g.Line.Points)
	case PolygonGeoType:
		err = json.Unmarshal(g.Coordinates, &g.Polygon.Lines)
	}

	g.Coordinates = []byte(nil)

	return err
}

// Marshal defines an marshaller for Geometry type.
func (g Geometry) Marshal() ([]byte, error) {

	var raw json.RawMessage
	var err error

	switch g.Type {
	case PointGeoType:
		raw, err = json.Marshal(&g.Point.Coordinates)
	case LineGeoType:
		raw, err = json.Marshal(&g.Line.Points)
	case PolygonGeoType:
		raw, err = json.Marshal(&g.Polygon.Lines)
	}

	if err != nil {
		return nil, err
	}

	g.Coordinates = raw

	type Alias Geometry
	aux := (*Alias)(&g)

	return json.Marshal(aux)
}
