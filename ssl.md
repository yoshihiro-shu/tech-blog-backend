# SSL証明書

## 鍵生成用のディレクトリ作成

```zsh
mkdir nginx/ssl
```

## 秘密鍵の作成


```zsh
sudo openssl genrsa -out nginx/ssl/server.key 2048
```

ドメインが決まっているなら「Common Name」だけ入力

## CSR（証明書署名要求）の作成

```zsh
sudo openssl req -new -key nginx/ssl/server.key -out nginx/ssl/server.csr
```

```zsh
sudo openssl x509 -days 3650 -req -signkey nginx/ssl/server.key -in nginx/ssl/server.csr -out nginx/ssl/server.crt
```

## CRT（SSLサーバ証明書）の作成

key chain accessに登録

https://www.curict.com/item/2a/2ace791.html
