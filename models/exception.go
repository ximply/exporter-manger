package models

import (
	"github.com/ximply/exporter-manger/structs"
	"github.com/ximply/exporter-manger/uploads"
)

func ProcessException(ex *structs.ExceptionItem) bool {
	ok, _ := uploads.ExceptionQueue().Put(ex)
	return ok
}
