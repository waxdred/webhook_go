# Package Discord_webhook_go

```
import "github.com/waxdred/Discord_webhook_go"
```
Package Discord_webhook_go provides functionality for sending messages to a Discord channel using a webhook.

## Types
- type Allowed
```go
type Allowed struct {
	Parse []string `json:"parse,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Users []string `json:"users,omitempty"`
}
```
Allowed represents a configuration for allowed mentions in a Discord message.

- type Embed
```go
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
```
Embed represents a message embed in a Discord message.

- type EmbedAuthor
```go
type EmbedAuthor struct {
	Name_s    string  `json:"name,omitempty"`
	URL_s     *string `json:"url,omitempty"`
	IconURL_s *string `json:"icon_url,omitempty"`
}
```
EmbedAuthor represents an author in a Discord message embed.

- type EmbedField
```go
type EmbedField struct {
	Name_s   string `json:"name"`
	Value_s  string `json:"value"`
	Inline_s bool   `json:"inline,omitempty"`
}
```
EmbedField represents a field in a Discord message embed.

- type EmbedFooter
```go
type EmbedFooter struct {
	Text_s    string  `json:"text"`
	IconURL_s *string `json:"icon_url,omitempty"`
}
```
EmbedFooter represents a footer in a Discord message embed.

- type EmbedImage
```go
type EmbedImage struct {
	URL_s string `json:"url"`
}
```
EmbedImage represents an image in a Discord message embed.

- type EmbedProvider
```go
type EmbedProvider struct {
	Name_s *string `json:"name,omitempty"`
	URL_s  *string `json:"url,omitempty"`
}
```
EmbedProvider represents a provider in a Discord message embed.

- type EmbedThumbnail
```go
type EmbedThumbnail struct {
	URL_s string `json:"url"`
}
```
EmbedThumbnail represents a thumbnail in a Discord message embed.

- type EmbedVideo
```go
type EmbedVideo struct {
	URL_s string `json:"url"`
}
```
EmbedVideo represents a video in a Discord message embed.

- type Message
```go
type Message struct {
	Content_s       *string  `json:"content,omitempty"`
	Username_s      *string  `json:"username,omitempty"`
	AvatarURL_s     *string  `json:"avatar_url,omitempty"`
	TTS_s           bool     `json:"tts"`
	Embeds_s        *[]Embed `json:"embeds,omitempty"`
	AllowMentions_s *Allowed `json:"allowed_mentions,omitempty"`
}
```
Message represents a message to be sent to a Discord channel.

- type WebhookClient
```go
type WebhookClient struct {
    client http.Client
	url    string
}
```
WebhookClient represents a client for sending messages to a Discord webhook.

## Functions
func (*WebhookClient) New
```go
func (c *WebhookClient) New(url string)
New sets up the client with the given Discord webhook URL.
```

func (*WebhookClient) Send
```go
func (c *WebhookClient) Send(f func(m *Message) *Message) (*http.Response, error)
```
Send sends a message to the configured webhook. It takes a function which is used to configure the message. The function should take a pointer to a Message struct and return the same pointer. The configured message is returned as the result of the function. The function can be used to set various options for the message, such as content, username, avatar URL, TTS, and embeds. Example usage:

```go
client := WebhookClient{}
client.New("https://discord.com/api/webhooks/...")
_, err := client.Send(func(m *Message) *Message {
    m.Content("Hello, world!").
        Username("my-bot").
        AvatarURL("https://i.imgur.com/abc123.jpg").
        TTS(false).
        Embed(func(embed *Embed) {
            embed.Title("My Title").
                Description("This is my description.").
                Color(0xFF0000).
                Field("My Field", "My Field Value", true)
        })
    return m
})
```
The Send function returns an http.Response struct and an error. If the request to the Discord API is successful, Send will return an HTTP response with a 204 status code and a nil error. If there is an error, Send will return an error that describes the problem.

## Types
- type Allowed
```go
type Allowed struct {
	Parse_s   []string `json:"parse,omitempty"`
	Roles_s   []string `json:"roles,omitempty"`
	Users_s   []string `json:"users,omitempty"`
	Replied_s bool     `json:"replied_user,omitempty"`
}
```
Allowed is a struct for setting up message mention permissions.

- type Embed
```go
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
```
Embed is a struct for setting up an embedded message.

- type EmbedAuthor
```go
type EmbedAuthor struct {
	Name_s    *string `json:"name,omitempty"`
	URL_s     *string `json:"url,omitempty"`
	IconURL_s *string `json:"icon_url,omitempty"`
}
```
EmbedAuthor is a struct for setting up an embedded message author.

- type EmbedField
```go
type EmbedField struct {
	Name_s   *string `json:"name,omitempty"`
	Value_s  *string `json:"value,omitempty"`
	Inline_s bool    `json:"inline,omitempty"`
}
```
EmbedField is a struct for setting up an embedded message field.
