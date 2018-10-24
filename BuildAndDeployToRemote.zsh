echo "deployToRemote script started running\n"

rm GoFileServer

GOOS=linux GOARCH=amd64 go build .

ssh -t sreekar339@139.59.93.218 'sudo systemctl stop GoFileServer'

scp GoFileServer sreekar339@139.59.93.218:~/backend/go/fileServer/

ssh -t sreekar339@139.59.93.218 'sudo systemctl start GoFileServer'