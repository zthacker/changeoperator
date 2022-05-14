package coSlackBot

func RunSlackCommands() {
	for {
		select {
		case slashCommand := <-CommandChannel:
			if slashCommand.Command == "/change" {
				GenerateChangeModal(slashCommand)
			}
		}
	}
}

func RunSlackInteractive() {
	for {
		select {
		case slackInteractive := <-InteractiveChannel:
			if slackInteractive.Type == "view_submission" {
				if slackInteractive.View.CallbackID == "change-modal" {
					HandleCreateChange(slackInteractive)
				}
			}
		}
	}
}
