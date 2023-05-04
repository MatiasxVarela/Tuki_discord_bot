const { joinVoiceChannel, getVoiceConnection, createAudioPlayer, createAudioResource, SodiumEncrypter } = require('@discordjs/voice');
const ytdl = require("ytdl-core-discord");
const _sodium = require('libsodium-wrappers');

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

  client.on('messageCreate', async (message) => {
    if (message.content.startsWith(':play')) {
      const voiceChannel = message.member.voice.channel;
      if (!voiceChannel) {
        return message.reply('Tenes que estar en un canal de voz');
      }

      const connection = joinVoiceChannel({
        channelId: voiceChannel.id,
        guildId: voiceChannel.guild.id,
        adapterCreator: voiceChannel.guild.voiceAdapterCreator,
        encrypterFactory: () => new SodiumEncrypter(_sodium)
      });
  
      const player = createAudioPlayer();
      connection.subscribe(player);
  
      const args = message.content.split(' ');
      const url = args[1];
  
      try {
        const stream = await ytdl(url, { filter: 'audioonly' });
        const resource = createAudioResource(stream);
        player.play(resource);
        message.reply(`Reproduciendo ${url}`);
      } catch (error) {
        console.error(error);
        message.reply('No se pudo reproducir la canción');
      }
    }
  });

  client.on("messageCreate", async (message) => {
    if (message.content === ":stop") {
      const connection = getVoiceConnection(message.member.voice.channel);
      console.log(connection)
      if (connection) {
        connection.destroy();
      }
    }
  });

}

module.exports = prueba;
