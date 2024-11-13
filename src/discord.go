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

			sha256, err := CompilerWrite(code, language)
			if err != nil {
				fmt.Println(err)
				return
			}

			var compileMess = "."

			message, err := s.ChannelMessageSend(m.ChannelID, "```"+"Processing."+"```")
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
						s.ChannelMessageEdit(message.Reference().ChannelID, message.Reference().MessageID, "```"+"Processing"+compileMess+"```")
					}
				}
			}()

			output, err := CreateVM(sha256, LanguageTypes[language], language)
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

			if len(output) == 0 {
				output = "null"
				s.ChannelMessageEdit(message.Reference().ChannelID, message.Reference().MessageID, fmt.Sprintf("```%s```", output))
				return
			}

			_, err = s.ChannelMessageEdit(message.Reference().ChannelID, message.Reference().MessageID, fmt.Sprintf("```%s```", output[:300]))
			if err != nil {
				fmt.Println(err)
				s.ChannelMessageEdit(message.Reference().ChannelID, message.Reference().MessageID, "```To many Requests```")
				return
			}
		}
	}
}

func CompilerWrite(value, dot string) (string, error) {
	sha256 := sha256.Sum256([]byte(fmt.Sprintf("%d", rand.IntN(1000))))
	file, err := os.Create(fmt.Sprintf("./scripts/%x.%s", sha256, dot))
	if err != nil {
		return "", err
	}
	file.Write([]byte(value))
	file.Close()

	return fmt.Sprintf("%x", sha256), nil
}
