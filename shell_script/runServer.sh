# Navigate to the website directory
cd ~/my_website

# Run the HTTP server on port 8000
python3 -m http.server 8000 --bind 0.0.0.0

echo "Server is running on http://<your-public-ip>:8000"
