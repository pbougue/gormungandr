package serializer

import "github.com/canaltp/gonavitia"
import "github.com/canaltp/gonavitia/pbnavitia"

func NewGeoJson(pb *pbnavitia.Section) *gonavitia.GeoJson {
	if pb == nil {
		return nil
	}
	g := gonavitia.GeoJson{
		Type:        "LineString",
		Properties:  []map[string]interface{}{{"length": pb.GetLength()}},
		Coordinates: make([][]float64, 0, len(pb.Shape)),
	}

	for _, pb_coord := range pb.Shape {
		coord := []float64{pb_coord.GetLon(), pb_coord.GetLat()}
		g.Coordinates = append(g.Coordinates, coord)
	}

	return &g
}

func NewGeoJsonMultistring(pb *pbnavitia.MultiLineString) *gonavitia.GeoJsonMultilineString {
	if pb == nil {
		return nil
	}
	g := gonavitia.GeoJsonMultilineString{
		Type:        "MultiLineString",
		Coordinates: make([][][]float64, 0, len(pb.Lines)),
	}

	for _, line := range pb.Lines {
		l := make([][]float64, 0, len(line.Coordinates))
		for _, coord := range line.Coordinates {
			l = append(l, []float64{coord.GetLon(), coord.GetLat()})
		}
		g.Coordinates = append(g.Coordinates, l)
	}

	return &g
}
