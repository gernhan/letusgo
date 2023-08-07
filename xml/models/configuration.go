package models

import (
	"github.com/gernhan/xml/concurrent"
	"github.com/gernhan/xml/concurrent/atomics"
)

type XmlGenerationConfiguration struct {
	MaxID int64 `json:"maxID"`
	MinID int64 `json:"minID"`

	GlobalConfig *XmlGenerationGlobalConfiguration `json:"globalConfig"`
}

type XmlGenerationGlobalConfiguration struct {
	BillRunId int64 `json:"billRunId"`

	FileCount         *atomics.AtomicInteger `json:"fileCount"`
	TarCount          *atomics.AtomicInteger `json:"tarCount"`
	UploadedFileCount *concurrent.Map        `json:"uploadedFileCount"`

	TimePrefix string `json:"timePrefix"`

	BatchInvoices int `json:"batchInvoices"`
	BatchFiles    int `json:"batchFiles"`
}
