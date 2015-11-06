package parser
import (
	"testing"
	"reflect"
	"fmt"
)

type TestCase struct {
	spec []byte
	expectedApi *Api
	expectedErr error
}

var testCases = []TestCase {
	{
		spec: []byte(`GET https://www.alvarloes.com/posts/:id/comments/:id
				<- {
					"id":"4567",
					"title":"I like it",
					"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target"
				}
				`),
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
						"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target",
					},
				},
			},
		},
		expectedErr: nil,
	},
}

func TestApi(t *testing.T) {
	for _,testCase := range testCases {
		api, err := NewApi(testCase.spec)
		fmt.Println(api,err)
		if ok := reflect.DeepEqual(testCase.expectedErr, err); !ok {
			t.Errorf("Expected error '%v', got: '%v'\n",testCase.expectedErr, err)
		}
		if ok := reflect.DeepEqual(testCase.expectedApi, api); !ok {
			t.Errorf("Expected api '%v', got: '%v'\n",testCase.expectedApi, api)
		}
	}
}
