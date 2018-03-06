# Deploy clippy from Terraform

## what

One click clippy on Azure.

## how

Clone the repo
Set image and environment variables in the main.tf file.

image=<your_image> (currently a snapshot public build from March)
API_KEY=<your_slack_api_key>
BOT_USER=<the name you want it to use>
KEY_WORD="clippy draw" # change this as you wish

Try it from Bash in Cloud Shell which has Terraform pre-installed and authenticated:

[![Launch Cloud Shell](https://shell.azure.com/images/launchcloudshell.png "Launch Cloud Shell")](https://shell.azure.com/bash)

1. git clone https://github.com/jeremyrickard/clippy-bot.git
2. cd terraform
3. update image, api key, bot user, and key word in main.tf
4. terraform init
5. terraform plan
6. terraform apply
7. terraform destroy