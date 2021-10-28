/*
	Copyright 2019 NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package main

import (
	"github.com/openziti/fabric/router"
	"github.com/openziti/fabric/router/xgress"
	"github.com/pkg/errors"
)

func Initialize(*router.Router) error {
	xgress.GlobalRegistry().Register("geneve", Factory{})
	return nil
}

type Factory struct{}

func (f Factory) CreateListener(optionsData xgress.OptionsData) (xgress.Listener, error) {
	return &listener{}, nil
}

func (f Factory) CreateDialer(optionsData xgress.OptionsData) (xgress.Dialer, error) {
	return nil, errors.New("dialer not supported")
}
