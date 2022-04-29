package matchers

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mniak/gomock-contrib/internal/utils"
)

func TestFunctionalTests(t *testing.T) {
	t.Run("Struct with field Message that is a string pointer containing JSON", func(t *testing.T) {
		expectedMessage := map[string]any{
			"domain": "migration",
			"org_id": gofakeit.UUID(),
			"data": map[string]any{
				"operation": "INSERT",
				"status":    "SUCCESS",
				"code":      "MIGR-0001",
				"message":   "The token was migrated with success",
				"migration": map[string]any{
					"id": "id-value",
				},
				"entity": map[string]any{},
			},
		}
		type SampleType struct {
			Message *string
		}

		sample := SampleType{
			Message: utils.ToPointer(`{
				"event_id": "event-id",
				"domain": "migration",
				"event_type": "account_outgoing",
				"schema_version": 1,
				"org_id": "TN-34125ec5-d288-4601-8073-f60ee429088c",
				"timestamp": "2022-04-29T14:50:41.342-03:00",
				"cid": "",
				"data": {
					"operation": "INSERT",
					"status": "SUCCESS",
					"code": "MIGR-0001",
					"message": "The token was migrated with success",
					"migration": {
						"id": "",
						"card_migration_id": "",
						"version_date": "",
						"phase_id": ""
					},
					"entity":{
						"card_id": null,
						"token_id": null,
						"account_id": null,
						"customer_id": null,
						"program_id": null
					}
				}
			}`),
		}
		sut := HasField("Message").ThatMatches(
			IsJSON().ThatMatches(LikeMap(expectedMessage)),
		)

		sut.Matches(sample)
		_ = sut.Got(sample)
		_ = sut.String()
	})
}
