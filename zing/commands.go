package zing

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func ZingCommand(service *Services) *cobra.Command {
	zingCmd := &cobra.Command{
		Use:   "zing",
		Short: "zing through repetitive commands",
	}

	zingCmd.AddCommand(listCommands(service), addCommand(service), runCommands(service))

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
				return err
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
							return err
						}
						msg := fmt.Sprintf("You have updated command for %s", tag)
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
				panic(err)
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
		Args:    cobra.RangeArgs(1, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			tag := strings.Join(args, " ")
			command, err := service.RunCommand(tag)
			if err != nil {
				LogError(err)
				return nil
			}
			LogInfo(fmt.Sprintf("Running command \"%s\" ...", command))
			// run a shell here in terminal
			resp, err := StreamShell(command, time.Minute*30)
			if err != nil {
				LogError(err)
			}
			LogNormalLn(resp)
			return nil
		},
	}
	return run
}
