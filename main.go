// Discord Bot用コード
package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {

	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Println("error:start\n", err)
	}

	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("error:start\n", err)
		return
	}

	//on message
	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error:wss\n", err)
		return
	}
	fmt.Println("BOT Running...")

	//シグナル受け取り可にしてチャネル受け取りを待つ（受け取ったら終了）
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	nick := m.Author.Username
	member, err := s.State.Member(m.GuildID, m.Author.ID)
	if err == nil && member.Nick != "" {
		nick = member.Nick
	}
	fmt.Println("< " + m.Content + " by " + nick)

	if m.Content == "日程調整" {
		s.ChannelMessageSend(m.ChannelID, "カレンダー表示")
		fmt.Println("> カレンダー表示")
	}
	if strings.Contains(m.Content, "日程") {
		s.ChannelMessageSend(m.ChannelID, "テストテスト")
		fmt.Println("> テストテスト")
	}
}
