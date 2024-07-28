On linux only, run:

```
chmod +x script.sh
sudo apt-get install npm
```

Then
```
sudo apt-get install npm
npm install -g http-server
npm init -y
npm install http-server --save-dev
```

package.json
```
{
  "name": "my-simple-server",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "start": "http-server -a 0.0.0.0 -p 8000"
  },
  "author": "",
  "license": "ISC",
  "devDependencies": {
    "http-server": "^14.0.0"
  }
}
```

Create in home directory my_server folder, then 

Then

```
go run main.go
```