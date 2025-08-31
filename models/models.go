package data

// AboutData holds the content for the "About" section of a resume or portfolio.
type AboutData struct {
	About string `json:"About"` // A brief description or bio.
}

// ContactData contains various contact details and professional profiles.
type ContactData struct {
	Email            string `json:"Email"`             // Email address.
	Mobile           string `json:"Mobile"`            // Mobile or phone number.
	LinkedIn         string `json:"LinkedIn"`          // LinkedIn profile URL.
	Hackerrank       string `json:"Hackerrank"`        // HackerRank profile URL.
	Leetcode         string `json:"Leetcode"`          // LeetCode profile URL.
	GeeksforGeeks    string `json:"GeeksforGeeks"`     // GeeksforGeeks profile URL.
	Github           string `json:"Github"`            // GitHub profile URL.
	PortfolioWebsite string `json:"Portfolio Website"` // Personal portfolio website URL.
}

// Experience represents a single work experience entry.
type Experience struct {
	Position    string   `json:"Position"`    // Job title or role.
	Company     string   `json:"Company"`     // Company or organization name.
	Location    string   `json:"Location"`    // Location of the job (remote or on-site).
	Duration    string   `json:"Duration"`    // Time period of employment.
	Description []string `json:"Description"` // Key responsibilities or achievements.
}

// ExperienceData holds a list of professional experiences.
type ExperienceData struct {
	Experiences []Experience `json:"Experiences"` // Slice of work experience entries.
}

// EducationEntry represents a single educational qualification.
type EducationEntry struct {
	Degree      string `json:"Degree"`      // Name of the degree or program.
	Institution string `json:"Institution"` // Name of the educational institution.
	Duration    string `json:"Duration"`    // Time period of study.
}

// EducationData holds a list of education entries.
type EducationData struct {
	Education []EducationEntry `json:"Education"` // Slice of educational qualifications.
}

// SkillsData categorizes different technical skills.
type SkillsData struct {
	Languages string `json:"Languages"` // Programming languages known.
	Frontend  string `json:"Frontend"`  // Frontend technologies and frameworks.
	Backend   string `json:"Backend"`   // Backend technologies and frameworks.
	Databases string `json:"Databases"` // Databases and data storage technologies.
	Cloud     string `json:"Cloud"`     // Cloud platforms and services.
	AIML      string `json:"AI/ML"`     // AI/ML tools, frameworks, or knowledge areas.
	Workspace string `json:"Workspace"` // Tools used in a development or team environment.
}

// Project represents a single project entry with relevant links.
type Project struct {
	Title           string   `json:"Title"`            // Title of the project.
	Description     []string `json:"Description"`      // Overview or key features of the project.
	SourceCode      string   `json:"Source Code"`      // Link to the source code repository.
	ApplicationLink string   `json:"Application Link"` // Link to the live application (if any).
}

// ProjectsData holds a list of personal or professional projects.
type ProjectsData struct {
	Projects []Project `json:"Projects"` // Slice of project entries.
}

// CertificatesData holds a list of certifications or achievements.
type CertificatesData struct {
	Certificates []string `json:"Certificates"` // Slice of certificate titles or descriptions.
}
