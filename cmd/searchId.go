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

// searchIDCmd represents the searchId command
var searchIDCmd = &cobra.Command{
	Use:   "searchId [ID]",
	Short: "Search Password with ID",
	Long:  `You can use either complete or part of ID for searching`,
	Args:  inputs.HasProvidedValidID(),
	RunE: func(cmd *cobra.Command, args []string) error {
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

		if !inputs.IsValidSingleArg(args) {
			return errors.New("Please give a ID")
		}

		passwordRepo, err := passwords.InitPasswordRepo(mPassword)
		if err != nil {
			return errors.Wrapf(err, "cannot initialize password repository")
		}

		searchID := args[0]
		passwordEntries, err := passwordRepo.SearchID(searchID, showPass)
		if err != nil {
			return errors.Wrapf(err, "cannot search ID")
		}

		if len(passwordEntries) != 0 {
			var idList []string
			for _, val := range passwordEntries {
				idList = append(idList, val.ID)
			}
			sID, _ := inputs.PromptForSelect("Choose", idList)
			err := passwordRepo.GetPassword(sID, showPass)
			if err != nil {
				return errors.Wrapf(err, "cannot get password for ID: %s", sID)
			}
		} else {
			return errors.New("cannot find any match")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchIDCmd)
	searchIDCmd.Flags().BoolP(ShowPassword, "s", false, "Print password to STDOUT")
}