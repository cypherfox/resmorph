//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config-server.yaml ResmorphTracker-0.0.1-oapi3.0.3.yml

package tracker

import (
	"context"

	"github.com/cypherfox/snacks-manager/internal/backend"
)

type ApiServer struct {
	Backend *backend.SnackBackEnd
}

var _ StrictServerInterface = (*ApiServer)(nil)

// GetResourcesResourceId implements StrictServerInterface.
func (a *ApiServer) GetResourcesResourceId(ctx context.Context, request GetResourcesResourceIdRequestObject) (GetResourcesResourceIdResponseObject, error) {
	panic("unimplemented")
}

// GetTest implements StrictServerInterface.
func (a *ApiServer) GetTest(ctx context.Context, request GetTestRequestObject) (GetTestResponseObject, error) {
	panic("unimplemented")
}

// PostNotify implements StrictServerInterface.
func (a *ApiServer) PostNotify(ctx context.Context, request PostNotifyRequestObject) (PostNotifyResponseObject, error) {
	panic("unimplemented")
}
