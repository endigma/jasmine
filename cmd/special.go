package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cmd_completion = &cobra.Command{
		Use:   "completion <zsh | fish | bash>",
		Short: "Generate shell completions",
		Long: fmt.Sprintf(`To load completions:

  Bash:
  
    $ source <(%[1]s completion bash)
  
    # To load completions for each session, execute once:
    # Linux:
    $ %[1]s completion bash > /etc/bash_completion.d/%[1]s
    # macOS:
    $ %[1]s completion bash > /usr/local/etc/bash_completion.d/%[1]s
  
  Zsh:
  
    # If shell completion is not already enabled in your environment,
    # you will need to enable it.  You can execute the following once:
  
    $ echo "autoload -U compinit; compinit" >> ~/.zshrc
  
    # To load completions for each session, execute once:
    $ %[1]s completion zsh > "${fpath[1]}/_%[1]s"
  
    # You will need to start a new shell for this setup to take effect.
  
  fish:
  
    $ %[1]s completion fish | source
  
    # To load completions for each session, execute once:
    $ %[1]s completion fish > ~/.config/fish/completions/%[1]s.fish
`, os.Args[0]),
		Args:      cobra.MinimumNArgs(1),
		ValidArgs: []string{"bash", "zsh", "fish"},
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				cmd.Root().GenFishCompletion(os.Stdout, true)
			}
		},
	}
)

func init() {
	cmd_root.AddCommand(cmd_completion)
}
