package src

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/rand/v2"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MessageContent(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!compiler") {
		value := strings.TrimPrefix(m.Content, "!compiler")

		split := Split(value)

		if strings.HasPrefix(split, "```") {
			lang := SplitPrefix(value)

			split = Split(lang)

			code, language := LanguageSele(split, lang)

			if language == "" || code == "" {
				return
			}

			sha256, err := CompilerWrite(code, language)
			if err != nil {
				fmt.Println(err)
				return
			}

			var compileMess = "."

			embed := &discordgo.MessageEmbed{
				Color:       0x00ff00,
				Description: language,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "console",
						Value:  "```" + "Processing." + "```",
						Inline: true,
					},
				},
			}

			message, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
			if err != nil {
				fmt.Println(err)
				return
			}

			ctx, cancel := context.WithCancel(context.Background())

			var is = false

			go func() {
				for {
					select {
					case <-ctx.Done():
						is = true
						return
					default:
						if len(compileMess) == 6 {
							compileMess = ""
						}

						time.Sleep(1 * time.Second)
						compileMess += "."
						embed.Fields[0].Value = "```" + "Processing" + compileMess + "```"

						s.ChannelMessageEditEmbed(message.Reference().ChannelID, message.Reference().MessageID, embed)
					}
				}
			}()

			output, exitcode, err := CreateVM(sha256, LanguageTypes[language], language)
			if err != nil {
				fmt.Println(err)
				cancel()
				return
			}

			cancel()

			err = os.Remove(fmt.Sprintf("./scripts/%s.%s", sha256, language))
			if err != nil {
				fmt.Println(err)
				return
			}

			for !is {
				time.Sleep(300 * time.Millisecond)
				continue
			}

			output = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`).ReplaceAllString(output, "")
			output = strings.ReplaceAll(output, "`", "\\`")

			embed = &discordgo.MessageEmbed{
				Color:       StatusColor(exitcode),
				Description: language,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "console",
						Value:  fmt.Sprintf("```%s```", output),
						Inline: false,
					},
					{
						Name:   "hash",
						Value:  sha256,
						Inline: false,
					},
				},
			}

			if len(output) == 0 {
				output = "null"
				embed.Fields[0].Value = "```\n" + output + "```"
				s.ChannelMessageEditEmbed(message.Reference().ChannelID, message.Reference().MessageID, embed)
				return
			}

			split := strings.Split(output, "\n")

			if len(split) >= 30 {
				output = ""
				for _, s := range split[:30] {
					output += s + "\n"
				}
			}

			embed.Fields[0].Value = "```\n" + output + "```"

			_, err = s.ChannelMessageEditEmbed(message.Reference().ChannelID, message.Reference().MessageID, embed)
			if err != nil {
				fmt.Println(err)
				embed.Fields[0].Value = "```To many Requests```"
				s.ChannelMessageEditEmbed(message.Reference().ChannelID, message.Reference().MessageID, embed)
				return
			}
		}
	}
}

func StatusColor(color int) int {
	if color != 0 {
		return 0xff0000
	} else {
		return 0x00ff00
	}
}

func CompilerWrite(value, dot string) (string, error) {
	sha256 := sha256.Sum256([]byte(fmt.Sprintf("%d", rand.IntN(1000))))
	code := fmt.Sprintf("%x", sha256)[:20]
	file, err := os.Create(fmt.Sprintf("./scripts/%s.%s", code, dot))
	if err != nil {
		return "", err
	}
	file.Write([]byte(value))
	file.Close()

	return code, nil
}
