package gen

import (
	"testing"

	"github.com/alvaroloes/sdkgen/parser"
	"github.com/alvaroloes/sdkgen/tests"
	"github.com/kr/pretty"
)

const (
	failModelsInfoFormat = "Test %q: Didn't get the expected models info. Differences are:\n%v"
)

type modelsInfoTestCase struct {
	name               string
	api                *parser.Api
	expectedModelsInfo map[string]*modelInfo
}

var modelsInfoTestCases = []modelsInfoTestCase{
	{
		name: "Simple. Only one GET endpoint",
		api: &parser.Api{
			Endpoints: []parser.Endpoint{
				{
					Method:    parser.GET,
					URLString: "https://www.alvarloes.com/posts/:id/comments/:id",
					URL:       tests.MustParseURL("https://www.alvarloes.com/posts/:id/comments/:id"),
					Resources: []parser.Resource{
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
		expectedModelsInfo: map[string]*modelInfo{
			"comments": {
				Name: "comment",
				Properties: map[string]property{
					"id": {
						Name:    "id",
						Type:    "string",
						IsArray: false,
					},
					"title": {
						Name:    "title",
						Type:    "string",
						IsArray: false,
					},
					"body": {
						Name:    "body",
						Type:    "string",
						IsArray: false,
					},
				},
				EndpointsInfo: []endpointInfo{
					{
						Method:  parser.GET,
						URLPath: "/posts/:id/comments/:id",
						SegmentParams: []string{
							"id",
							"id",
						},
						ResponseType: 0,
					},
				},
			},
		},
	},
}

func TestModelsInfo(t *testing.T) {
	for _, testCase := range modelsInfoTestCases {
		gen := Generator{
			gen:    nil,
			api:    testCase.api,
			config: Config{},
		}
		gen.extractModelsInfo()

		diffs := pretty.Diff(testCase.expectedModelsInfo, gen.modelsInfo)

		if len(diffs) > 0 {
			t.Errorf(failModelsInfoFormat, testCase.name, tests.FormattedDiff(diffs))
		}
	}
}
