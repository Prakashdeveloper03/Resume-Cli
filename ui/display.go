package ui

import (
	"fmt"
	data "resumecli/models"
	"strings"

	"github.com/fatih/color"
)

// ShowAbout displays the "About Me" section with word-wrapping.
func ShowAbout(d data.AboutData) {
	printSectionHeader("About Me")
	aboutColor := color.New(color.FgWhite)
	words := strings.Fields(d.About)
	var line string

	// Word wrapping based on terminal width
	for _, word := range words {
		if len(line)+len(word)+1 > TerminalWidth-4 {
			_, err := aboutColor.Println("  " + line)
			if err != nil {
				return
			}
			line = ""
		}
		line += word + " "
	}
	_, err := aboutColor.Println("  " + line)
	if err != nil {
		return
	}
}

// ShowContact displays the "Contact Information" section with links and styling.
func ShowContact(d data.ContactData) {
	printSectionHeader("Contact Information")

	labelColor := color.New(color.FgHiCyan)                    // Label headings
	valueColor := color.New(color.FgWhite)                     // Text values
	linkColor := color.New(color.FgHiMagenta, color.Underline) // Hyperlinks

	// Email with clickable mailto link
	emailLink := createLink(d.Email, "mailto:"+d.Email)
	fmt.Printf("  %-20s : ", labelColor.Sprint("📧 Email"))
	_, err := valueColor.Println(emailLink)
	if err != nil {
		return
	}

	// Mobile
	fmt.Printf("  %-20s : ", labelColor.Sprint("📱 Mobile"))
	_, err = valueColor.Println(d.Mobile)
	if err != nil {
		return
	}
	fmt.Println()

	// Professional Links
	fmt.Printf("  %-20s : %s\n", labelColor.Sprint("💼 LinkedIn"), createLink(linkColor.Sprint(d.LinkedIn), d.LinkedIn))
	fmt.Printf("  %-20s : %s\n", labelColor.Sprint("🐱 Github"), createLink(linkColor.Sprint(d.Github), d.Github))
	fmt.Printf("  %-20s : %s\n", labelColor.Sprint("🌐 Portfolio"), createLink(linkColor.Sprint(d.PortfolioWebsite), d.PortfolioWebsite))
	fmt.Println()

	// Coding Platforms
	fmt.Printf("  %-20s : %s\n", labelColor.Sprint("💻 Hackerrank"), createLink(linkColor.Sprint(d.Hackerrank), d.Hackerrank))
	fmt.Printf("  %-20s : %s\n", labelColor.Sprint("💡 Leetcode"), createLink(linkColor.Sprint(d.Leetcode), d.Leetcode))
	fmt.Printf("  %-20s : %s\n", labelColor.Sprint("📖 GeeksforGeeks"), createLink(linkColor.Sprint(d.GeeksforGeeks), d.GeeksforGeeks))
}

// ShowExperience displays the "Work Experience" section in a readable layout.
func ShowExperience(d data.ExperienceData) {
	printSectionHeader("Work Experience")

	positionColor := color.New(color.FgHiCyan, color.Bold) // Job title
	metaColor := color.New(color.FgYellow)                 // Company/location/duration
	descColor := color.New(color.FgWhite)                  // Description text
	bulletColor := color.New(color.FgMagenta)              // Bullet symbol

	for i, exp := range d.Experiences {
		_, err := positionColor.Printf("  %s\n", exp.Position)
		if err != nil {
			return
		}
		_, err = metaColor.Printf("    %s │ %s │ %s\n", exp.Company, exp.Location, exp.Duration)
		if err != nil {
			return
		}
		fmt.Println()

		for _, desc := range exp.Description {
			fmt.Printf("      %s %s\n", bulletColor.Sprint("›"), descColor.Sprint(desc))
		}

		// Separator line between experiences
		if i < len(d.Experiences)-1 {
			_, err := color.New(color.FgBlack, color.Bold).Println("\n" + strings.Repeat("· ", (TerminalWidth-2)/2))
			if err != nil {
				return
			}
		}
	}
}

// ShowEducation displays the "Education" section with styling.
func ShowEducation(d data.EducationData) {
	printSectionHeader("Education")

	degreeColor := color.New(color.FgHiYellow, color.Bold) // Degree name
	institutionColor := color.New(color.FgHiGreen)         // Institution
	durationColor := color.New(color.FgWhite)              // Duration

	for _, entry := range d.Education {
		fmt.Printf("  🎓 %s\n", degreeColor.Sprint(entry.Degree))
		fmt.Printf("     %s\n", institutionColor.Sprint(entry.Institution))
		fmt.Printf("     %s\n", durationColor.Sprint(entry.Duration))
		fmt.Println()
	}
}

// ShowSkills displays the "Technical Skills" section categorized by domain.
func ShowSkills(d data.SkillsData) {
	printSectionHeader("Technical Skills")

	keyColor := color.New(color.FgHiCyan, color.Bold) // Skill category title
	bracketColor := color.New(color.FgBlue)           // Brackets
	skillTextColor := color.New(color.FgWhite)        // Skill names

	skillCategories := map[string]string{
		"👨‍💻 Languages": d.Languages,
		"🌐 Frontend":    d.Frontend,
		"⚙️  Backend":   d.Backend,
		"🛢️  Databases": d.Databases,
		"☁️ Cloud":      d.Cloud,
		"🧠 AI/ML":       d.AIML,
		"🛠️  Workspace": d.Workspace,
	}

	for category, skillsStr := range skillCategories {
		fmt.Printf("  %-22s : ", keyColor.Sprint(category))
		skills := strings.Split(skillsStr, ", ")
		for _, skill := range skills {
			fmt.Printf("%s%s%s ",
				bracketColor.Sprint("["),
				skillTextColor.Sprint(skill),
				bracketColor.Sprint("]"),
			)
		}
		fmt.Printf("\n\n")
	}
}

// ShowProjects displays the "Projects" section with descriptions and links.
func ShowProjects(d data.ProjectsData) {
	printSectionHeader("Projects")

	titleColor := color.New(color.FgHiCyan, color.Bold)        // Project title
	descColor := color.New(color.FgWhite)                      // Description
	linkColor := color.New(color.FgHiMagenta, color.Underline) // Links
	bulletColor := color.New(color.FgMagenta)                  // Bullet symbol
	tagColor := color.New(color.FgBlue)                        // Source/live tag

	for i, project := range d.Projects {
		_, err := titleColor.Printf("  %s\n", project.Title)
		if err != nil {
			return
		}
		fmt.Println()

		for _, desc := range project.Description {
			fmt.Printf("      %s %s\n", bulletColor.Sprint("›"), descColor.Sprint(desc))
		}
		fmt.Println()

		// Source code and live link
		fmt.Printf("        %s %s\n", tagColor.Sprint("└─> Source:"), createLink(linkColor.Sprint(project.SourceCode), project.SourceCode))
		fmt.Printf("        %s %s\n", tagColor.Sprint("└─> Live:  "), createLink(linkColor.Sprint(project.ApplicationLink), project.ApplicationLink))

		// Separator between projects
		if i < len(d.Projects)-1 {
			_, err := color.New(color.FgBlack, color.Bold).Println("\n" + strings.Repeat("· ", (TerminalWidth-2)/2))
			if err != nil {
				return
			}
		}
	}
}

// ShowCertificates displays the "Certificates" section.
// Assumes each certificate string is formatted as "Certificate Name - Issuer".
func ShowCertificates(d data.CertificatesData) {
	printSectionHeader("Certificates")

	nameColor := color.New(color.FgHiWhite)    // Certificate title
	issuerColor := color.New(color.FgHiYellow) // Issuer name

	for _, certificate := range d.Certificates {
		parts := strings.SplitN(certificate, "-", 2)
		if len(parts) == 2 {
			fmt.Printf("  🏅 %s from %s\n", nameColor.Sprint(strings.TrimSpace(parts[0])), issuerColor.Sprint(strings.TrimSpace(parts[1])))
		} else {
			// Fallback in case the string is not formatted correctly
			fmt.Printf("  🏅 %s\n", nameColor.Sprint(certificate))
		}
	}
}
