package setup

import (
	"fmt"
	"strconv"

	"github.com/OnFinality-io/onf-cli/pkg/service"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "initialize the configuration",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			profile, err := cmd.Flags().GetString("profile")
			if err != nil {
				fmt.Printf("err:%s\n", err)
				return
			}
			Flow(profile)
		},
	}
	cmd.AddCommand(
		listCmd(),
	)
	return cmd
}

func Flow(section string) {

	credential := &Credential{}

	// access key
	accessKeyPrompt := promptui.Prompt{
		Label: "Please input your access key",
	}
	result, err := accessKeyPrompt.Run()
	if err != nil {
		fmt.Printf("Fail to add access key %v\n", err)
		return
	}
	credential.AccessKey = result

	// secret key
	secretKeyPrompt := promptui.Prompt{
		Label: "Please input your secret key",
	}
	result, err = secretKeyPrompt.Run()
	if err != nil {
		fmt.Printf("Fail to add secret key %v\n", err)
		return
	}
	credential.SecretKey = result

	// workspace id key
	service.Init(credential.AccessKey, credential.SecretKey)
	list, err := service.GetWorkspaceList()
	if len(list) == 1 {
		credential.WorkspaceID = list[0].ID
	} else if len(list) > 0 {
		var name []string
		var tmp string
		for _, ws := range list {
			tmp = ws.Name + "(" + strconv.FormatUint(ws.ID, 10) + ")"
			name = append(name, tmp)
		}
		workspaceIDPrompt := promptui.Select{
			Label: "Please select your workspace",
			Items: name,
		}
		index, _, err := workspaceIDPrompt.Run()
		if err != nil {
			fmt.Printf("Fail to add workspace id %v\n", err)
			return
		}
		credential.WorkspaceID = list[index].ID
	}

	config := &CredentialConfig{
		Credential: credential,
		Section:    section,
	}
	PersistentCredential(config)
}
