package zing

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func ZingCommand(service *Services) *cobra.Command {
	zingCmd := &cobra.Command{
		Use:   "zing",
		Short: "zing through repetitive commands",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	zingCmd.AddCommand(listCommands(service), addCommand(service), runCommands(service), previewCommand(service))

	return zingCmd
}

func addCommand(service *Services) *cobra.Command {
	var tag string
	var action string
	add := &cobra.Command{
		Use:     "add [command pieces...]",
		Short:   "add a new stored command (pass the full shell command as args)",
		Example: "zing add --tag deploy --cmd docker compose build && docker compose push && kubectl rollout restart deploy/{{ .service }} -n {{ .ns }}",
		Args:    cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			// join args into a single command string
			full := strings.Join(args, " ")
			LogInfo("adding: ", full, " tag: ", tag)
			LogInfo("cmd: ", action)

			// call your service to persist the command (uncomment and set correct import)
			prompt, err := service.SetCommand(tag, action)
			if err != nil {
				LogError(err)
				return nil
			}

			if prompt {
				reader := bufio.NewReader(os.Stdin)
				for {
					LogMessage("You are trying to add a tag that already exist.")
					LogNormal("Proceed? [y/N]: ")
					input, _ := reader.ReadString('\n')
					input = strings.TrimSpace(strings.ToLower(input))

					if input == "" || input == "n" || input == "no" {
						LogNormalLn("Aborted.")
						return nil
					}
					if input == "y" || input == "yes" {
						err := service.UpdateCommand(tag, action)
						if err != nil {
							LogError(err)
							return nil
						}
						msg := fmt.Sprintf("You have updated the zing for %s", tag)
						LogSuccess(msg)
						break
					}
					LogError("Please answer yes or no.")
				}
			}

			return nil
		},
	}

	add.Flags().StringVarP(&tag, "tag", "n", "default", "name of the command")
	add.Flags().StringVarP(&action, "cmd", "c", "default", "the command to store")

	return add
}

func listCommands(service *Services) *cobra.Command {
	list := &cobra.Command{
		Use:   "list",
		Short: "use this to list all available commands stored",
		RunE: func(cmd *cobra.Command, args []string) error {
			commands, err := service.ListCommands()
			if err != nil {
				LogError(err)
			}
			LogInfo(commands)
			return nil
		},
	}

	return list
}

func runCommands(service *Services) *cobra.Command {

	run := &cobra.Command{
		Use:     "run",
		Short:   "use this to run a stored command",
		Example: "zing run <tag>",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			tag := args[0]
			command, err := service.GetCommand(tag)
			if err != nil {
				LogError(err)
				return nil
			}
			variables := GetVariables(command)
			fields := GetFieldsMap(command)

			full := slices.Delete(args, 0, 1)

			argMap := GetVariablesMap(full)

			newCommand := command

			for _, variable := range variables {
				va, ok := argMap[variable]
				if !ok {
					reader := bufio.NewReader(os.Stdin)
					LogMessage(fmt.Sprintf("You did not pass %s variable", variable))
					LogNormal(fmt.Sprintf("Enter variable name for %s: ", variable))
					input, _ := reader.ReadString('\n')
					input = strings.TrimSpace(strings.ToLower(input))
					va = input
				}
				field, ok := fields[variable]
				if !ok {
					LogError("Failed to find filed")
					return nil
				}
				newCommand = strings.ReplaceAll(newCommand, field.(string), va.(string))
			}

			LogInfo(fmt.Sprintf("Running command \"%s\" ...", newCommand))
			// run a shell here in terminal
			resp, err := StreamShell(newCommand, time.Minute*30)
			if err != nil {
				LogError(err)
			}
			LogNormalLn(resp)
			return nil
		},
	}
	return run
}

func previewCommand(service *Services) *cobra.Command {
	preview := &cobra.Command{
		Use:   "preview",
		Short: "Use this to preview all a command and see the args required",
		Args:  cobra.RangeArgs(1, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			tag := args[0]
			command, err := service.GetCommand(tag)
			if err != nil {
				LogError(err)
			}
			variables := len(GetVariables(command))

			LogInfo(command)
			LogMessage(fmt.Sprintf("There are %d variables needed for this command", variables))
			return nil
		},
	}

	return preview
}
