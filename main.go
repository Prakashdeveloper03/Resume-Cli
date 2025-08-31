// Package main is the entry point for the portfolio CLI application.
// It initializes the user interface, handles user input through a menu,
// and calls functions to display different sections of the portfolio.
package main

import (
	data "resumecli/models"
	"resumecli/ui"
	"resumecli/utils"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

// main is the primary function that runs the CLI application.
func main() {
	// Ensure required data directories exist
	utils.EnsureDirs([]string{
		"./data/about", "./data/experience", "./data/education",
		"./data/skills", "./data/projects", "./data/certificates", "./data/contact",
	})

	// Welcome the user
	utils.CallClear()
	ui.ShowIntro()

	// Main menu options
	options := []string{
		"👤 About",
		"💼 Experience",
		"🎓 Education",
		"✨ Skills",
		"🚀 Projects",
		"📜 Certificates",
		"📧 Contact",
		"🚪 Exit",
	}

	// Main interaction loop
	for {
		utils.CallClear()

		var choice string
		prompt := &survey.Select{
			Message: "\nWhat would you like to know about me?",
			Options: options,
		}

		err := survey.AskOne(prompt, &choice, survey.WithIcons(func(icons *survey.IconSet) {
			icons.Question.Format = "cyan+b"
			icons.SelectFocus.Text = "›"
			icons.SelectFocus.Format = "yellow+b"
		}))

		// If Ctrl+C or other input error, exit gracefully
		if err != nil {
			ui.ShowExitMessage()
			return
		}

		utils.CallClear()

		var loadErr error

		// Handle each menu option
		switch choice {
		case "👤 About":
			var d data.AboutData
			loadErr = utils.LoadJSON("./data/about/about.json", &d)
			if loadErr == nil {
				ui.ShowAbout(d)
			}
		case "💼 Experience":
			var d data.ExperienceData
			loadErr = utils.LoadJSON("./data/experience/experience.json", &d)
			if loadErr == nil {
				ui.ShowExperience(d)
			}
		case "🎓 Education":
			var d data.EducationData
			loadErr = utils.LoadJSON("./data/education/education.json", &d)
			if loadErr == nil {
				ui.ShowEducation(d)
			}
		case "✨ Skills":
			var d data.SkillsData
			loadErr = utils.LoadJSON("./data/skills/skills.json", &d)
			if loadErr == nil {
				ui.ShowSkills(d)
			}
		case "🚀 Projects":
			var d data.ProjectsData
			loadErr = utils.LoadJSON("./data/projects/projects.json", &d)
			if loadErr == nil {
				ui.ShowProjects(d)
			}
		case "📜 Certificates":
			var d data.CertificatesData
			loadErr = utils.LoadJSON("./data/certificates/certificates.json", &d)
			if loadErr == nil {
				ui.ShowCertificates(d)
			}
		case "📧 Contact":
			var d data.ContactData
			loadErr = utils.LoadJSON("./data/contact/contact.json", &d)
			if loadErr == nil {
				ui.ShowContact(d)
			}
		case "🚪 Exit":
			ui.ShowExitMessage()
			return
		}

		// Show error if file failed to load
		if loadErr != nil {
			color.Red("Could not display section due to a file loading error: %v", loadErr)
		}

		// Wait before returning to menu
		ui.PromptToContinue()
	}
}
