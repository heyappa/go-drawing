https://outcrawl.com/realtime-collaborative-drawing-go/

// first install the following packages
go get -u github.com/gorilla/websocket \
  github.com/lucasb-eyer/go-colorful \
    github.com/tidwall/gjson \
      github.com/satori/go.uuid


// to run
"go build -o server && ./server"  OR  "go run *.go"
and open client/index.html in browser

// (2018/06/05) localhost:3000 added

