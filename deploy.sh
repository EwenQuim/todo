#!/bin/bash
# Description: Deploy the app to the VPS


SSH_KEY_PATH=~/.ssh/id_rsa
TARGET_HOST=vps

# exit when any command fails
set -e

echo "===> Updating remote server dependencies"
ssh -i $SSH_KEY_PATH $TARGET_HOST 'sudo apt-get update && sudo apt-get upgrade -y && sudo apt-get autoremove -y'

echo "===> Generating the binary"
go generate ./...

echo "===> Building the binary"
GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o todo-app

echo "===> Zipping the binary"
zip todoapp.zip todo-app

echo "===> Copying the binary into server with temporary location so the downtime is minimal"
scp -i $SSH_KEY_PATH todoapp.zip $TARGET_HOST:/home/ubuntu/todo-tmp.zip

echo "===> Unzipping the binary"
ssh -i $SSH_KEY_PATH $TARGET_HOST 'unzip -o /home/ubuntu/todo-tmp.zip -d /home/ubuntu/'

# echo "===> Move the migration files"
# scp -i $SSH_KEY_PATH -r db/ $TARGET_HOST:/home/ubuntu

# echo "===> Copying the service file"
# scp -i $SSH_KEY_PATH ./scripts/todo.service $TARGET_HOST:/tmp/todo.service

# echo "===> Moving the service file at the right place"
# ssh -i $SSH_KEY_PATH $TARGET_HOST 'sudo mv /tmp/todo.service /etc/systemd/system/todo.service'

echo "===> Reloading the daemon"
ssh -i $SSH_KEY_PATH $TARGET_HOST 'sudo systemctl daemon-reload'

echo "===> Stopping the service"
ssh -i $SSH_KEY_PATH $TARGET_HOST 'sudo systemctl stop todo'

echo "===> Moving the binary at the right place (overwriting the old one quickly)"
ssh -i $SSH_KEY_PATH $TARGET_HOST 'sudo mv /home/ubuntu/todo-app /home/ubuntu/todo/'

echo "===> Starting the service"
echo "If it fails, you can check the logs with: journalctl -u todo -f. Possible errors are:"
echo "- the binary is not executable, in that case, you can fix it with: sudo chmod +x /home/ubuntu/todo-back"
echo "- the .env file is not present"
ssh -i $SSH_KEY_PATH $TARGET_HOST 'sudo systemctl start todo'

echo "===> Cleaning up"
rm todo-app
rm todoapp.zip
ssh -i $SSH_KEY_PATH $TARGET_HOST 'rm /home/ubuntu/todo-tmp.zip'

echo "===> Checking the status of the service"
ssh -i $SSH_KEY_PATH $TARGET_HOST 'sudo systemctl status todo'

echo "===> Done"
