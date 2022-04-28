# https密钥和证书生成



## 1、创建密钥

使用openssl工具生成一个RSA私钥

```bash
openssl genrsa -des3 -out server.key 2048
```

## 2、生成CSR（证书签名请求）

```bash
openssl req -new -key server.key -out server.csr
```

说明：需要依次输入国家，地区，城市，组织，组织单位，Common Name和Email。其中Common Name，可以写自己的名字或者域名，如果要支持https，Common Name应该与域名保持一致，否则会引起浏览器警告。

## 3、删除密钥中的密码

```bash
openssl rsa -in server.key -out server.key
```

说明：如果不删除密码，在应用加载的时候会出现输入密码进行验证的情况，不方便自动化部署。

## 4、生成自签名证书

```bash
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
```

## 5、生成pem格式的公钥

```bash
openssl x509 -in server.crt -out server.pem -outform PEM
```


