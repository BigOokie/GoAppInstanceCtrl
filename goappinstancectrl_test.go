// Copyright © 2019 BigOokie
//
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.
package goappinstancectrl

import (
	"testing"
)

func Test_InitAppInstance_InvalidAppID(t *testing.T) {
	// Test that we get an error if the appID parameter is empty
	_, err := InitAppInstance("")
	if err == nil {
		t.Error("An error is expected if InitAppInstace is called with an empty appID parameter")
	}
}

func Test_InitAppInstance_ValidAppID(t *testing.T) {
	appID := "Test_InitAppInstance_ValidAppID"
	s, err := InitAppInstance(appID)
	if err != nil {
		t.Errorf("An unexpected error occurred calling InitAppInstance with appID = %s. Error = %s", appID, err)
	}

	if s == nil {
		t.Error("A Single struct is expected to be returned if InitAppInstace returns without errors")
	}

	s.Unlock()
}

func Test_InitAppInstance_FailOnSameAppID(t *testing.T) {
	appID := "Test_InitAppInstance_FailOnSameAppID"
	s, err := InitAppInstance(appID)
	if err != nil {
		t.Errorf("An unexpected error occurred calling InitAppInstance with appID = %s. Error = %s", appID, err)
	}

	if s == nil {
		t.Error("A Single struct is expected to be returned if InitAppInstace returns without errors")
	}

	// Attempt to init the app instance for the same appID
	_, err = InitAppInstance(appID)
	if err != nil {
		t.Error("An error is expected after calling InitAppInstance a second time with the same appID")
	}

	s.Unlock()
}

func Test_ReleaseAppInstance_InvalidAppID(t *testing.T) {
	// Test that we get an error if the appID parameter is empty
	err := ReleaseAppInstance(nil)
	if err == nil {
		t.Error("An error is expected if ReleaseAppInstance is called with an unassigned Single struct")
	}
}

func Test_ReleaseAppInstance_ValidAppID(t *testing.T) {
	appID := "Test_ReleaseAppInstance_ValidAppID"
	s, err := InitAppInstance(appID)
	if (err != nil) || (s == nil) {
		t.Errorf("An unexpected error occurred calling InitAppInstance with appID = %s. Error = %s", appID, err)
	}

	err = ReleaseAppInstance(s)
	if err != nil {
		t.Errorf("An unexpected error occurred calling ReleaseAppInstance. Error = %s", err)
	}
}
