
# Webhook Go
![GitHub Workflow Status](https://github.com/waxdred/webhook_go/actions/workflows/go.yml/badge.svg)
![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)

Webhook-Go is a Go library for sending webhooks. With this library, you can send messages, and embeds to a using a webhook URL.
## Table of Contents
- [Installation](#Installation)
- [Usage](#Usage)
- [Exemple](#Exemple)
- [Message](#Message)
- [Embed](#Embed)

# Installation
To use this library in your Go project, you can install it using go get:

```go
go get github.com/waxdred/webhook_go
import (
	wk "github.com/waxdred/webhook_go"
)
```
# Usage
- To use the Discord_webhook_go library, you first need to create a new Webhook object with the webhook URL:
- You can also send an embed:

## Exemple:
```go
package main

import (
   wk "github.com/waxdred/webhook_go"
)

func main() {
   // Create a new WebhookClient with the URL of your webhook
   client := wk.WebhookClient{}
   url := "webhook url"
   client.New(url)

   // Call the Send function with a closure that constructs the message and returns it
   _, err := client.Send(func(m *wk.Message) *wk.Message {
   // Set the content, username, and avatar URL of the message
      m.Content("Hello, world!").
         Username("My Bot").
     	 AvatarURL("https://i.imgur.com/4m34hi2.png").
	 // Add an embed to the message
         Embed(func(embed *wk.Embed) {
	 // Set the title, author, URL, color, and description of the embed
         embed.Title("My Embed").
	 Author("waxdred", "https://i.imgur.com/4m34hi2.png", "https://i.imgur.com/R66g1Pe.jpg").
	 Url("https://google.com/").
	 Color(15258703).
	 Desc("This is an example embed!").
	 // Add some fields to the embed
	 Field("Field 1", "This is the first field", false).
	 Field("Field 2", "This is the second field", false).
	 Field("Field 3", "This is the third field", false).
	 Field("Inline Field", "This is an inline field", true).
	 // Set the thumbnail and image of the embed
	 Thumbnail("https://i.imgur.com/4m34hi2.png").
	 Image("https://i.imgur.com/4m34hi2.png").
	 // Set the footer of the embed
	 Footer("Created by My Bot", "")
	 })
      return m
   })
   if err != nil {
      panic(err)
   }
}
```

## Message
- content_s (string, optional): the content of the message
- username_s (string, optional): the username that will be shown for the webhook message
- avatarURL_s (string, optional): the URL of the image that will be shown for the webhook message
- tts_s (bool, optional): whether the message should be sent as a text-to-speech message
- embeds_s ([]Embed, optional): an array of Embed objects to include in the message
- allowMentions_s (Allowed, optional): an object specifying which user mentions are allowed in the message
## Embed
- author_s (EmbedAuthor, optional): an object specifying the author of the embed
- title_s (string, optional): the title of the embed
- type_s (string, optional): the type of the embed
- description_s (string, optional): the description of the embed
- url_s (string, optional): the URL that the title should link to
- timestamp_s (string, optional): a timestamp to show at the bottom of the embed
- color_s (int64, optional): the color of the embed
- fields_s ([]EmbedField, optional): an array of EmbedField objects to include in the embed
- thumbnail_s (EmbedThumbnail, optional): an object specifying the thumbnail image of the embed
- image_s (EmbedImage, optional): an object specifying the main image of the embed
- video_s (EmbedVideo, optional): an object specifying a video to include in the embed
- footer_s (EmbedFooter, optional): an object specifying the footer of the embed
- provider_s (EmbedProvider, optional): an object specifying the provider of the embed

For more examples and detailed usage instructions, please see the GoDoc [documentation](https://pkg.go.dev/github.com/waxdred/webhook_go?utm_source=godoc).

Contributing
If you find a bug or would like to contribute to this library, please open an issue or pull request on the GitHub repository.

License
This library is licensed under the MIT License. See the LICENSE file for details.
