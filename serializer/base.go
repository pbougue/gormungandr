package serializer

import (
	"strings"

	"github.com/canaltp/gonavitia"
	"github.com/canaltp/gonavitia/pbnavitia"
)

func NewError(pb *pbnavitia.Error) *gonavitia.Error {
	if pb == nil {
		return nil
	}
	id := pb.Id.Enum().String()
	return &gonavitia.Error{
		Id:      &id,
		Message: pb.Message,
	}
}

func NewCode(pb *pbnavitia.Code) *gonavitia.Code {
	if pb == nil {
		return nil
	}
	return &gonavitia.Code{
		Type:  pb.Type,
		Value: pb.Value,
	}
}

func NewPlace(pb *pbnavitia.PtObject) *gonavitia.Place {
	if pb == nil {
		return nil
	}
	t := strings.ToLower(pb.EmbeddedType.String())
	place := gonavitia.Place{
		Id:           pb.Uri,
		Name:         pb.Name,
		EmbeddedType: &t,
		Quality:      pb.Quality,
		StopPoint:    NewStopPoint(pb.StopPoint),
		StopArea:     NewStopArea(pb.StopArea),
		Admin:        NewAdmin(pb.AdministrativeRegion),
		Address:      NewAddress(pb.Address),
	}
	return &place
}

func NewAdmin(pb *pbnavitia.AdministrativeRegion) *gonavitia.Admin {
	if pb == nil {
		return nil
	}
	admin := gonavitia.Admin{
		Id:      pb.Uri,
		Name:    pb.Name,
		Label:   pb.Label,
		Coord:   NewCoord(pb.Coord),
		Insee:   pb.Insee,
		ZipCode: pb.ZipCode,
	}
	return &admin
}

func NewCoord(pb *pbnavitia.GeographicalCoord) *gonavitia.Coord {
	if pb == nil {
		return nil
	}
	coord := gonavitia.Coord{
		Lat: pb.GetLat(),
		Lon: pb.GetLon(),
	}
	return &coord
}

func NewStopPoint(pb *pbnavitia.StopPoint) *gonavitia.StopPoint {
	if pb == nil {
		return nil
	}
	sp := gonavitia.StopPoint{
		Id:       pb.Uri,
		Name:     pb.Name,
		Label:    pb.Label,
		Coord:    NewCoord(pb.Coord),
		Admins:   make([]*gonavitia.Admin, len(pb.AdministrativeRegions)),
		StopArea: NewStopArea(pb.StopArea),
		Codes:    make([]*gonavitia.Code, 0, len(pb.Codes)),
	}
	for _, pb_admin := range pb.AdministrativeRegions {
		sp.Admins = append(sp.Admins, NewAdmin(pb_admin))
	}
	for _, code := range pb.Codes {
		sp.Codes = append(sp.Codes, NewCode(code))
	}
	return &sp
}

func NewStopArea(pb *pbnavitia.StopArea) *gonavitia.StopArea {
	if pb == nil {
		return nil
	}
	sa := gonavitia.StopArea{
		Id:       pb.Uri,
		Name:     pb.Name,
		Label:    pb.Label,
		Timezone: pb.Timezone,
		Coord:    NewCoord(pb.Coord),
		Admins:   make([]*gonavitia.Admin, 0, len(pb.AdministrativeRegions)),
		Codes:    make([]*gonavitia.Code, 0, len(pb.Codes)),
	}
	for _, pb_admin := range pb.AdministrativeRegions {
		sa.Admins = append(sa.Admins, NewAdmin(pb_admin))
	}
	for _, code := range pb.Codes {
		sa.Codes = append(sa.Codes, NewCode(code))
	}
	for _, sp := range pb.StopPoints {
		sa.StopPoints = append(sa.StopPoints, NewStopPoint(sp))
	}
	return &sa
}

func NewAddress(pb *pbnavitia.Address) *gonavitia.Address {
	if pb == nil {
		return nil
	}
	address := gonavitia.Address{
		Id:          pb.Uri,
		Name:        pb.Name,
		Label:       pb.Label,
		Coord:       NewCoord(pb.Coord),
		HouseNumber: pb.HouseNumber,
		Admins:      make([]*gonavitia.Admin, 0),
	}
	for _, pb_admin := range pb.AdministrativeRegions {
		address.Admins = append(address.Admins, NewAdmin(pb_admin))
	}
	return &address
}
