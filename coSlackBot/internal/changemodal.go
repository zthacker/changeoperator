package coSlackBot

import (
	"fmt"
	"github.com/slack-go/slack"
)

func generateChangeModal(s slack.SlashCommand) slack.ModalViewRequest {
	fmt.Println("Generating modal")
	titleText := slack.NewTextBlockObject("plain_text", "Create a Change", false, false)
	closeText := slack.NewTextBlockObject("plain_text", "Cancel", false, false)
	submitText := slack.NewTextBlockObject("plain_text", "Submit", false, false)

	sectionText := slack.NewTextBlockObject("mrkdwn", "Fill out the fields to create a Change.", false, false)
	sectionBlock := slack.NewSectionBlock(sectionText, nil, nil)

	descriptionText := slack.NewTextBlockObject("plain_text", "Description", false, false)
	descriptionPlaceholder := slack.NewTextBlockObject("plain_text", "Brief description of the change", false, false)
	descriptionElement := slack.PlainTextInputBlockElement{
		Type:         slack.METPlainTextInput,
		ActionID:     "description",
		Placeholder:  descriptionPlaceholder,
		Multiline:    true,
		InitialValue: s.Text,
	}
	descriptionBlock := slack.NewInputBlock("Description", descriptionText, descriptionElement)

	customerImpactLabel := slack.NewTextBlockObject("plain_text", "Could this impact a customer?", false, false)
	customerImpactPlaceholder := slack.NewTextBlockObject("plain_text", "Yes, No or Not Sure?", false, false)
	var customerImpactOptions []*slack.OptionBlockObject
	for c := 1; c <= 3; c++ {
		if c == 1 {
			customerImpactOption := slack.OptionBlockObject{
				Text:  slack.NewTextBlockObject("plain_text", "Yes", false, false),
				Value: "yes",
			}
			customerImpactOptions = append(customerImpactOptions, &customerImpactOption)
		}
		if c == 2 {
			customerImpactOption := slack.OptionBlockObject{
				Text:  slack.NewTextBlockObject("plain_text", "No", false, false),
				Value: "no",
			}
			customerImpactOptions = append(customerImpactOptions, &customerImpactOption)
		}
		if c == 3 {
			customerImpactOption := slack.OptionBlockObject{
				Text:  slack.NewTextBlockObject("plain_text", "Not sure", false, false),
				Value: "not_sure",
			}
			customerImpactOptions = append(customerImpactOptions, &customerImpactOption)
		}

	}
	customerImpactElement := slack.NewOptionsSelectBlockElement("static_select", customerImpactPlaceholder, "customerimpact", customerImpactOptions...)
	customerImpactBlock := slack.NewInputBlock("CustomerImpact", customerImpactLabel, customerImpactElement)

	envLabel := slack.NewTextBlockObject("plain_text", "Select which environment", false, false)
	envPlaceholder := slack.NewTextBlockObject("plain_text", "alpha, prod, etc", false, false)
	var envOptions []*slack.OptionBlockObject
	for i := 0; i <= 3; i++ {
		if i == 0 {
			envOption := slack.OptionBlockObject{
				Text:  slack.NewTextBlockObject("plain_text", "Alpha", false, false),
				Value: "alpha",
			}
			envOptions = append(envOptions, &envOption)
		}
		if i == 1 {
			envOption := slack.OptionBlockObject{
				Text:  slack.NewTextBlockObject("plain_text", "Prod", false, false),
				Value: "prod",
			}
			envOptions = append(envOptions, &envOption)
		}
		if i == 2 {
			envOption := slack.OptionBlockObject{
				Text:  slack.NewTextBlockObject("plain_text", "Dev", false, false),
				Value: "dev",
			}
			envOptions = append(envOptions, &envOption)
		}
		if i == 3 {
			envOption := slack.OptionBlockObject{
				Text:  slack.NewTextBlockObject("plain_text", "Lab", false, false),
				Value: "lab",
			}
			envOptions = append(envOptions, &envOption)
		}
	}
	envElement := slack.NewOptionsSelectBlockElement("static_select", envPlaceholder, "environment", envOptions...)
	envBlock := slack.NewInputBlock("Environment", envLabel, envElement)

	typeLabel := slack.NewTextBlockObject("plain_text", "Select the Type of Change", false, false)
	typePlaceholder := slack.NewTextBlockObject("plain_text", "network, maintenance, migration, etc", false, false)
	var typeOptions []*slack.OptionBlockObject
	for i := 0; i <= 16; i++ {
		if i == 0 {
			typeOption := slack.OptionBlockObject{
				Text:  slack.NewTextBlockObject("plain_text", "Maintenance", false, false),
				Value: "maintenance",
			}
			typeOptions = append(typeOptions, &typeOption)
		}
		if i == 1 {
			typeOption := slack.OptionBlockObject{
				Text:  slack.NewTextBlockObject("plain_text", "Client", false, false),
				Value: "client",
			}
			typeOptions = append(typeOptions, &typeOption)
		}
		if i == 2 {
			typeOption := slack.OptionBlockObject{
				Text:  slack.NewTextBlockObject("plain_text", "Deploy", false, false),
				Value: "deploy",
			}
			typeOptions = append(typeOptions, &typeOption)
		}
	}
	typeElement := slack.NewOptionsSelectBlockElement("static_select", typePlaceholder, "changetype", typeOptions...)
	typeBlock := slack.NewInputBlock("Changetype", typeLabel, typeElement)

	linkText := slack.NewTextBlockObject("plain_text", "Relevant Link", false, false)
	linkPlaceholder := slack.NewTextBlockObject("plain_text", "Any relevant link?", false, false)
	linkElement := slack.PlainTextInputBlockElement{
		Type:         slack.METPlainTextInput,
		ActionID:     "link",
		Placeholder:  linkPlaceholder,
		Multiline:    false,
		InitialValue: s.Text,
	}
	linkBlock := slack.NewInputBlock("Link", linkText, linkElement)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			sectionBlock,
			customerImpactBlock,
			envBlock,
			typeBlock,
			descriptionBlock,
			linkBlock,
		},
	}

	return slack.ModalViewRequest{
		Type:       slack.ViewType("modal"),
		Title:      titleText,
		Close:      closeText,
		Submit:     submitText,
		Blocks:     blocks,
		CallbackID: "change-modal",
	}
}
