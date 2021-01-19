// FIXME remove this file

package cmd

import "github.com/spf13/cobra"

func (c *Config) newReadPasswordCmd() *cobra.Command {
	readPasswordCmd := &cobra.Command{
		Use:    "read-password",
		Args:   cobra.NoArgs,
		Hidden: true,
		RunE:   c.runReadPasswordCmdE,
	}
	return readPasswordCmd
}

func (c *Config) runReadPasswordCmdE(cmd *cobra.Command, args []string) error {
	password, err := c.readPassword("Password: ")
	if err != nil {
		return err
	}
	return c.writeOutputString(password)
}
