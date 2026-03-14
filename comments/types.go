package comments

import (
	"time"

	"github.com/mathisbot/notionary-go/blocks"
)

type CommentParentType string

const (
	CommentParentTypePageID  CommentParentType = "page_id"
	CommentParentTypeBlockID CommentParentType = "block_id"
)

type CommentParent struct {
	Type    CommentParentType `json:"type"`
	PageID  string            `json:"page_id,omitempty"`
	BlockID string            `json:"block_id,omitempty"`
}

func NewPageCommentParent(pageID string) CommentParent {
	return CommentParent{Type: CommentParentTypePageID, PageID: pageID}
}

func NewBlockCommentParent(blockID string) CommentParent {
	return CommentParent{Type: CommentParentTypeBlockID, BlockID: blockID}
}

type CommentAttachmentCategory string

const (
	CommentAttachmentCategoryAudio        CommentAttachmentCategory = "audio"
	CommentAttachmentCategoryImage        CommentAttachmentCategory = "image"
	CommentAttachmentCategoryPDF          CommentAttachmentCategory = "pdf"
	CommentAttachmentCategoryProductivity CommentAttachmentCategory = "productivity"
	CommentAttachmentCategoryVideo        CommentAttachmentCategory = "video"
)

type FileWithExpiry struct {
	URL        string    `json:"url"`
	ExpiryTime time.Time `json:"expiry_time"`
}

type CommentAttachment struct {
	Category CommentAttachmentCategory `json:"category"`
	File     FileWithExpiry            `json:"file"`
}

type CommentAttachmentInputType string

const (
	CommentAttachmentInputTypeFileUpload CommentAttachmentInputType = "file_upload"
)

type CommentAttachmentInput struct {
	Type         CommentAttachmentInputType `json:"type"`
	FileUploadID string                     `json:"file_upload_id"`
}

func NewCommentAttachmentInput(fileUploadID string) CommentAttachmentInput {
	return CommentAttachmentInput{
		Type:         CommentAttachmentInputTypeFileUpload,
		FileUploadID: fileUploadID,
	}
}

type CommentDisplayNameType string

const (
	CommentDisplayNameTypeIntegration CommentDisplayNameType = "integration"
	CommentDisplayNameTypeUser        CommentDisplayNameType = "user"
	CommentDisplayNameTypeCustom      CommentDisplayNameType = "custom"
)

type CustomDisplayName struct {
	Name string `json:"name"`
}

type CommentDisplayNameInput struct {
	Type   CommentDisplayNameType `json:"type"`
	Custom *CustomDisplayName     `json:"custom,omitempty"`
}

func IntegrationDisplayNameInput() CommentDisplayNameInput {
	return CommentDisplayNameInput{Type: CommentDisplayNameTypeIntegration}
}

func UserDisplayNameInput() CommentDisplayNameInput {
	return CommentDisplayNameInput{Type: CommentDisplayNameTypeUser}
}

func CustomCommentDisplayNameInput(name string) CommentDisplayNameInput {
	return CommentDisplayNameInput{
		Type:   CommentDisplayNameTypeCustom,
		Custom: &CustomDisplayName{Name: name},
	}
}

type CommentDisplayName struct {
	Type         CommentDisplayNameType `json:"type"`
	ResolvedName string                 `json:"resolved_name"`
}

type UserRef struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}

type Comment struct {
	Object         string              `json:"object"`
	ID             string              `json:"id"`
	Parent         CommentParent       `json:"parent"`
	DiscussionID   string              `json:"discussion_id"`
	CreatedTime    time.Time           `json:"created_time"`
	LastEditedTime time.Time           `json:"last_edited_time"`
	CreatedBy      UserRef             `json:"created_by"`
	RichText       []blocks.RichText   `json:"rich_text"`
	Attachments    []CommentAttachment `json:"attachments"`
	DisplayName    *CommentDisplayName `json:"display_name,omitempty"`
}

type commentListResponse struct {
	Object     string    `json:"object"`
	Results    []Comment `json:"results"`
	NextCursor *string   `json:"next_cursor,omitempty"`
	HasMore    bool      `json:"has_more"`
}

type CommentCreateRequest struct {
	RichText     []blocks.RichText        `json:"rich_text"`
	Parent       *CommentParent           `json:"parent,omitempty"`
	DiscussionID string                   `json:"discussion_id,omitempty"`
	DisplayName  *CommentDisplayNameInput `json:"display_name,omitempty"`
	Attachments  []CommentAttachmentInput `json:"attachments,omitempty"`
}

func NewPageCommentRequest(pageID string, richText []blocks.RichText) CommentCreateRequest {
	parent := NewPageCommentParent(pageID)
	return CommentCreateRequest{
		RichText: richText,
		Parent:   &parent,
	}
}

func NewBlockCommentRequest(blockID string, richText []blocks.RichText) CommentCreateRequest {
	parent := NewBlockCommentParent(blockID)
	return CommentCreateRequest{
		RichText: richText,
		Parent:   &parent,
	}
}

func NewDiscussionCommentRequest(discussionID string, richText []blocks.RichText) CommentCreateRequest {
	return CommentCreateRequest{
		RichText:     richText,
		DiscussionID: discussionID,
	}
}
