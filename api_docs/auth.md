## ログイン (POST "/login")
メールアドレスとパスワードを送ると、メールアドレスと電話番号が認証されている場合、アクセストークンとリフレッシュトークンが付与されます。

メールアドレスが認証されてない場合はinvalid.infoが返却されます。

Request

```json
{
  "email": "test@example.com",
  "password": "password1234"
}
```

Response(success 200)

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}

Response(error 400)

```zsh
"invalid.info"
```

## リフレッシュ (POST "/refresh_token")

リフレッシュトークンをリクエストすると、有効だった場合アクセストークンとリフレッシュトークンが新規発行され、返却されます。

Request

```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

Response (success 200)

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```

Response(error 400) トークン無効

```zsh
"invalid.token"
```
