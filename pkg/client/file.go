package client

import (
	"fmt"
	"io"
	"net/url"
)

type InputFile interface {
	inputFileTag()
}

type InputFileType int

type InputFileUpload struct {
	Filename string
	Data     io.Reader
}

func (InputFileUpload) inputFileTag() {}

func (i *InputFileUpload) MarshalJSON() ([]byte, error) {
	return []byte(`"@` + i.Filename + `"`), nil
}

type InputFileString struct {
	Data string
}

func (InputFileString) inputFileTag() {}

func (i *InputFileString) MarshalJSON() ([]byte, error) {
	return []byte(`"` + i.Data + `"`), nil
}

type SendPhotoParams struct {
	ChatID                   int                  `json:"chat_id"`
	MessageThreadID          int                  `json:"message_thread_id,omitempty"`
	Photo                    InputFile            `json:"photo"`
	Caption                  string               `json:"caption,omitempty"`
	ParseMode                string               `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity      `json:"caption_entities,omitempty"`
	HasSpoiler               bool                 `json:"has_spoiler,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageID         int                  `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (s SendPhotoParams) getParams() map[string]string {
	var m = make(map[string]string)

	m["chat_id"] = fmt.Sprintf("%d", s.ChatID)
	m["photo"] = jsonAnything(s.Photo)
	return m
}

func (s SendPhotoParams) getEndpoint() string {
	return "sendPhoto"
}

func NewPhotoUpload(chatID int64, file any) PhotoConfig {
	return PhotoConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			File:        file,
			UseExisting: false,
		},
	}
}

// PhotoConfig contains information about a SendPhoto request.
type PhotoConfig struct {
	BaseFile
	Thumb           RequestFileData
	Caption         string
	ParseMode       string
	CaptionEntities []MessageEntity
}

// Params represents a set of parameters that gets passed to a request.
type Params map[string]string

//// AddNonEmpty adds a value if it not an empty string.
//func (p Params) AddNonEmpty(key, value string) {
//	if value != "" {
//		p[key] = value
//	}
//}
//
//// AddInterface adds an interface if it is not nil and can be JSON marshalled.
//func (p Params) AddInterface(key string, value interface{}) error {
//	if value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) {
//		return nil
//	}
//
//	b, err := json.Marshal(value)
//	if err != nil {
//		return err
//	}
//
//	p[key] = string(b)
//
//	return nil
//}

func (config PhotoConfig) getParams() map[string]string {
	params, _ := config.BaseFile.params()

	if config.Caption != "" {
		params["caption"] = config.Caption
		if config.ParseMode != "" {
			params["parse_mode"] = config.ParseMode
		}
	}

	return params
}

// Values returns a url.Values representation of PhotoConfig.
func (config PhotoConfig) values() (url.Values, error) {
	v, err := config.BaseChat.values()
	if err != nil {
		return v, err
	}

	v.Add(config.name(), config.FileID)
	if config.Caption != "" {
		v.Add("caption", config.Caption)
		if config.ParseMode != "" {
			v.Add("parse_mode", config.ParseMode)
		}
	}

	return v, nil
}

// name returns the field name for the Photo.
func (config PhotoConfig) name() string {
	return "photo"
}

// method returns Telegram API method name for sending Photo.
func (config PhotoConfig) method() string {
	return "sendPhoto"
}

func (p PhotoConfig) getEndpoint() string {
	return "sendPhoto"
}

// BaseFile is a base type for all file config types.
type BaseFile struct {
	BaseChat
	File        interface{}
	FileID      string
	UseExisting bool
	MimeType    string
	FileSize    int
}

func (file BaseFile) params() (Params, error) {
	return file.BaseChat.params()
}

func (chat *BaseChat) params() (Params, error) {
	params := make(Params)

	params["chat_id"] = fmt.Sprint(chat.ChatID)

	params["reply_markup"] = jsonAnything(chat.ReplyMarkup)

	return params, nil
}

// RequestFileData represents the data to be used for a file.
type RequestFileData interface {
	// NeedsUpload shows if the file needs to be uploaded.
	NeedsUpload() bool

	// UploadData gets the file name and an `io.Reader` for the file to be uploaded. This
	// must only be called when the file needs to be uploaded.
	UploadData() (string, io.Reader, error)
	// SendData gets the file data to send when a file does not need to be uploaded. This
	// must only be called when the file does not need to be uploaded.
	SendData() string
}

// FileID is an ID of a file already uploaded to Telegram.
type FileID string

func (fi FileID) NeedsUpload() bool {
	return false
}

func (fi FileID) UploadData() (string, io.Reader, error) {
	panic("FileID cannot be uploaded")
}

func (fi FileID) SendData() string {
	return string(fi)
}

// EditMessageMediaConfig allows you to make an editMessageMedia request.
type EditMessageMediaConfig struct {
	BaseEdit

	Media interface{}
}

func (EditMessageMediaConfig) method() string {
	return "editMessageMedia"
}

// InputMediaPhoto is a photo to send as part of a media group.
type InputMediaPhoto struct {
	BaseInputMedia
}

// BaseInputMedia is a base type for the InputMedia types.
type BaseInputMedia struct {
	Type            string          `json:"type"`
	Media           RequestFileData `json:"media"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
}

func NewInputMediaPhoto(media RequestFileData) InputMediaPhoto {
	return InputMediaPhoto{
		BaseInputMedia{
			Type:  "photo",
			Media: media,
		},
	}
}

// fileAttach is an internal file type used for processed media groups.
type fileAttach string

func (fa fileAttach) NeedsUpload() bool {
	return false
}

func (fa fileAttach) UploadData() (string, io.Reader, error) {
	panic("fileAttach cannot be uploaded")
}

func (fa fileAttach) SendData() string {
	return string(fa)
}

func (config EditMessageMediaConfig) params() (map[string]string, error) {
	params, err := config.BaseEdit.params()
	if err != nil {
		return params, err
	}
	md := (config.Media).(InputMediaPhoto)
	if md.Media.NeedsUpload() {
		md.Media = fileAttach(fmt.Sprintf("attach://file-%d", 0))
	}
	params["media"] = jsonAnything(md)

	return params, err
}

func (config EditMessageMediaConfig) files() []RequestFile {
	files := make([]RequestFile, 0, 10)
	md := (config.Media).(InputMediaPhoto)
	if md.Media.NeedsUpload() {
		files = append(files, RequestFile{
			Name: fmt.Sprintf("file-%d", 0),
			Data: md.Media,
		})
	}

	return files
}

// RequestFile represents a file associated with a field name.
type RequestFile struct {
	// The file field name.
	Name string
	// The file data to include.
	Data RequestFileData
}
