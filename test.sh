go build -o godaddy main.go

./godaddy domain set-record -d 8lr.uk -n dev -t A -v 11
