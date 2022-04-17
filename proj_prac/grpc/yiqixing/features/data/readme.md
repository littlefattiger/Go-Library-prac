openssl genrsa -out ca.key 2048
openssl req -new -key ca.key -out ca.csr  -subj "/C=US/ST=AZ/L=Phoenix/O=PYPL/OU=EDP/CN=xiaomingdou.com"
openssl req -new -x509 -days 3650 -key ca.key -out ca.crt  -subj "/C=US/ST=AZ/L=Phoenix/O=PYPL/OU=EDP/CN=xiaomingdou.com"
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/C=GB/L=China/O=lixd/CN=www.lixueduan.com" -reqexts SAN -config ./openssl_local.cnf	
openssl x509 -req -days 3650 -in server.csr -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extensions SAN -extfile ./openssl_local.cnf
