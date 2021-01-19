// FIXME remove this file

package cmd

import "github.com/spf13/cobra"

func (c *Config) newReadLineCmd() *cobra.Command {
	readLineCmd := &cobra.Command{
		Use:    "read-line",
		Args:   cobra.NoArgs,
		Hidden: true,
		RunE:   c.runReadLineCmdE,
	}
	return readLineCmd
}

func (c *Config) runReadLineCmdE(cmd *cobra.Command, args []string) error {
	line, err := c.readLine("Line: ")
	if err != nil {
		return err
	}
	return c.writeOutputString(line)
}
