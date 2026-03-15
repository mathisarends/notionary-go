package entity

import (
	"github.com/mathisbot/notionary-go/blocks"
	"github.com/mathisbot/notionary-go/user"
)

type EntityType string

const (
	EntityTypePage       EntityType = "page"
	EntityTypeDataSource EntityType = "data_source"
	EntityTypeDatabase   EntityType = "database"
)

// EntityResponseDto mirrors the Pydantic EntityResponseDto
type EntityResponseDto struct {
	Object         EntityType       `json:"object"`
	ID             string           `json:"id"`
	CreatedTime    string           `json:"created_time"`
	CreatedBy      user.PartialUser `json:"created_by"`
	LastEditedTime string           `json:"last_edited_time"`
	LastEditedBy   user.PartialUser `json:"last_edited_by"`
	Cover          *File            `json:"cover,omitempty"`
	Icon           *Icon            `json:"icon,omitempty"`
	Parent         Parent           `json:"parent"`
	InTrash        bool             `json:"in_trash"`
	URL            string           `json:"url"`
	PublicURL      *string          `json:"public_url,omitempty"`
}

type EntityUpdateRequest struct {
	Icon    *Icon `json:"icon,omitempty"`
	Cover   *File `json:"cover,omitempty"`
	InTrash *bool `json:"in_trash,omitempty"`
}

type Titled interface {
	GetTitle() []blocks.RichText
}

type Describable interface {
	GetDescription() []blocks.RichText
}

type URL struct {
	URL string `json:"url"`
}

type FileType string

const (
	FileTypeExternal   FileType = "external"
	FileTypeFile       FileType = "file"
	FileTypeFileUpload FileType = "file_upload"
)

type FileRef struct {
	URL        string  `json:"url"`
	ExpiryTime *string `json:"expiry_time,omitempty"`
}

type File struct {
	Type       FileType `json:"type"`
	External   *URL     `json:"external,omitempty"`
	File       *FileRef `json:"file,omitempty"`
	FileUpload *struct {
		ID string `json:"id"`
	} `json:"file_upload,omitempty"`
}