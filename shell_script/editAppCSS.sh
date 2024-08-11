#!/bin/bash

# Navigate to the directory
cd ~/my-react-app/src

# Create the HTML file
cat <<EOL > App.css
$1
EOL

echo "App.css file updated at ~/my-react-app/App.css"
