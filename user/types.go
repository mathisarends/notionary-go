package user

import (
	"encoding/json"
	"fmt"
)

type UserType string

const (
	UserTypePerson UserType = "person"
	UserTypeBot    UserType = "bot"
)

type WorkspaceOwnerType string

const (
	WorkspaceOwnerTypeUser      WorkspaceOwnerType = "user"
	WorkspaceOwnerTypeWorkspace WorkspaceOwnerType = "workspace"
)

type PersonUserDto struct {
	Email *string `json:"email,omitempty"`
}

type BotOwnerDto struct {
	Type      WorkspaceOwnerType `json:"type"`
	Workspace *bool              `json:"workspace,omitempty"`
}

type WorkspaceLimits struct {
	MaxFileUploadSizeInBytes int `json:"max_file_upload_size_in_bytes"`
}

type BotUserDto struct {
	Owner           *BotOwnerDto     `json:"owner,omitempty"`
	WorkspaceName   *string          `json:"workspace_name,omitempty"`
	WorkspaceLimits *WorkspaceLimits `json:"workspace_limits,omitempty"`
}

type NotionUserBase struct {
	Object    string   `json:"object"`
	ID        string   `json:"id"`
	Type      UserType `json:"type"`
	Name      *string  `json:"name,omitempty"`
	AvatarURL *string  `json:"avatar_url,omitempty"`
}

type PersonUserResponseDto struct {
	NotionUserBase
	Person PersonUserDto `json:"person"`
}

type BotUserResponseDto struct {
	NotionUserBase
	Bot BotUserDto `json:"bot"`
}

type UserResponseDto struct {
	Person *PersonUserResponseDto
	Bot    *BotUserResponseDto
}

func (u *UserResponseDto) UnmarshalJSON(data []byte) error {
	var base NotionUserBase
	if err := json.Unmarshal(data, &base); err != nil {
		return err
	}

	switch base.Type {
	case UserTypePerson:
		u.Person = &PersonUserResponseDto{}
		return json.Unmarshal(data, u.Person)
	case UserTypeBot:
		u.Bot = &BotUserResponseDto{}
		return json.Unmarshal(data, u.Bot)
	default:
		return fmt.Errorf("unknown user type: %s", base.Type)
	}
}

type NotionUsersListResponse struct {
	Results    []UserResponseDto `json:"results"`
	NextCursor *string           `json:"next_cursor,omitempty"`
	HasMore    bool              `json:"has_more"`
}

type PartialUserDto struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}

type Person struct {
	ID        string
	Name      *string
	AvatarURL *string
	Email     *string
}

func NewPerson(dto PersonUserResponseDto) Person {
	return Person{
		ID:        dto.ID,
		Name:      dto.Name,
		AvatarURL: dto.AvatarURL,
		Email:     dto.Person.Email,
	}
}

type Bot struct {
	ID            string
	Name          *string
	AvatarURL     *string
	WorkspaceName *string
	OwnerType     *WorkspaceOwnerType
}

func NewBot(dto BotUserResponseDto) Bot {
	var ownerType *WorkspaceOwnerType
	if dto.Bot.Owner != nil {
		ownerType = &dto.Bot.Owner.Type
	}
	return Bot{
		ID:            dto.ID,
		Name:          dto.Name,
		AvatarURL:     dto.AvatarURL,
		WorkspaceName: dto.Bot.WorkspaceName,
		OwnerType:     ownerType,
	}
}

type PartialUser struct {
	Object string `json:"object"` // always "user"
	ID     string `json:"id"`
}