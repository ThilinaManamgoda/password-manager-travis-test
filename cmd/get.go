// Copyright © 2019 Thilina Manamgoda
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


package cmd

import (
	"github.com/ThilinaManamgoda/password-manager/pkg/inputs"
	"github.com/ThilinaManamgoda/password-manager/pkg/passwords"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// ShowPassword flag
const (
	// ShowPassword flag
	ShowPassword = "show-pass"
	// ErrMSGCannotGetFlag message
	ErrMSGCannotGetFlag = "cannot get value of %s flag"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [ID]",
	Short: "Get a password",
	Long:  `Get a password`,
	Args:  inputs.HasProvidedValidID(),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		mPassword, err := inputs.GetFlagStringVal(cmd, MasterPassword)
		if err != nil {
			return errors.Wrapf(err, ErrMSGCannotGetFlag, mPassword)
		}
		if mPassword == "" {
			mPassword, err = promptForMPassword()
			if err != nil {
				return errors.Wrap(err, "cannot prompt for Master password")
			}
		}
		showPass, err := inputs.GetFlagBoolVal(cmd, ShowPassword)
		if err != nil {
			return errors.Wrapf(err, ErrMSGCannotGetFlag, Password)
		}

		passwordRepo, err := passwords.InitPasswordRepo(mPassword)
		if err != nil {
			return errors.Wrapf(err, "cannot initialize password repository")
		}

		err = passwordRepo.GetPassword(id, showPass)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolP(ShowPassword, "s", false, "Print password to STDOUT")
}
