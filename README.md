# clippy draw 

## what

It's a clippy bot that will draw clippy saying whatever you want

## how

Clone the repo
Set some environment variables
export API_KEY=<your_slack_api_key>
export BOT_USER=<the name you want it to use>
export KEY_WORD="clippy draw" # change this as you wish

go build
./clippy-bot 

You can also just use a container:

```
docker run -d -e API_KEY=<yourkey> \
    -e BOT_USER=<user> \
    -e KEY_WORD="clippy draw" \
    jeremyrickard/clippy-bot
``` 



