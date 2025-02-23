//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Film struct {
	FilmID             uint16 `sql:"primary_key"`
	Title              string
	Description        *string
	ReleaseYear        *int16
	LanguageID         uint8
	OriginalLanguageID *uint8
	RentalDuration     uint8
	RentalRate         float64
	Length             *uint16
	ReplacementCost    float64
	Rating             *FilmRating
	SpecialFeatures    *string
	LastUpdate         time.Time
}
