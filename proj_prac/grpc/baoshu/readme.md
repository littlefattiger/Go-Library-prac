It is from https://www.cnblogs.com/baoshu/p/13488106.html
---
### proto
Need following command to generate message struct and service
protoc --go_out=./service .\pbfiles\Product.proto
protoc --go-grpc_out=./service .\pbfiles\Product.proto

### SSL
https://ningyu1.github.io/site/post/51-ssl-cert/
https://monkeywie.cn/2019/11/15/create-ssl-cert-with-san/
https://colobu.com/2016/06/07/simple-golang-tls-examples/
Know about Root authorication and server CA, server request: https://www.ethanzhang.xyz/cfssl%E4%BD%BF%E7%94%A8%E6%96%B9%E6%B3%95/
根服务器证书，服务器证书， 客户端证书
根证书请求文件， 服务器证书请求文件


Command:
.\openssl genrsa -out ca.key 2048
.\openssl req -sha256 -new -x509 -days 365 -key ca.key -out ca.crt   -subj "/C=US/ST=AZ/L=Phoenix/O=PYPL/OU=EDP/CN=testRoot"
    
.\openssl genrsa -out server.key 2048
# .\openssl req -new -subj "/CN=foobar.mydomain.svc" -addext "subjectAltName = DNS:foobar.mydomain.svc" -key server.key -out server.csr
.\openssl req -new -sha256 -key server.key -subj "/C=US/ST=AZ/L=Phoenix/O=PYPL/OU=EDP/CN=testRoot" -reqexts SAN -config ./openssl_local.cnf -out server.csr
    

.\openssl rsa -in server.key -out server_no_password.key
#.\openssl x509 -req -days 365 -in server.csr -signkey server_no_password.key -out server.crt

.\openssl ca -in server.csr -md sha256 -keyfile ca.key -cert ca.crt -extensions SAN -config ./openssl_local.cnf -out server.crt
    
    

2022/04/16 14:45:54 调用gRPC方法错误: rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: x509: certificate relies on legacy Common Name field, use SANs instead"

What I learn here:
1, self sign certification for root authorization;
2, Sign the CA for server
3, Generate a key for server and client