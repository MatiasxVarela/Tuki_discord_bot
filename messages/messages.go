package messages

import (
	"bot/discord/services"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, ":milei") {
		file, _ := os.Open("./assets/244081852_10158638416964891_8635652479027886497_n.jpg")
		defer file.Close()
		s.ChannelFileSend(m.ChannelID, "milei.jpg", file)
	}

	if strings.HasPrefix(m.Content, ":juliJugandoRisk") {
		file, _ := os.Open("./assets/Juli.jpg")
		defer file.Close()
		s.ChannelFileSend(m.ChannelID, "julito.jpg", file)
	}

	if strings.HasPrefix(m.Content, ":cat") {

		catPic, err := services.GetRandomCat()
		if err != nil {
			s.ChannelMessage(m.ChannelID, "Error al obtener la imagen de gato")
			return
		}
		defer catPic.Close()

		s.ChannelFileSend(m.ChannelID, "gato.jpg", catPic)
	}

	if strings.HasPrefix(m.Content, ":join") {

		voiceState, err := s.State.VoiceState(m.GuildID, m.Author.ID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error al conectar al chat de voz")
		}

		voiceChannel, err := s.Channel(voiceState.ChannelID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error al conectar al chat de voz")
		}

		_, err = s.ChannelVoiceJoin(m.GuildID, voiceChannel.ID, false, true)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error al conectar al chat de voz")
			return
		}

	}
}
