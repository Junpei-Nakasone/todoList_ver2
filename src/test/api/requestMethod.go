package api

type RequestMethod int

const (
	Get RequestMethod = 1 << iota

	Post

	Put

	Delete
)

func (r RequestMethod) Strings() []string {
	var result []string

	if r&Get == Get {
		result = append(result, "GET")
	}

	if r&Post == Post {
		result = append(result, "POST")
	}

	if r&Put == Put {
		result = append(result, "PUT")
	}

	if r&Delete == Delete {
		result = append(result, "DELETE")
	}
	return result
}
