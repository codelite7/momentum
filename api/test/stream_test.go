package test

import (
	"bufio"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

type Request struct {
	Input []Input `json:"input"`
}

type Input struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

func TestStream(t *testing.T) {
	r := &Request{
		Input: []Input{{
			Content: "Tell me a short joke",
			Type:    "human"}},
	}
	client := resty.New()
	resp, err := client.R().SetDoNotParseResponse(true).SetBody(r).Post("http://0.0.0.0:8000/openai/stream_log")
	require.NoError(t, err)
	defer resp.RawResponse.Body.Close()

	scanner := bufio.NewScanner(resp.RawResponse.Body)
	reply := ""
	for scanner.Scan() {
		_res := scanner.Text()
		if _res == "" {
			continue
		}
		reply += _res
		t.Log(_res)
	}

	t.Log(reply)
	//dec := json.NewDecoder(resp.RawResponse.Body)
	//for {
	//	var x map[string]interface{}
	//	err := dec.Decode(&x)
	//	if err != nil {
	//		break
	//	}
	//	log.Println(x)
	//}
}
