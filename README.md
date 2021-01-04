# install

At the moment, there is no real install script provided. You need to have go installed (locally I have developed and tested it with version 1.14.3), so you can call

```zsh
go install
```

After the binary was produced, you can make an alias for your terminal, if you want to - I made an alias called "portscan" and following examples will use this alias.

## list commands

```zsh
portscan
```

## scan an ip

For example, use this to scan local IP, use 20 threads to scan and store the results in a file called portscan, so you can later grep the results.

```zsh
portscan scan -i 127.0.0.1 -t 20 > portscan
```

Or only interested in open?

```zsh
portscan scan -i 127.0.0.1 -t 20 | grep ^open
```


## flood an ip

Following command will try to establish 400 TCP connections at 127.0.0.1:8002 and uses the local client IP 127.0.4.1 for binding to the socket.

```zsh
portscan flood -i 127.0.0.1 -p 8002 -c 400 -l 127.0.4.1
```

Be careful what you do - use this in production might be illegal, when you are not allowed to do penetration testing at your own / customers network! I had developed this to can test firewalls or other blue teams strategies avoiding tcp connections flooding.

## ping an ip / port combination

Following command will try to ping the default message "Ping \r\n\r\n" to 127.0.0.1:3306 and print the response at stdout. (This example will show the used version of MySQL in output!)

```zsh
portscan ping -i 127.0.0.1 -p 3306
```