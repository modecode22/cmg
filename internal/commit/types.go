package commit

// Commit represents a conventional commit message structure
type Commit struct {
	Type     string
	Scope    string
	Subject  string
	Body     string
	Breaking string
	Issues   string
}

// TypeOption represents a commit type with its description
type TypeOption struct {
	Label       string
	Description string
	Emoji      string
}

// DefaultTypes returns the default commit types
func DefaultTypes() []TypeOption {
	return []TypeOption{
		{"feat", "A new feature", "✨"},
		{"fix", "A bug fix", "🐛"},
		{"docs", "Documentation only changes", "📚"},
		{"style", "Changes that do not affect the meaning of the code", "🎨"},
		{"refactor", "A code change that neither fixes a bug nor adds a feature", "🔨"},
		{"perf", "A code change that improves performance", "⚡️"},
		{"test", "Adding missing tests or correcting existing tests", "✅"},
		{"build", "Changes that affect the build system or external dependencies", "🛠️"},
		{"ci", "Changes to our CI configuration files and scripts", "👷"},
		{"chore", "Other changes that don't modify src or test files", "🧹"},
		{"revert", "Reverts a previous commit", "⏪"},
		{"init", "Initial commit", "🎉"},
	}
}