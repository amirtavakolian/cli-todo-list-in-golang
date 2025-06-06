package response

import "strconv"

type Response struct {
	Content string
	Status  int
}

func (response *Response) SetContent(content string) {
	response.Content = content
}

func (response *Response) SetStatus(status int) {
	response.Status = status
}

func (response *Response) BuildResponse() string {
	return "\n" + response.Content + "\nstatus code: " + strconv.Itoa(response.Status)
}
