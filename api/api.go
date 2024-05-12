package api

import (
	_ "embed"
)

// the embed is done out side the *RegisterRoutes* Function because
// it does not apply var inside any function

//go:embed openapi.yaml
var openApi string

func OpenApiDocs() string {
	return openApi
}
