// Copyright © 2019 BigOokie
//
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package goappinstancectrl

import (
	"github.com/marcsauter/single"
	"github.com/pkg/errors"
)

// InitAppInstance initialises an instance of the application based on the value of `appID`.
// An error will be returned if another instance of the application instance is
// detected (based on the `appID`). You should handle this error as appropriate for your situation.
// Call this when your application starts.
func InitAppInstance(appID string) (s *single.Single, err error) {
	// Check that an appID has been provided
	if appID == "" {
		return nil, errors.New("appID must be set")
	}

	s = single.New(appID)
	if err := s.CheckLock(); err != nil && err == single.ErrAlreadyRunning {
		// An instance of appID is already running.
		return nil, errors.Wrapf(err, "another instance of the application (appID = %s) is already running", appID)
	} else if err != nil {
		// An unknown error occurred
		return nil, errors.Wrapf(err, "an error occurred attempting to obtain an application instance lock (appID = %s)", appID)
	}
	return s, nil
}

// ReleaseAppInstance will release(unlock) an instance of the application.
// An error will be returned if this should fail for any reason. You should handle the error as appropriate for you situation.
// Call this when your application is exiting.
func ReleaseAppInstance(s *single.Single) (err error) {
	if s == nil {
		return errors.New("an application instance (single.Single) was not specified")
	}

	// Try to unlock the application instance
	err = s.TryUnlock()
	if err != nil {
		return errors.Wrap(err, "an error occurred attempting to release the application instance lock")
	}

	return nil
}
