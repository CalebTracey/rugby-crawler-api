package routes

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
	"net/http"
)

//go:generate go run ../../cmd/openapi-gen/main.go -path ../../swagger-ui

//go:generate oapi-codegen -package openapi3 -generate types  -o ../../pkg/openapi3/types.gen.go ../../swagger-ui/openapi3.yaml
//go:generate oapi-codegen -package openapi3 -generate client -o ../../pkg/openapi3/client.gen.go ../../swagger-ui/openapi3.yaml

//go:generate statik -src=/Users/calebtracey/Desktop/Code/rugby-crawler-api/swagger-ui

// NewOpenAPI3 instantiates the OpenAPI specification for this service.
func NewOpenAPI3() openapi3.T {
	swagger := openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "Rugby Data Crawler REST API",
			Description: "REST API rugby data web crawler",
			Version:     "0.0.0",
			License: &openapi3.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
			Contact: &openapi3.Contact{
				URL: "https://github.com/CalebTracey/rugby-crawler-api",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Local development",
				URL:         "http://0.0.0.0:6080",
			},
		},
	}

	swagger.Components.Schemas = openapi3.Schemas{
		"CompetitionCrawlRequest": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("competitionID", openapi3.NewStringSchema()).
				WithProperty("date", openapi3.NewStringSchema())),
		"CompetitionCrawlResponse": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("compId", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("name", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("teamIds", openapi3.NewArraySchema().
					WithItems(&openapi3.Schema{
						Type: openapi3.TypeString,
					})).
				WithPropertyRef("message", &openapi3.SchemaRef{
					Ref: "#/components/schemas/Message",
				}).
				WithNullable()),
		"ErrorLog": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("scope", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("status", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("trace", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("rootCause", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("query", openapi3.NewStringSchema().
					WithNullable())),
		"ErrorLogs": openapi3.NewArraySchema().
			WithItems(&openapi3.Schema{
				Type: openapi3.TypeArray,
				Items: &openapi3.SchemaRef{
					Ref: "#/components/schemas/ErrorLog",
				}}).Items,
		"Message": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithPropertyRef("errorLog", &openapi3.SchemaRef{
					Ref: "#/components/schemas/ErrorLogs",
				}).
				WithProperty("hostName", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("status", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("timeTaken", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("count", openapi3.NewStringSchema().
					WithNullable())),
	}

	swagger.Components.RequestBodies = openapi3.RequestBodies{
		"CompetitionCrawlRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for crawling competition data").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("competitionID", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("date", openapi3.NewStringSchema().
						WithMinLength(1))),
		},
	}

	swagger.Components.Responses = openapi3.Responses{
		"CompetitionCrawlResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response with competition crawl results").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("compId", openapi3.NewStringSchema().
						WithNullable()).
					WithProperty("name", openapi3.NewStringSchema().
						WithNullable()).
					WithItems(&openapi3.Schema{
						Title: "teamIds",
						Type:  openapi3.TypeString,
					}).
					WithPropertyRef("message", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Message",
					}).
					WithNullable())),
		},
	}

	swagger.Paths = openapi3.Paths{
		"/competition": &openapi3.PathItem{
			Summary:     "Crawl Requests",
			Description: "Crawl Competition",
			Post: &openapi3.Operation{
				OperationID: "CrawlComp",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/CompetitionCrawlRequest",
				},
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/CompetitionCrawlResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/CompetitionCrawlResponse",
					},
					"201": &openapi3.ResponseRef{
						Ref: "#/components/responses/CompetitionCrawlResponse",
					},
				},
			},
		},
	}

	return swagger
}

func RegisterOpenAPI(r *mux.Router) {
	swagger := NewOpenAPI3()

	r.HandleFunc("/openapi3.json", func(w http.ResponseWriter, r *http.Request) {
		renderResponse(w, &swagger, http.StatusOK)
	}).Methods(http.MethodGet)

	r.HandleFunc("/openapi3.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")

		data, _ := yaml.Marshal(&swagger)

		_, _ = w.Write(data)

		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
}
