#!/bin/bash

# Navigate to the directory
cd ~/my-react-app/src

# Create the HTML file
cat <<EOL > App.js
$1
EOL

echo "App.js file updated at ~/my-react-app/App.js"
