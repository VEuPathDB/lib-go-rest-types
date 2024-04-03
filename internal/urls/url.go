package urls

import "strings"

func New(url string) URL {
	idx := strings.Index(url, "://")

	if idx < 1 {
		panic("URL missing protocol: " + url)
	}

	protocol := url[:idx]
	split := strings.Split(url[idx+3:], "/")
	tmp := make([]string, 0, len(split)+1)

	tmp = append(tmp, protocol)

	for _, part := range split {
		if len(part) > 1 {
			tmp = append(tmp, part)
		}
	}

	if len(tmp) == 1 {
		panic("Malformed URL: " + url)
	}

	out := make([]string, len(tmp))
	copy(out, tmp)

	return URL{out}
}

type URL struct {
	parts []string
}

func (u URL) Protocol() string {
	return u.parts[0]
}

func (u URL) Host() string {
	return u.parts[1]
}
