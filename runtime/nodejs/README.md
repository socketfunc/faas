# runtime for Node.js

`PARCEL`を使用してハンドラーをビルドして1つにまとめる

### config

```yaml
runtime: nodejs
handler:
    name: room.join
```

### handler

```javascript
exports.handler = async (ctx, req, res) => {
  res.send('topic', 'event', 'message')
}
```
