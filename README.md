## Webhook watcher service for automated deployments
A webhook watcher service for Boo's git repo. It is used for automated deployments

## **Prerequisites**
1. Make sure `go` is installed
2. Installation steps [here](https://go.dev/dl/)
3. Place the [deploy.sh](./deploy.sh) file under `/home/ifkash`
4. Make the [deploy.sh](./deploy.sh) executable by running `chmod +x deploy.sh`

## **Steps**
1. The current user for this is `ifkash`
2. Place the [webgook_server.go](./webhook_server.go) file under `/home/ifkash`
3. Build the executable
```sh
go build -o webhook_server webhook_server.go
```
4. Set up a systemd service to run the webhook server. Create a file `/etc/systemd/system/webhook-server.service`
```
[Unit]
Description=GitHub Webhook Server
After=network.target

[Service]
User=ifkash
WorkingDirectory=/home/ifkash
ExecStart=/home/ifkash/webhook_server
Restart=always

[Install]
WantedBy=multi-user.target
```
5. Reload the systemd daemon and restart the service
```sh
sudo systemctl daemon-reload
sudo systemctl restart webhook-server
```

## **Setup GitHub Webhook**
1. Go to repo settings, navigate to webhooks
2. Add/Edit a webhook
3. Set the payload URL to `http://GCP_VM_IP:5000/webhook`
4. Ensure the content type is set to `application/json`
5. Select "Just the push event" for the trigger

## **Monitor logs**
To monitor logs for the systemd service
```sh
sudo journalctl -u webhook-server -f
```