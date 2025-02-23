//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type FilmListRating string

const (
	FilmListRating_G    FilmListRating = "G"
	FilmListRating_Pg   FilmListRating = "PG"
	FilmListRating_Pg13 FilmListRating = "PG-13"
	FilmListRating_R    FilmListRating = "R"
	FilmListRating_Nc17 FilmListRating = "NC-17"
)

func (e *FilmListRating) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "G":
		*e = FilmListRating_G
	case "PG":
		*e = FilmListRating_Pg
	case "PG-13":
		*e = FilmListRating_Pg13
	case "R":
		*e = FilmListRating_R
	case "NC-17":
		*e = FilmListRating_Nc17
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for FilmListRating enum")
	}

	return nil
}

func (e FilmListRating) String() string {
	return string(e)
}
