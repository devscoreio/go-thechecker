package thechecker

type Response struct {
	shared     *Client
	Result     string `json:"result"`
	Reason     string `json:"reason"`
	Email      string `json:"email"`
	User       string `json:"user"`
	Domain     string `json:"domain"`
	Role       bool   `json:"role"`
	Disposable bool   `json:"disposable"`
	AcceptAll  bool   `json:"accept_all"`
	DidYouMean string `json:"did_you_mean"`
}

func (c *Response) Check(email string) (*Response, error) {
	if email == "" {
		return nil, fmt.Errorf("Missing email")
	}
	r, err := c.shared.get(email)
	if err != nil {
		return nil, err
	}
	return c.search(r)
}

func (c *Response) search(r *http.Request) (*Response, error) {
	resp, err := c.shared.do(r)
	if err != nil {
		return nil, err
	}
	var response Response
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
