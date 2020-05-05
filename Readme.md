# japanese-zip-code
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

japanese-zip-code は[ZipCloudAPI](https://zip-cloud.appspot.com/rule/api) を簡素に実施するためのパッケージです。

APIを無駄に叩かないように、APIに合わせた郵便番号のチェックを行い、不正なものは事前にエラーとします。

## Usage
```
jzc -z zip-code
  -z : Input zip code. This format is "[0-9]{3}[-]?[0-9]{4}"
```

## Note
このパッケージは[ZipCloud](https://zip-cloud.appspot.com/rule/api) 非公式パッケージです。  
上記サービス運用者の迷惑にならない利用をお願いいたします。

```
This package is an unofficial ZipCloud package.
Please use the service without disturbing the service operator.
```

## License
japanes-zip-code is released under the MIT License, see LICENSE.txt.