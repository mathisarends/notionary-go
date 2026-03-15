package entity

import (
	"context"
	"fmt"

	fileupload "github.com/mathisbot/notionary-go/file_upload"
	notionhttp "github.com/mathisbot/notionary-go/http"
)

type Entity struct {
    ID             string
    CreatedTime    string
    LastEditedTime string
    InTrash        bool
    URL            string
    PublicURL      *string

    emojiIcon       *string
    externalIconURL *string
    coverImageURL   *string

    http       *notionhttp.Client
    fileUpload *fileupload.Client
    meta       MetadataClient
}

func New(dto EntityResponseDto, meta MetadataClient, http *notionhttp.Client) Entity {
    e := Entity{
        ID:             dto.ID,
        CreatedTime:    dto.CreatedTime,
        LastEditedTime: dto.LastEditedTime,
        InTrash:        dto.InTrash,
        URL:            dto.URL,
        PublicURL:      dto.PublicURL,
        meta:           meta,
        http:           http,
        fileUpload:     fileupload.New(http),
    }
    e.emojiIcon = extractEmojiIcon(dto)
    e.externalIconURL = extractExternalIconURL(dto)
    e.coverImageURL = extractCoverImageURL(dto)
    return e
}

func (e *Entity) SetEmojiIcon(ctx context.Context, emoji string) error {
    resp, err := e.meta.PatchEmojiIcon(ctx, emoji)
    if err != nil {
        return err
    }
    e.emojiIcon = extractEmojiIcon(*resp)
    e.externalIconURL = extractExternalIconURL(*resp)
    return nil
}

func (e *Entity) SetExternalIcon(ctx context.Context, url string) error {
    resp, err := e.meta.PatchExternalIcon(ctx, url)
    if err != nil {
        return err
    }
    e.emojiIcon = extractEmojiIcon(*resp)
    e.externalIconURL = extractExternalIconURL(*resp)
    return nil
}

func (e *Entity) SetIconFromFileUpload(ctx context.Context, fileUploadID string) error {
    resp, err := e.meta.PatchIconFromFileUpload(ctx, fileUploadID)
    if err != nil {
        return err
    }
    e.emojiIcon = extractEmojiIcon(*resp)
    e.externalIconURL = extractExternalIconURL(*resp)
    return nil
}

func (e *Entity) RemoveIcon(ctx context.Context) error {
    if err := e.meta.RemoveIcon(ctx); err != nil {
        return err
    }
    e.emojiIcon = nil
    e.externalIconURL = nil
    return nil
}

func (e *Entity) SetExternalCover(ctx context.Context, url string) error {
    resp, err := e.meta.PatchExternalCover(ctx, url)
    if err != nil {
        return err
    }
    e.coverImageURL = extractCoverImageURL(*resp)
    return nil
}

func (e *Entity) SetCoverFromFileUpload(ctx context.Context, fileUploadID string) error {
    resp, err := e.meta.PatchCoverFromFileUpload(ctx, fileUploadID)
    if err != nil {
        return err
    }
    e.coverImageURL = extractCoverImageURL(*resp)
    return nil
}

func (e *Entity) RemoveCover(ctx context.Context) error {
    if err := e.meta.RemoveCover(ctx); err != nil {
        return err
    }
    e.coverImageURL = nil
    return nil
}

func (e *Entity) MoveToTrash(ctx context.Context) error {
    if e.InTrash {
        return fmt.Errorf("entity %s is already in trash", e.ID)
    }
    resp, err := e.meta.MoveToTrash(ctx)
    if err != nil {
        return err
    }
    e.InTrash = resp.InTrash
    return nil
}

func (e *Entity) RestoreFromTrash(ctx context.Context) error {
    if !e.InTrash {
        return fmt.Errorf("entity %s is not in trash", e.ID)
    }
    resp, err := e.meta.RestoreFromTrash(ctx)
    if err != nil {
        return err
    }
    e.InTrash = resp.InTrash
    return nil
}

func (e *Entity) EmojiIcon() (string, bool) {
    if e.emojiIcon == nil {
        return "", false
    }
    return *e.emojiIcon, true
}

func (e *Entity) ExternalIconURL() (string, bool) {
    if e.externalIconURL == nil {
        return "", false
    }
    return *e.externalIconURL, true
}

func (e *Entity) CoverImageURL() (string, bool) {
    if e.coverImageURL == nil {
        return "", false
    }
    return *e.coverImageURL, true
}

func extractEmojiIcon(dto EntityResponseDto) *string {
    if dto.Icon == nil {
        return nil
    }
    if emoji, ok := dto.Icon.EmojiValue(); ok {
        return &emoji
    }
    return nil
}

func extractExternalIconURL(dto EntityResponseDto) *string {
    if dto.Icon == nil {
        return nil
    }
    if url, ok := dto.Icon.ExternalURL(); ok {
        return &url
    }
    return nil
}

func extractCoverImageURL(dto EntityResponseDto) *string {
    if dto.Cover == nil {
        return nil
    }
    switch dto.Cover.Type {
    case FileTypeExternal:
        if dto.Cover.External != nil {
            url := dto.Cover.External.URL
            return &url
        }
    case FileTypeFile:
        if dto.Cover.File != nil {
            url := dto.Cover.File.URL
            return &url
        }
    }
    return nil
}