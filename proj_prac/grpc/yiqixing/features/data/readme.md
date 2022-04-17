openssl genrsa -out ca.key 2048
openssl req -new -key ca.key -out ca.csr  -subj "/C=US/ST=AZ/L=Phoenix/O=PYPL/OU=EDP/CN=xiaomingdou.com"
openssl req -new -x509 -days 3650 -key ca.key -out ca.crt  -subj "/C=US/ST=AZ/L=Phoenix/O=PYPL/OU=EDP/CN=xiaomingdou.com"
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/C=US/ST=AZ/L=Phoenix/O=PYPL/OU=EDP/CN=xiaomingdou.com"  -reqexts SAN -config ./openssl_local.cnf	
openssl x509 -req -days 3650 -in server.csr -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extensions SAN -extfile ./openssl_local.cnf

openssl genrsa -out client.key 2048

openssl req -new -key client.key -out client.csr -subj "/C=US/ST=AZ/L=Phoenix/O=PYPL/OU=EDP/CN=xiaomingdou.com"  -reqexts SAN -config ./openssl_local.cnf	
openssl x509 -req -days 3650 -in client.csr -out client.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extensions SAN -extfile ./openssl_local.cnf

---
```
error:rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0"
```
Will see this after go 1.15
Cause:

Go 1.15 版本开始废弃 CommonName 并且推荐使用 SAN 证书，导致依赖 CommonName 的证书都无法使用了。

Solution:

- 1）开启兼容：设置环境变量 GODEBUG 为 x509ignoreCN=0
- 2）使用 SAN 证书
```
[SAN]
subjectAltName=DNS:*.xiaomingdou.com,DNS:*.xiaomingdou1.com
```