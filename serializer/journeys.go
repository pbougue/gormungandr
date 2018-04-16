package serializer

import "github.com/CanalTP/gonavitia"
import "github.com/CanalTP/gonavitia/pbnavitia"
import "time"
import "strings"
import "github.com/golang/protobuf/proto"

func NewJourneysReponse(pb *pbnavitia.Response) *gonavitia.JourneysResponse {
	if pb == nil {
		return nil
	}
	r := gonavitia.JourneysResponse{}
	for _, pb_journey := range pb.Journeys {
		r.Journeys = append(r.Journeys, NewJourney(pb_journey))
	}
	return &r
}

func NewJourney(pb *pbnavitia.Journey) *gonavitia.Journey {
	if pb == nil {
		return nil
	}
	journey := gonavitia.Journey{
		From:              NewPlace(pb.Origin),
		To:                NewPlace(pb.Destination),
		Duration:          pb.GetDuration(),
		NbTransfers:       pb.GetNbTransfers(),
		DepartureDateTime: gonavitia.NavitiaDatetime(time.Unix(int64(pb.GetDepartureDateTime()), 0)),
		ArrivalDateTime:   gonavitia.NavitiaDatetime(time.Unix(int64(pb.GetArrivalDateTime()), 0)),
		RequestedDateTime: gonavitia.NavitiaDatetime(time.Unix(int64(pb.GetRequestedDateTime()), 0)),
		Status:            pb.GetMostSeriousDisruptionEffect(),
		Durations:         NewDurations(pb.Durations),
		Distances:         NewDistances(pb.Distances),
		Tags:              make([]string, 0),
	}
	for _, pb_section := range pb.Sections {
		journey.Sections = append(journey.Sections, NewSection(pb_section))
	}
	return &journey
}

func NewSection(pb *pbnavitia.Section) *gonavitia.Section {
	if pb == nil {
		return nil
	}
	var mode *string
	if sn := pb.StreetNetwork; sn != nil {
		m := strings.ToLower(sn.Mode.String())
		mode = &m
	}
	var transferType *string
	if pb.TransferType != nil {
		t := strings.ToLower(pb.TransferType.String())
		transferType = &t
	}
	section := gonavitia.Section{
		Id:                pb.GetId(),
		From:              NewPlace(pb.Origin),
		To:                NewPlace(pb.Destination),
		DepartureDateTime: gonavitia.NavitiaDatetime(time.Unix(int64(pb.GetBeginDateTime()), 0)),
		ArrivalDateTime:   gonavitia.NavitiaDatetime(time.Unix(int64(pb.GetEndDateTime()), 0)),
		Duration:          pb.GetDuration(),
		Type:              strings.ToLower(pb.GetType().String()),
		GeoJson:           NewGeoJson(pb),
		Mode:              mode,
		TransferType:      transferType,
		DisplayInfo:       NewPtDisplayInfoForVJ(pb.PtDisplayInformations),
		Co2Emission:       NewCo2Emission(pb.Co2Emission),
		AdditionalInfo:    NewAdditionalInformations(pb.AdditionalInformations),
		Links:             NewLinksFromUris(pb.PtDisplayInformations),
	}

	return &section
}

func NewDurations(pb *pbnavitia.Durations) *gonavitia.Durations {
	if pb == nil {
		return nil
	}
	durations := gonavitia.Durations{
		Total:       pb.GetTotal(),
		Walking:     pb.GetWalking(),
		Bike:        pb.GetBike(),
		Car:         pb.GetCar(),
		Ridesharing: pb.GetRidesharing(),
	}
	return &durations
}

func NewDistances(pb *pbnavitia.Distances) *gonavitia.Distances {
	if pb == nil {
		return nil
	}
	distances := gonavitia.Distances{
		Walking:     pb.GetWalking(),
		Bike:        pb.GetBike(),
		Car:         pb.GetCar(),
		Ridesharing: pb.GetRidesharing(),
	}
	return &distances
}

func NewCo2Emission(pb *pbnavitia.Co2Emission) *gonavitia.Amount {
	if pb == nil {
		return nil
	}
	co2 := gonavitia.Amount{
		Value: pb.GetValue(),
		Unit:  pb.GetUnit(),
	}
	return &co2

}

func NewLinksFromUris(pb *pbnavitia.PtDisplayInfo) []gonavitia.Link {
	if pb == nil || pb.Uris == nil {
		return nil
	}
	uris := pb.Uris
	res := make([]gonavitia.Link, 0)
	res = appendLinksFromUri(uris.Company, "company", &res)
	res = appendLinksFromUri(uris.VehicleJourney, "vehicle_journey", &res)
	res = appendLinksFromUri(uris.Line, "line", &res)
	res = appendLinksFromUri(uris.Route, "route", &res)
	res = appendLinksFromUri(uris.CommercialMode, "commercial_mode", &res)
	res = appendLinksFromUri(uris.PhysicalMode, "physical_mode", &res)
	res = appendLinksFromUri(uris.Network, "Network", &res)
	res = appendLinksFromUri(uris.Note, "note", &res)
	res = appendLinksFromUri(uris.JourneyPattern, "journey_pattern", &res)
	return res
}

func appendLinksFromUri(pb *string, typ string, links *[]gonavitia.Link) []gonavitia.Link {
	if pb == nil {
		return *links
	}
	return append(*links, gonavitia.Link{Id: pb, Type: proto.String(typ)})
}
