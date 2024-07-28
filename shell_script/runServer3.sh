#!/bin/bash

# Install necessary packages
sudo apt-get update
sudo apt-get install -y python3

# Create a directory for the server
mkdir -p ~/my_website
cd ~/my_website

# Create a simple HTTP server script
cat <<EOF > server.py
import http.server
import socketserver

PORT = 8000
Handler = http.server.SimpleHTTPRequestHandler

with socketserver.TCPServer(("", PORT), Handler) as httpd:
    print("Serving at port", PORT)
    httpd.serve_forever()
EOF

# Create a systemd service file for the daemon
sudo bash -c 'cat <<EOF > /etc/systemd/system/my_website.service
[Unit]
Description=Simple HTTP Server

[Service]
ExecStart=/usr/bin/python3 /home/$USER/my_website/server.py
Restart=always
User=$USER

[Install]
WantedBy=multi-user.target
EOF'

# Reload systemd to apply the new service
sudo systemctl daemon-reload

# Enable and start the service
sudo systemctl enable my_website.service
sudo systemctl start my_website.service

echo "Simple HTTP server is now running as a daemon on port 8000."
