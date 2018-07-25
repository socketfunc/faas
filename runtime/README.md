# runtime

## Architecture

`Runtime Engine`のPoolをしておいて必要になったときに専用エンジンに置き換える

置き換える部分で`kubernetes`のServiceを生成する

## Manager

- [ ] yamlでのconfig設定
- [ ] ソースコード管理のStorage
- [ ] メトリックス収集

### Go

GolangのPlugin機構を用いてHandlerを動かす

- [ ] ビルド環境
- [ ] plugin機構
- [ ] メトリックス

ゴルーチン、ヒープ、

### Node.js

Webpackを用いてminifyする

- [ ] ビルド環境
- [ ] plugin機構
- [ ] メトリックス

ヒープ

## Runtime Config

```yaml
port: 9000
codePath: /path/to/plugin
```

### Go

version: 1.8 over

拡張子: so

```yaml
runtime: go1.10
```

### Node.js

version: 8.10 over

拡張子: js

```yaml
runtime: nodejs10.6
```
