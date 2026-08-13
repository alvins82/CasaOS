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

type fakeSystemPackageService struct {
	service.SystemService
	updates      service.SystemPackageUpdates
	updateStatus service.SystemPackageUpdateStatus
	checkErr     error
	startErr     error
}

func (f *fakeSystemPackageService) GetSystemPackageUpdates() (service.SystemPackageUpdates, error) {
	return f.updates, f.checkErr
}

func (f *fakeSystemPackageService) StartSystemPackageUpdate() (service.SystemPackageUpdateStatus, error) {
	return f.updateStatus, f.startErr
}

func (f *fakeSystemPackageService) GetSystemPackageUpdateStatus() service.SystemPackageUpdateStatus {
	return f.updateStatus
}

type fakeSystemPackageRepository struct {
	service.Repository
	system service.SystemService
}

func (f fakeSystemPackageRepository) System() service.SystemService {
	return f.system
}

func performSystemPackageRequest(t *testing.T, method, path string) *httptest.ResponseRecorder {
	t.Helper()
	e := echo.New()
	request := httptest.NewRequest(method, path, nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)

	switch {
	case method == http.MethodGet && path == "/v1/sys/packages":
		if err := GetSystemPackageUpdates(context); err != nil {
			t.Fatalf("GetSystemPackageUpdates() error = %v", err)
		}
	case method == http.MethodPost && path == "/v1/sys/packages/update":
		if err := StartSystemPackageUpdate(context); err != nil {
			t.Fatalf("StartSystemPackageUpdate() error = %v", err)
		}
	case method == http.MethodGet && path == "/v1/sys/packages/update/status":
		if err := GetSystemPackageUpdateStatus(context); err != nil {
			t.Fatalf("GetSystemPackageUpdateStatus() error = %v", err)
		}
	default:
		t.Fatalf("unsupported test request %s %s", method, path)
	}
	return recorder
}

func TestGetSystemPackageUpdatesResponse(t *testing.T) {
	original := service.MyService
	t.Cleanup(func() { service.MyService = original })

	fakeSystem := &fakeSystemPackageService{
		updates: service.SystemPackageUpdates{
			Supported: true,
			Manager:   "apt",
			Count:     1,
			Updates: []service.SystemPackageUpdate{{
				Name:             "libc6",
				CurrentVersion:   "1.0",
				CandidateVersion: "1.1",
			}},
		},
	}
	service.MyService = fakeSystemPackageRepository{system: fakeSystem}

	recorder := performSystemPackageRequest(t, http.MethodGet, "/v1/sys/packages")
	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}
	var response model.Result
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if response.Success != http.StatusOK {
		t.Fatalf("response success = %d, want %d", response.Success, http.StatusOK)
	}
	data, ok := response.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("response data type = %T, want object", response.Data)
	}
	if data["count"] != float64(1) {
		t.Fatalf("response data = %#v", data)
	}
}

func TestGetSystemPackageUpdatesReturnsConflictWhenRunning(t *testing.T) {
	original := service.MyService
	t.Cleanup(func() { service.MyService = original })

	fakeSystem := &fakeSystemPackageService{
		updates:  service.SystemPackageUpdates{Supported: true, Manager: "apt"},
		checkErr: service.ErrSystemPackageUpdateRunning,
	}
	service.MyService = fakeSystemPackageRepository{system: fakeSystem}

	recorder := performSystemPackageRequest(t, http.MethodGet, "/v1/sys/packages")
	if recorder.Code != http.StatusConflict {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusConflict)
	}
}

func TestStartSystemPackageUpdateReturnsConflictWhenRunning(t *testing.T) {
	original := service.MyService
	t.Cleanup(func() { service.MyService = original })

	fakeSystem := &fakeSystemPackageService{
		startErr:     service.ErrSystemPackageUpdateRunning,
		updateStatus: service.SystemPackageUpdateStatus{State: "running", Supported: true, Manager: "apt"},
	}
	service.MyService = fakeSystemPackageRepository{system: fakeSystem}

	recorder := performSystemPackageRequest(t, http.MethodPost, "/v1/sys/packages/update")
	if recorder.Code != http.StatusConflict {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusConflict)
	}
	if !errors.Is(fakeSystem.startErr, service.ErrSystemPackageUpdateRunning) {
		t.Fatalf("fake error was changed unexpectedly: %v", fakeSystem.startErr)
	}
}

func TestGetSystemPackageUpdateStatusResponse(t *testing.T) {
	original := service.MyService
	t.Cleanup(func() { service.MyService = original })

	fakeSystem := &fakeSystemPackageService{
		updateStatus: service.SystemPackageUpdateStatus{
			Supported:      true,
			Manager:        "apt",
			State:          "succeeded",
			Log:            "CASAOS_PACKAGE_UPDATE_SUCCESS",
			RebootRequired: true,
		},
	}
	service.MyService = fakeSystemPackageRepository{system: fakeSystem}

	recorder := performSystemPackageRequest(t, http.MethodGet, "/v1/sys/packages/update/status")
	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}
	var response model.Result
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	data, ok := response.Data.(map[string]interface{})
	if !ok || data["state"] != "succeeded" || data["reboot_required"] != true {
		t.Fatalf("response data = %#v", response.Data)
	}
}
