#!/bin/bash

# Create a directory for the website
mkdir -p ~/my_website

# Navigate to the directory
cd ~/my_website

# Create the HTML file
cat <<EOL > index.html
$1
EOL

echo "HTML file created at ~/my_website/index.html"
