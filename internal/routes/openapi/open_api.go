package openapi

import (
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
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
		// Models for database interaction
		"PSQLLeaderboardData": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("comp_id", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("comp_name", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("team_id", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("team_name", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("team_abbr", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("games_played", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("win_count", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("draw_count", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("loss_count", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("bye", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("points_for", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("points_against", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("tries_for", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("tries_against", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("bonus_points_try", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("bonus_points_losing", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("bonus_points", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("points_diff", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("points", openapi3.NewStringSchema().
					WithNullable())),
		"PSQLLeaderboardDataList": openapi3.NewArraySchema().
			WithItems(&openapi3.Schema{
				Type: openapi3.TypeArray,
				Items: &openapi3.SchemaRef{
					Ref: "#/components/schemas/PSQLLeaderboardData",
				}}).Items,
		// Data transfer objects
		"TeamCompetitionStats": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("gamesPlayed", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("winCount", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("drawCount", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("lossCount", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("bye", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("pointsFor", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("pointsAgainst", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("triesFor", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("triesAgainst", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("bonusPointsTry", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("bonusPointsLosing", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("bonusPoints", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("pointsDiff", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("points", openapi3.NewStringSchema().
					WithNullable())),
		"TeamLeaderboardData": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("id", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("name", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("abbr", openapi3.NewStringSchema().
					WithNullable()).
				WithPropertyRef("competitionStats", &openapi3.SchemaRef{
					Ref: "#/components/schemas/TeamCompetitionStats",
				})),
		"TeamLeaderboardDataList": openapi3.NewArraySchema().
			WithItems(&openapi3.Schema{
				Type: openapi3.TypeArray,
				Items: &openapi3.SchemaRef{
					Ref: "#/components/schemas/TeamLeaderboardData",
				}}).Items,
		"CompetitionLeaderboardData": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("compId", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("compName", openapi3.NewStringSchema().
					WithNullable()).
				WithPropertyRef("teams", &openapi3.SchemaRef{
					Ref: "#/components/schemas/TeamLeaderboardDataList",
				})),
		"CompetitionLeaderboardDataList": openapi3.NewArraySchema().
			WithItems(&openapi3.Schema{
				Type: openapi3.TypeArray,
				Items: &openapi3.SchemaRef{
					Ref: "#/components/schemas/CompetitionLeaderboardData",
				}}).Items,
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
		"LeaderboardRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for fetching leaderboard data").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("source", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("compId", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("compName", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("date", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("source", openapi3.NewStringSchema().
						WithMinLength(1))),
		},
	}

	swagger.Components.Responses = openapi3.Responses{
		"LeaderboardResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response with competition data").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("leaderboardData", &openapi3.SchemaRef{
						Ref: "#/components/schemas/CompetitionLeaderboardData",
					}).
					WithPropertyRef("message", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Message",
					}).
					WithNullable())),
		},
		"AllLeaderboardsResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response with competition data").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("leaderboardDataList", &openapi3.SchemaRef{
						Ref: "#/components/schemas/CompetitionLeaderboardDataList",
					}).
					WithPropertyRef("message", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Message",
					}).
					WithNullable())),
		},
	}

	swagger.Paths = openapi3.Paths{
		"/leaderboard": &openapi3.PathItem{
			Description: "Leaderboard Data",
			Post: &openapi3.Operation{
				OperationID: "GetLeaderboardData",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/LeaderboardRequest",
				},
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/LeaderboardResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/LeaderboardResponse",
					},
					"201": &openapi3.ResponseRef{
						Ref: "#/components/responses/LeaderboardResponse",
					},
				},
			},
		},
		"/leaderboards": &openapi3.PathItem{
			Description: "All Leaderboard Data",
			Get: &openapi3.Operation{
				OperationID: "GetAllLeaderboardData",
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/AllLeaderboardsResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/AllLeaderboardsResponse",
					},
					"201": &openapi3.ResponseRef{
						Ref: "#/components/responses/AllLeaderboardsResponse",
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
		response.RenderResponse(w, &swagger, http.StatusOK)
	}).Methods(http.MethodGet)

	r.HandleFunc("/openapi3.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")

		data, _ := yaml.Marshal(&swagger)

		_, _ = w.Write(data)

		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
}
