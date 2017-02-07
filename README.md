# run
```sh
usage of ./mock:
--adminport string    admin port (default "8091")
--datapath string     path to store mockinfo
--serverport string   server port (default "8090")
```

#usage 
```sh
./mock --datapath ./test --severport 8090
curl '127.0.0.1:8089/mocks/' -d '{"url":"http://www.baidu.com:8090/","body":"test","headers":{"a":"b", "c":"d", "Content-Type":"Text/Html2", "Wanghaibo":"haha"}}'
echo "127.0.0.1 www.baidu.com" >> /etc/hosts
curl 'http://www.baidu.com:8090/'  -I 
```
