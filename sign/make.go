package sign

import "net/url"

// Make 生成签名
func Make(secret, host, path string, values Values) string {
	return make(secret, host, path, values)
}

func make(secret, host, path string, values Values) string {
	u := &url.URL{
		Host: host,
		Path: path,
	}
	u.RawQuery = values.Encode()
	sign := md5V(u.String() + secret)
	return sign
}
