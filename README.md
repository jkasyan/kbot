# kbot

Telegram bot - https://t.me/ekasyan_bot
Currently kbot supports 2 commands help and hello.

## Usage
1. Clone the project https://github.com/JKasyan/kbot.git
2. Run 
```bash 
go build -ldflags "-X="github.com/JKasyan/cmd.appVersion=v1.0.0
```
3. Set environment variable from telegram
```bash export TELE_TOKEN={{your token}}```

4. Run ```bash ./kbot```