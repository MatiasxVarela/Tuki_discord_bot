const { joinVoiceChannel, getVoiceConnection } = require('@discordjs/voice');

function prueba(client) {

  client.on("messageCreate", (message) => {
    if (message.content === ":milei") {
      message.channel.send({files: ["./src/assets/244081852_10158638416964891_8635652479027886497_n.jpg"]});
    }
  });
    
  client.on("messageCreate", (message) => {
    if (message.content === ":juliJugandoRisk") {
      message.channel.send({files: ["./src/assets/Juli.jpg"]});
    }
  });

  client.on("messageCreate", async (message) => {
    if (message.content.startsWith(":play")) {
      
    }
  });

  client.on("messageCreate", async (message) => {

    if (message.content.startsWith(":play")) {
      const voiceChannel = message.member.voice.channel;
      if (!voiceChannel) {
        return message.reply("Tenes que estar en un canal de voz");
      }

      joinVoiceChannel({
        channelId: voiceChannel.id,
        guildId: voiceChannel.guild.id,
        adapterCreator: voiceChannel.guild.voiceAdapterCreator,
      });
    }
  });

  client.on("messageCreate", async (message) => {
    if (message.content === ":stop") {
      const connection = getVoiceConnection(message.guild.id);
      if (connection) {
        connection.destroy();
      }
    }
  });

}

module.exports = prueba;
