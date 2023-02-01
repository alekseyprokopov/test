package main

import (
	"encoding/json"
	"fmt"
)

const response = `{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
}`

type (
	Response struct {
		Header HeaderType `json:"header"`
		Data   DataType   `json:"data,"`
	}

	HeaderType struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	DataType []DataItem

	DataItem struct {
		Type       string        `json:"type"`
		Id         int           `json:"id"`
		Attributes AtributesType `json:"attributes"`
	}
	AtributesType struct {
		Email      string `json:"email"`
		ArticleIds []int  `json:"articleIds"`
	}
)

func ReadResponse(item string) (Response, error) {
	resp := Response{}
	if err := json.Unmarshal([]byte(item), &resp); err != nil {
		return Response{}, fmt.Errorf("JSON unmarshal: %w", err)
	}
	return resp, nil
}
