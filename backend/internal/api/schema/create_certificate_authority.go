package schema

import "fmt"

// CreateCertificateAuthority is the schema for incoming data validation
func CreateCertificateAuthority() string {
	return fmt.Sprintf(`
		{
			"type": "object",
			"additionalProperties": false,
			"required": [
				"name",
				"acmesh_server",
				"max_domains"
			],
			"properties": {
				"name": %s,
				"acmesh_server": %s,
				"max_domains": %s
			}
		}
	`, stringMinMax(1, 100), stringMinMax(2, 255), intMinOne)
}
