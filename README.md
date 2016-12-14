xcrawl
==============================

Tools to crawl html and read in xpath.

## Install
```
$ go get github.com/kyokomi/xcrawl
```

## Config Example
```yaml
headers:
  User-Agent: Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Mobile Safari/537.36
```

## Usage
```
$ xcrawl -i <input html url> -x <scraping xpath> -c <crawler config>
```

## Demo
```
$ xcrawl -i 'http://example.com' -x '//*[@id="main-contents"]/ul[1]/li[3]/div/div[1]/div/a[1]/@href' -c config.yaml
http://hogehoge.com
```

## Licence
[MIT](https://github.com/kyokomi/xcrawl/blob/master/LICENCE)
