// Discord Bot用コード
package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	vcsession *discordgo.VoiceConnection
)

// http://localhost:8080/ へアクセスしたときの処理
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	//Discordのセッションを作成
	discord, err := discordgo.New()
	discord.Token = os.Getenv("TOKEN")
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
	}

	discord.AddHandler(onMessageCreate) //全てのWSAPIイベントが発生した時のイベントハンドラを追加
	// websocketを開いてlistening開始
	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer discord.Close()

	fmt.Println("Listening...")
	<-make(chan bool) //プログラムが終了しないようロック
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	err, _ := discordgo.New()
	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}
	fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)

	switch {
	case strings.HasPrefix(m.Content, fmt.Sprintf("%s %s", os.Getenv("APPLICATION_ID"), "通話開始")): //Bot宛に通話開始コマンドが実行された時
		sendMessage(s, m.ChannelID, "Hello world！")

	case strings.HasPrefix(m.Content, fmt.Sprintf("%s %s", os.Getenv("APPLICATION_ID"), "退出")):
		vcsession.Disconnect() //今いる通話チャンネルから抜ける
	}
}

//メッセージを送信する関数
func sendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)

	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
