const { Client } = require("discord.js");
const prueba = require("./src/Messages/messages")
require("dotenv").config();

const TOKEN = process.env.TOKEN
const client = new Client({intents: [131071]});

client.login(TOKEN)
  .then(() => {
    console.log(`Cliente ${client.user.username}`);
    client.user.setActivity("🥴")
  })
  .catch((err) => console.log(err));

prueba(client)