package main

import "github.com/DataDrake/cli-ng/v2/cmd"

func main() {
	root := &cmd.Root{
		Name:  "minecraftcp-cli",
		Short: "Terminal control panel for Minecraft servers",
	}

	cmd.Register(&cmd.Sub{
		Name:  "run",
		Short: "Starts the server control panel",
		Run:   Run,
	})

	root.Run()
}
