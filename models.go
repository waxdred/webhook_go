package webhook_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type (
	AllowedMention string
	Snowflake      string
)

const (
	RoleMention     AllowedMention = "roles"
	UserMention     AllowedMention = "users"
	EveryoneMention AllowedMention = "everyone"
)

// WebhookClient struct  
type WebhookClient struct {
	client http.Client
	url    string
}

// New method  
func (c *WebhookClient) New(url string) {
	c.client = http.Client{}
	c.url = url
}

// Send method  
func (c *WebhookClient) Send(function func(*Message) *Message) (*http.Response, error) {
	m := Message{}
	function(&m)
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	jsonBuffer := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("POST", c.url, jsonBuffer)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}

type Message struct {
	Content_s       *string  `json:"content,omitempty"`
	Username_s      *string  `json:"username,omitempty"`
	AvatarURL_s     *string  `json:"avatar_url,omitempty"`
	TTS_s           bool     `json:"tts"`
	Embeds_s        *[]Embed `json:"embeds,omitempty"`
	AllowMentions_s *Allowed `json:"allowed_mentions,omitempty"`
}

func (c *Message) Content(content interface{}) *Message {
	co := getString(content)
	c.Content_s = &co
	return c
}

func (c *Message) Username(username interface{}) *Message {
	u := getString(username)
	c.Username_s = &u
	return c
}

func (c *Message) AvatarURL(avatarURL interface{}) *Message {
	a := getString(avatarURL)
	c.AvatarURL_s = &a
	return c
}

func (c *Message) TTS(tts bool) *Message {
	c.TTS_s = tts
	return c
}

func (c *Message) AllowMentions(parse []AllowedMention, roles []Snowflake, users []Snowflake, repliedUser bool) *Message {
	c.AllowMentions_s = &Allowed{
		Parse:       parse,
		Roles:       roles,
		Users:       users,
		RepliedUser: repliedUser,
	}
	return c
}

func (c *Message) Embed(embedFunc func(*Embed)) *Message {
	if c.Embeds_s == nil {
		c.Embeds_s = &[]Embed{}
	}
	embed := Embed{}
	embedFunc(&embed)
	*c.Embeds_s = append(*c.Embeds_s, embed)
	return c
}

// Embed struct  
type Embed struct {
	Author_s      *EmbedAuthor    `json:"author,omitempty"`
	Title_s       *string         `json:"title,omitempty"`
	Type_s        *string         `json:"type,omitempty"`
	Description_s *string         `json:"description,omitempty"`
	URL_s         *string         `json:"url,omitempty"`
	Timestamp_s   *string         `json:"timestamp,omitempty"`
	Color_s       *int64          `json:"color,omitempty"`
	Fields_s      *[]EmbedField   `json:"fields,omitempty"`
	Thumbnail_s   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Image_s       *EmbedImage     `json:"image,omitempty"`
	Video_s       *EmbedVideo     `json:"video,omitempty"`
	Footer_s      *EmbedFooter    `json:"footer,omitempty"`
	Provider_s    *EmbedProvider  `json:"provider,omitempty"`
}

func (e *Embed) Title(title interface{}) *Embed {
	t := getString(title)
	e.Title_s = &t
	return e
}

func (e *Embed) Type(typestring interface{}) *Embed {
	t := getString(typestring)
	e.Type_s = &t
	return e
}

func (e *Embed) Desc(desc interface{}) *Embed {
	d := getString(desc)
	e.Description_s = &d
	return e
}

func (e *Embed) Url(url interface{}) *Embed {
	u := getString(url)
	e.URL_s = &u
	return e
}

func (e *Embed) Timestamp(timestamp interface{}) *Embed {
	t := getString(timestamp)
	e.Timestamp_s = &t
	return e
}

// AddColor method  
// color need be in decimal
func (e *Embed) Color(color interface{}) *Embed {
	c := getInt(color)
	e.Color_s = &c
	return e
}

func (e *Embed) Footer(text, icon_url interface{}) *Embed {
	e.Footer_s = &EmbedFooter{
		Text:     getString(text),
		Icon_url: getString(icon_url),
	}
	return e
}

func (e *Embed) Image(url interface{}) *Embed {
	e.Image_s = &EmbedImage{
		Url: getString(url),
	}
	return e
}

func (e *Embed) Video(url interface{}) *Embed {
	e.Video_s = &EmbedVideo{
		Url: getString(url),
	}
	return e
}

func (e *Embed) Thumbnail(url interface{}) *Embed {
	e.Thumbnail_s = &EmbedThumbnail{
		Url: getString(url),
	}
	return e
}

func (e *Embed) Provider(name, url interface{}) *Embed {
	e.Provider_s = &EmbedProvider{
		Name: getString(name),
		Url:  getString(url),
	}
	return e
}

func (e *Embed) Author(name, url, icon_url interface{}) *Embed {
	e.Author_s = &EmbedAuthor{
		Name:     getString(name),
		Url:      getString(url),
		Icon_url: getString(icon_url),
	}
	return e
}

func (e *Embed) Field(name interface{}, value interface{}, inline bool) *Embed {
	if e.Fields_s == nil {
		e.Fields_s = &[]EmbedField{}
	}
	if len(*e.Fields_s) > 25 {
		panic("You can't have more than 25 fields in an embed!")
	}
	fields := EmbedField{
		Name:   getString(name),
		Value:  getString(value),
		Inline: inline,
	}
	*e.Fields_s = append(*e.Fields_s, fields)
	return e
}

func getString(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	case int:
		return strconv.FormatInt(int64(value.(int)), 10)
	case int8:
		return strconv.FormatInt(int64(value.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(value.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 10)
	case int64:
		return strconv.FormatInt(value.(int64), 10)
	case uint:
		return strconv.FormatUint(uint64(value.(uint)), 10)
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(value.(uint64), 10)
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	default:
		return ""
	}
}

func getInt(value interface{}) int64 {
	switch value.(type) {
	case string:
		ret, _ := strconv.ParseInt(value.(string), 10, 64)
		return ret
	case int:
		return value.(int64)
	case int8:
		return value.(int64)
	case int16:
		return value.(int64)
	case int32:
		return value.(int64)
	case int64:
		return value.(int64)
	default:
		return 0
	}
}

type EmbedFooter struct {
	Text     string `json:"text,omitempty"`
	Icon_url string `json:"icon_url,omitempty"`
}

type EmbedImage struct {
	Url string `json:"url,omitempty"`
}

type EmbedVideo struct {
	Url string `json:"url,omitempty"`
}

type EmbedThumbnail struct {
	Url string `json:"url,omitempty"`
}
type EmbedUrlSource struct {
	Url string `json:"url,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name     string `json:"name,omitempty"`
	Url      string `json:"url,omitempty"`
	Icon_url string `json:"icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type Allowed struct {
	Parse       []AllowedMention `json:"parse,omitempty"`
	Roles       []Snowflake      `json:"roles,omitempty"`
	Users       []Snowflake      `json:"users,omitempty"`
	RepliedUser bool             `json:"replied_user"`
}
