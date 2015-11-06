package parser
import (
	"testing"
	"reflect"
)

const (
	failErrorFormat = `Test "%v": Expected error '%v', got: '%v'\n`
	failApiFormat = `Test "%v": Expected api '%v', got: '%v'\n`
)

type TestCase struct {
	name string
	spec []byte
	expectedApi *Api
	expectedErr error
}

var testCases = []TestCase {
	{
		name: "Simple. Only response object",
		spec: []byte(`GET https://www.alvarloes.com/posts/:id/comments/:id
			<- {
				"id":"4567",
				"title":"I like it",
				"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client languages your API target"
			}`),
		expectedApi: &Api {
			Endpoints:[]Endpoint {
				{
					Method: "GET",
					FullURL: "https://www.alvarloes.com/posts/:id/comments/:id",
					Resources: []Resource{
						{
							Name:"posts",
							Parameters:[]string{"id"},
						}, {
							Name:"comments",
							Parameters:[]string{"id"},
						},
					},
					RequestBody: nil,
					ResponseBody: map[string]interface{} {
						"id":"4567",
						"title":"I like it",
						"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client languages your API target",
					},
				},
			},
		},
		expectedErr: nil,
	},{
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
		expectedApi: &Api {
			Endpoints:[]Endpoint {
				{
					Method: "POST",
					FullURL: "https://www.alvarloes.com/posts/:id/comments",
					Resources: []Resource{
						{
							Name:"posts",
							Parameters:[]string{"id"},
						}, {
							Name:"comments",
							Parameters:nil,
						},
					},
					RequestBody: map[string]interface{} {
						"title":"I like it",
						"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client languages your API target",
					},
					ResponseBody: map[string]interface{} {
						"id":"4567",
						"title":"I like it",
						"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client languages your API target",
					},
				},
			},
		},
		expectedErr: nil,
	},{
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
		expectedApi: &Api {
			Endpoints:[]Endpoint {
				{
					Method: "GET",
					FullURL: "https://www.alvarloes.com/posts/:id/comments",
					Resources: []Resource{
						{
							Name:"posts",
							Parameters:[]string{"id"},
						}, {
							Name:"comments",
							Parameters:nil,
						},
					},
					RequestBody:nil,
					ResponseBody: []interface{} {
						map[string]interface{} {
							"id":"4567",
							"title":"I like it",
							"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client languages your API target",
						},
						map[string]interface{} {
							"id":"4567",
							"title":"I like it too",
							"body":"I like it, but it seems to be really hard to come up with a powerfull and flexible generator",
						},

					},
				},
			},
		},
		expectedErr: nil,
	},{
		name: "Simple. No request nor response",
		spec: []byte(`DELETE https://www.alvarloes.com/posts/:id`),
		expectedApi: &Api {
			Endpoints:[]Endpoint {
				{
					Method: "DELETE",
					FullURL: "https://www.alvarloes.com/posts/:id",
					Resources: []Resource{
						{
							Name:"posts",
							Parameters:[]string{"id"},
						},
					},
					RequestBody:nil,
					ResponseBody:nil,
				},
			},
		},
		expectedErr: nil,
	},
}

func TestApi(t *testing.T) {
	for _,testCase := range testCases {
		api, err := NewApi(testCase.spec)

		if ok := reflect.DeepEqual(testCase.expectedErr, err); !ok {
			t.Errorf(failErrorFormat, testCase.name, testCase.expectedErr, err)
		}
		if ok := reflect.DeepEqual(testCase.expectedApi, api); !ok {
			t.Errorf(failApiFormat, testCase.name, testCase.expectedApi, api)
		}
	}
}
