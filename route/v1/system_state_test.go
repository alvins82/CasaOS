package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IceWhaleTech/CasaOS/model"
	"github.com/IceWhaleTech/CasaOS/service"
	"github.com/labstack/echo/v4"
)

type fakeSystemStateService struct {
	service.SystemService
	rebootErr   error
	shutdownErr error
}

func (f *fakeSystemStateService) SystemReboot() error {
	return f.rebootErr
}

func (f *fakeSystemStateService) SystemShutdown() error {
	return f.shutdownErr
}

type fakeSystemStateRepository struct {
	service.Repository
	system service.SystemService
}

func (f fakeSystemStateRepository) System() service.SystemService {
	return f.system
}

func TestPutSystemStatePropagatesPowerOperationErrors(t *testing.T) {
	original := service.MyService
	t.Cleanup(func() { service.MyService = original })

	fakeSystem := &fakeSystemStateService{
		rebootErr:   errors.New("power operation failed: reboot"),
		shutdownErr: errors.New("power operation failed: shutdown"),
	}
	service.MyService = fakeSystemStateRepository{system: fakeSystem}

	for _, test := range []struct {
		name  string
		state string
		err   error
	}{
		{name: "shutdown", state: "off", err: fakeSystem.shutdownErr},
		{name: "reboot", state: "restart", err: fakeSystem.rebootErr},
	} {
		t.Run(test.name, func(t *testing.T) {
			e := echo.New()
			request := httptest.NewRequest(http.MethodPut, "/v1/sys/state/"+test.state, nil)
			recorder := httptest.NewRecorder()
			context := e.NewContext(request, recorder)
			context.SetPath("/v1/sys/state/:state")
			context.SetParamNames("state")
			context.SetParamValues(test.state)

			if err := PutSystemState(context); err != nil {
				t.Fatalf("PutSystemState() error = %v", err)
			}
			if recorder.Code != http.StatusInternalServerError {
				t.Fatalf("status = %d, want %d", recorder.Code, http.StatusInternalServerError)
			}

			var response model.Result
			if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
				t.Fatalf("decode response: %v", err)
			}
			if response.Success != http.StatusInternalServerError {
				t.Fatalf("response success = %d, want %d", response.Success, http.StatusInternalServerError)
			}
			if response.Message != test.err.Error() {
				t.Fatalf("response message = %q, want %q", response.Message, test.err.Error())
			}
		})
	}
}
