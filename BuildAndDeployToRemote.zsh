#!/bash/zsh
echo "deployToRemote script started running\n"

# rm lol.lol \
# \
GOOS=linux GOARCH=amd64 go build . \
&& echo "0" \
&& scp GoFileServer.service sreekar339@139.59.93.218:~/  \
&& echo "1" \
&& ssh -t sreekar339@139.59.93.218 'sudo systemctl stop GoFileServer' \
&& echo "2" \
&& scp GoFileServer ServerConfig.yml sreekar339@139.59.93.218:~/backend/go/fileServer \
&& echo "3" \
&& ssh -t sreekar339@139.59.93.218 '
echo "4";
sudo rm /etc/systemd/system/GoFileServer.service;
echo "5";
sudo mv ~/GoFileServer.service /etc/systemd/system/GoFileServer.service;
echo "6";
sudo systemctl daemon-reload &&
echo "7" &&
sudo systemctl enable GoFileServer &&
echo "8" &&
sudo systemctl start GoFileServer &&
echo "9"
' \
&& echo "10"