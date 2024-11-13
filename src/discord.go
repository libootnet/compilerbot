package src

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/rand/v2"
	"os"
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

			message, err := s.ChannelMessageSend(m.ChannelID, "```"+"compiling."+"```")
			if err != nil {
				fmt.Println(err)
				return
			}

			ctx, cancel := context.WithCancel(context.Background())

			go func() {
				for {
					select {
					case <-ctx.Done():
						return
					default:
						if len(compileMess) == 5 {
							compileMess = ""
						}
						time.Sleep(2 * time.Second)
						compileMess += "."
						s.ChannelMessageEdit(message.Reference().ChannelID, message.Reference().MessageID, "```"+"compiling"+compileMess+"```")
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

			s.ChannelMessageEdit(message.Reference().ChannelID, message.Reference().MessageID, fmt.Sprintf("```%s```", output))
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
