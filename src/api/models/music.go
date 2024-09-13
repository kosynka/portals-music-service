package models

type Music struct {
	TITLE string `json:"title"`
	URL   string `json:"url"`
	SIZE  int64  `json:"size"`
	ETAG  string `json:"etag"`
}
