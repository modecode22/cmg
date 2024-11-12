package commit

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

func RunCLI() error {
	color.Cyan("Welcome to the Conventional Commit CLI!")

	c := &Commit{}
	var typeOptions []string
	for _, t := range DefaultTypes() {
		typeOptions = append(typeOptions, 
			fmt.Sprintf("%s: %s %s", 
				t.Label, t.Description, t.Emoji))
	}

	questions := []*survey.Question{
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "Select the type of change you're committing:",
				Options: typeOptions,
			},
			Validate: survey.Required,
		},
		{
			Name: "subject",
			Prompt: &survey.Input{
				Message: "Write a short, imperative tense description of the change:",
			},
			Validate: survey.Required,
		},
		{
			Name: "scope",
			Prompt: &survey.Input{
				Message: "What is the scope of this change (e.g. component or file name): (press enter to skip)",
			},
		},
		{
			Name: "body",
			Prompt: &survey.Input{
				Message: "Provide a longer description of the change: (press enter to skip)",
			},
		},
		{
			Name: "breaking",
			Prompt: &survey.Input{
				Message: "List any breaking changes: (press enter to skip)",
			},
		},
		{
			Name: "issues",
			Prompt: &survey.Input{
				Message: "List any issues closed by this change: (press enter to skip)",
			},
		},
	}

	err := survey.Ask(questions, c)
	if err != nil {
		return fmt.Errorf("error occurred while prompting: %w", err)
	}
	commitMsg := formatCommitMessage(c)
	fmt.Println(color.CyanString("\nYour commit message:"))
	fmt.Println(color.YellowString(commitMsg))

	return commitChanges(commitMsg)
}

func formatCommitMessage(c *Commit) string {
	// Split by ":" and take first part, which includes emoji and type
	typeWithEmoji := strings.TrimSpace(strings.Split(c.Type, ":")[0])
	
	commitMsg := fmt.Sprintf("%s ", typeWithEmoji)
	if c.Scope != "" {
		commitMsg += fmt.Sprintf("(%s)", c.Scope)
	}
	commitMsg += fmt.Sprintf(": %s", c.Subject)
	if c.Body != "" {
		commitMsg += fmt.Sprintf("\n\n%s", c.Body)
	}
	if c.Breaking != "" {
		commitMsg += fmt.Sprintf("\n\nBREAKING CHANGE: %s", c.Breaking)
	}
	if c.Issues != "" {
		commitMsg += fmt.Sprintf("\n\nCloses %s", c.Issues)
	}
	return commitMsg
}

func commitChanges(commitMsg string) error {
	// Check if there are staged changes
	stagedChanges, err := exec.Command("git", "diff", "--cached", "--name-only").Output()
	if err != nil {
		return fmt.Errorf("error checking for staged changes: %w", err)
	}

	if len(stagedChanges) == 0 {
		// No staged changes, ask if user wants to commit all changes
		commitAll := false
		prompt := &survey.Confirm{
			Message: "No staged changes. Do you want to commit all changes?",
			Default: true,
		}
		survey.AskOne(prompt, &commitAll)
		if !commitAll {
			fmt.Println(color.YellowString("Commit cancelled. Please stage your changes and try again."))
			return nil
		}
		// Stage all changes
		gitAdd := exec.Command("git", "add", ".")
		gitAdd.Stdout = os.Stdout
		gitAdd.Stderr = os.Stderr
		err = gitAdd.Run()
		if err != nil {
			return fmt.Errorf("error occurred while staging changes: %w", err)
		}
	}

	// Commit the changes
	gitCommit := exec.Command("git", "commit", "-m", commitMsg)
	gitCommit.Stdout = os.Stdout
	gitCommit.Stderr = os.Stderr
	err = gitCommit.Run()
	if err != nil {
		return fmt.Errorf("error occurred while committing: %w", err)
	}

	fmt.Println(color.GreenString("Successfully committed!"))
	return nil
}