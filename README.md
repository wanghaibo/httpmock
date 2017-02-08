# run
```sh
usage of ./mock:
--adminport string    admin port (default "8089")
--datapath string     path to store mockinfo
--serverport string   server port (default "80")
```

```sh
docker run -ti -p 80:80 -p 8089:8089 wanghaibo/httpmock
curl '127.0.0.1:8089/mocks/' -d '{"url":"http://www.baidu.com/","body":"test","headers":{"a":"b", "c":"d", "Content-Type":"Text/Html2", "Wanghaibo":"haha"}}'
echo "127.0.0.1 www.baidu.com" >> /etc/hosts
curl 'http://www.baidu.com/'  -I 
```

#usage 
```sh
./mock --datapath ./test --severport 8090
curl '127.0.0.1:8089/mocks/' -d '{"url":"http://www.baidu.com:8090/","body":"test","headers":{"a":"b", "c":"d", "Content-Type":"Text/Html2", "Wanghaibo":"haha"}}'
echo "127.0.0.1 www.baidu.com" >> /etc/hosts
curl 'http://www.baidu.com:8090/'  -I 
```

#build
```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
```
