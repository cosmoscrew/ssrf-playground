# SSRF Playground

SSRF Playground is a platform to practice exploitation and testing for Server Side Request Forgery issues in web applications.

### Setup

Because this using logrus package, you need to get logrus first:
```bash
export GOPATH=~/go
export PATH=$PATH:$GOPATH/bin #you can save this to .bashrc or .zshrc
mkdir -p $GOPATH
cd $GOPATH
go get github.com/sirupsen/logrus
```

### How to play?

Run the server by typing `go run .`

There are different challenges with varying level of difficulty. The end goal is accessing a service running at localhost to get the flag.

```bash
http://localhost:8082/flag
```

The flag is randomly generated with a hash of timestamp at the end.

Have fun playing. The UI is accessible at http://localhost:8001/

### Contributing

The project is still in very early stage hence contributions are more than welcomed.
