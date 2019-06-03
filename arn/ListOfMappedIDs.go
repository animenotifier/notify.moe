package arn

// import (
// 	"github.com/akyoto/color"
// )

// // ListOfMappedIDs ...
// type ListOfMappedIDs struct {
// 	IDList []*MappedID `json:"idList"`
// }

// // MappedID ...
// type MappedID struct {
// 	Type string `json:"type"`
// 	ID   string `json:"id"`
// }

// // Append appends the given mapped ID to the end of the list.
// func (idList *ListOfMappedIDs) Append(typeName string, id string) {
// 	idList.IDList = append(idList.IDList, &MappedID{
// 		Type: typeName,
// 		ID:   id,
// 	})
// }

// // Resolve ...
// func (idList *ListOfMappedIDs) Resolve() []interface{} {
// 	var data []interface{}

// 	for _, mapped := range idList.IDList {
// 		obj, err := DB.Get(mapped.Type, mapped.ID)

// 		if err != nil {
// 			color.Red(err.Error())
// 			continue
// 		}

// 		data = append(data, obj)
// 	}

// 	return data
// }
