package parser

import (
	"reflect"
	"testing"

	"github.com/alvaroloes/sdkgen/tests"
	"github.com/kr/pretty"
)

const (
	failErrorFormat = "Test %q: Expected error %q, got: %q"
	failAPIFormat   = "Test %q: Didn't get the expected API. Differences are:\n%v"
)

type testCase struct {
	name        string
	spec        []byte
	expectedAPI *API
	expectedErr error
}

var testCases = []testCase{
	{
		name: "Simple. Only response object",
		spec: []byte(`GET https://www.alvarloes.com/posts/:id/comments/:id
			<- {
				"id":"4567",
				"title":"I like it",
				"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client languages your API target"
			}`),
		expectedAPI: &API{
			Endpoints: []Endpoint{
				{
					Method:    GET,
					URL:       tests.MustParseURL("https://www.alvarloes.com/posts/:id/comments/:id"),
					Resources: []Resource{
						{
							Name:       "posts",
							Parameters: []string{"id"},
						}, {
							Name:       "comments",
							Parameters: []string{"id"},
						},
					},
					RequestBody: nil,
					ResponseBody: map[string]interface{}{
						"id":    "4567",
						"title": "I like it",
						"body":  "I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client languages your API target",
					},
				},
			},
		},
		expectedErr: nil,
	}, {
		name: "Simple. Request and response object",
		spec: []byte(`POST https://www.alvarloes.com/posts/:id/comments
			-> {
				"title":"I like it",
				"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client languages your API target"
			}
			<- {
				"id":"4567",
				"title":"I like it",
				"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client languages your API target"
			}`),
		expectedAPI: &API{
			Endpoints: []Endpoint{
				{
					Method:    POST,
					URL:       tests.MustParseURL("https://www.alvarloes.com/posts/:id/comments"),
					Resources: []Resource{
						{
							Name:       "posts",
							Parameters: []string{"id"},
						}, {
							Name:       "comments",
							Parameters: nil,
						},
					},
					RequestBody: map[string]interface{}{
						"title": "I like it",
						"body":  "I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client languages your API target",
					},
					ResponseBody: map[string]interface{}{
						"id":    "4567",
						"title": "I like it",
						"body":  "I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client languages your API target",
					},
				},
			},
		},
		expectedErr: nil,
	}, {
		name: "Simple. Response array",
		spec: []byte(`GET https://www.alvarloes.com/posts/:id/comments
			<- [
				{
					"id":"4567",
					"title":"I like it",
					"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client languages your API target"
				},{
					"id":"4567",
					"title":"I like it too",
					"body":"I like it, but it seems to be really hard to come up with a powerfull and flexible generator"
				}
			]`),
		expectedAPI: &API{
			Endpoints: []Endpoint{
				{
					Method:    GET,
					URL:       tests.MustParseURL("https://www.alvarloes.com/posts/:id/comments"),
					Resources: []Resource{
						{
							Name:       "posts",
							Parameters: []string{"id"},
						}, {
							Name:       "comments",
							Parameters: nil,
						},
					},
					RequestBody: nil,
					ResponseBody: []interface{}{
						map[string]interface{}{
							"id":    "4567",
							"title": "I like it",
							"body":  "I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client languages your API target",
						},
						map[string]interface{}{
							"id":    "4567",
							"title": "I like it too",
							"body":  "I like it, but it seems to be really hard to come up with a powerfull and flexible generator",
						},
					},
				},
			},
		},
		expectedErr: nil,
	}, {
		name: "Simple. No request nor response",
		spec: []byte(`DELETE https://www.alvarloes.com/posts/:id`),
		expectedAPI: &API{
			Endpoints: []Endpoint{
				{
					Method:    DELETE,
					URL:       tests.MustParseURL("https://www.alvarloes.com/posts/:id"),
					Resources: []Resource{
						{
							Name:       "posts",
							Parameters: []string{"id"},
						},
					},
					RequestBody:  nil,
					ResponseBody: nil,
				},
			},
		},
		expectedErr: nil,
	}, {
		name: "Complex. Response array",
		spec: []byte(`GET https://www.alvarloes.com/posts
			<- [
				{
					"id":"1234",
					"author":{
						"name":"John",
						"age":20
					},
					"title":"We really need a client SDK generator",
					"body":"(...) we to make the machine work for us, thus we should write generators to make the computer write the non-creative part of the code for us",
					"comments":[
						{
							"id":"4567",
							"title":"I like it",
							"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target"
						},{
							"id":"4567",
							"title":"I like it too",
							"body":"I like it, but it seems to be really hard to come up with a powerfull and flexible generator"
						}
					]
				},{
					"id":"12345",
					"author":{
						"name":"John",
						"age":20
					},
					"title":"We really need a client SDK generator",
					"body":"(...) we to make the machine work for us, thus we should write generators to make the computer write the non-creative part of the code for us",
					"comments":[
						{
							"id":"4567",
							"title":"I like it",
							"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target"
						},{
							"id":"4567",
							"title":"I like it too",
							"body":"I like it, but it seems to be really hard to come up with a powerfull and flexible generator"
						}
					]
				}
			]`),
		expectedAPI: &API{
			Endpoints: []Endpoint{
				{
					Method:    GET,
					URL:       tests.MustParseURL("https://www.alvarloes.com/posts"),
					Resources: []Resource{
						{
							Name:       "posts",
							Parameters: nil,
						},
					},
					RequestBody: nil,
					ResponseBody: []interface{}{
						map[string]interface{}{
							"id": "1234",
							"author": map[string]interface{}{
								"name": "John",
								"age":  float64(20),
							},
							"title": "We really need a client SDK generator",
							"body":  "(...) we to make the machine work for us, thus we should write generators to make the computer write the non-creative part of the code for us",
							"comments": []interface{}{
								map[string]interface{}{
									"id":    "4567",
									"title": "I like it",
									"body":  "I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target",
								},
								map[string]interface{}{
									"id":    "4567",
									"title": "I like it too",
									"body":  "I like it, but it seems to be really hard to come up with a powerfull and flexible generator",
								},
							},
						}, map[string]interface{}{
							"id": "12345",
							"author": map[string]interface{}{
								"name": "John",
								"age":  float64(20),
							},
							"title": "We really need a client SDK generator",
							"body":  "(...) we to make the machine work for us, thus we should write generators to make the computer write the non-creative part of the code for us",
							"comments": []interface{}{
								map[string]interface{}{
									"id":    "4567",
									"title": "I like it",
									"body":  "I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target",
								},
								map[string]interface{}{
									"id":    "4567",
									"title": "I like it too",
									"body":  "I like it, but it seems to be really hard to come up with a powerfull and flexible generator",
								},
							},
						},
					},
				},
			},
		},
		expectedErr: nil,
	},
}

func TestAPI(t *testing.T) {
	for _, testCase := range testCases {
		api, err := NewAPI(testCase.spec)

		if ok := reflect.DeepEqual(testCase.expectedErr, err); !ok {
			t.Errorf(failErrorFormat, testCase.name, testCase.expectedErr, err)
		}

		if diff := pretty.Diff(testCase.expectedAPI, api); len(diff) > 0 {
			t.Errorf(failAPIFormat, testCase.name, tests.FormattedDiff(diff))
		}
	}
}
