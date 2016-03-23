package volume

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	inspectCmd = &cobra.Command{
		Use:   "inspect <volume ID>",
		Short: "Inspect a volume",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("volume ID missing")
			}

			fmt.Printf("Volume ID = %s\n", args[0])

			// TODO(amitshukla): Send it to the Manager thru grpc

			return nil
		},
	}
)
