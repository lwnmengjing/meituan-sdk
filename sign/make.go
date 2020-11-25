package sign

import "net/url"

// Make 生成签名
func Make(secret, scheme, host, path string, values Values) string {
	return make(secret, scheme, host, path, values)
}

// make 生成签名(私有)
func make(secret, scheme, host, path string, values Values) string {
	u := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
	u.RawQuery = values.Encode()
	sign := md5V(u.String() + secret)
	return sign
}
