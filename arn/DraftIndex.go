package arn

import (
	"errors"
	"reflect"
)

// DraftIndex has references to unpublished drafts a user created.
type DraftIndex struct {
	UserID       string  `json:"userId" primary:"true"`
	GroupID      string  `json:"groupId"`
	SoundTrackID string  `json:"soundTrackId"`
	CompanyID    string  `json:"companyId"`
	QuoteID      string  `json:"quoteId"`
	CharacterID  string  `json:"characterId"`
	AnimeID      AnimeID `json:"animeId"`
	AMVID        string  `json:"amvId"`
}

// NewDraftIndex ...
func NewDraftIndex(userID UserID) *DraftIndex {
	return &DraftIndex{
		UserID: userID,
	}
}

// DraftID gets the ID for the given type name.
func (index *DraftIndex) DraftID(typeName string) (string, error) {
	v := reflect.ValueOf(index).Elem()
	fieldValue := v.FieldByName(typeName + "ID")

	if !fieldValue.IsValid() {
		return "", errors.New("Invalid draft index ID type: " + typeName)
	}

	return fieldValue.String(), nil
}

// SetDraftID sets the ID for the given type name.
func (index *DraftIndex) SetDraftID(typeName string, id string) error {
	v := reflect.ValueOf(index).Elem()
	fieldValue := v.FieldByName(typeName + "ID")

	if !fieldValue.IsValid() {
		return errors.New("Invalid draft index ID type: " + typeName)
	}

	fieldValue.SetString(id)
	return nil
}

// GetID returns the ID.
func (index *DraftIndex) GetID() string {
	return index.UserID
}

// GetDraftIndex ...
func GetDraftIndex(id string) (*DraftIndex, error) {
	obj, err := DB.Get("DraftIndex", id)

	if err != nil {
		return nil, err
	}

	return obj.(*DraftIndex), nil
}
