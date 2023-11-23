package data

import (
	"github.com/Oudwins/stackifyer/db"
)

type FileObject struct {
	ID        string `json:"id,omitempty"`
	CreatedAt int    `json:"created_at,omitempty"`
	UserId    string `json:"userid,omitempty"`
	Ftype     string `json:"ftype"` // file,image
	Filename  string `json:"filename,omitempty"`
	Key       string `json:"key,omitempty"`
	KeySmall  string `json:"key_small,omitempty"`
}

func PostNewFileEntry(f *FileObject) (inserted *FileObject, err error) {
	var res []FileObject
	if err = db.DbClient.DB.From("files").Insert(f).Execute(&res); err != nil {
		return nil, err
	}

	return &res[0], nil
}

func PostNewFileEntries(f *[]FileObject) (inserted *[]FileObject, err error) {
	var res []FileObject
	err = db.DbClient.DB.From("files").Insert(f).Execute(&res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
