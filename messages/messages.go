package messages

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	CAT_API := os.Getenv("CAT_API")

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

		res, err := http.Get("https://api.thecatapi.com/v1/images/search?api_key=" + CAT_API)
		if err != nil {
			s.ChannelMessage(m.ChannelID, "Error en la petición")
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			s.ChannelMessage(m.ChannelID, "Error en la petición")
		}

		var data []map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			s.ChannelMessage(m.ChannelID, "Error en la petición")
		}

		url := data[0]["url"].(string)
		res, err = http.Get(url)
		if err != nil {
			s.ChannelMessage(m.ChannelID, "Error en la petición")
		}

		defer res.Body.Close()

		s.ChannelFileSend(m.ChannelID, "gato.jpg", res.Body)

	}

}
