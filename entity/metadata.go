package entity

import "context"

type MetadataClient interface {
    PatchEmojiIcon(ctx context.Context, emoji string) (*EntityResponseDto, error)
    PatchExternalIcon(ctx context.Context, url string) (*EntityResponseDto, error)
    PatchIconFromFileUpload(ctx context.Context, fileUploadID string) (*EntityResponseDto, error)
    RemoveIcon(ctx context.Context) error
    PatchExternalCover(ctx context.Context, url string) (*EntityResponseDto, error)
    PatchCoverFromFileUpload(ctx context.Context, fileUploadID string) (*EntityResponseDto, error)
    RemoveCover(ctx context.Context) error
    MoveToTrash(ctx context.Context) (*EntityResponseDto, error)
    RestoreFromTrash(ctx context.Context) (*EntityResponseDto, error)
}