package main
import (
	"fmt"
	"log"

	"github.com/xanzy/go-gitlab"
)
import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	seedUserData("data/user_seed_data.json")

	// register routes
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/users", UsersHandler)
	http.HandleFunc("/subscribe", SubscribeHandler)
	http.HandleFunc("/wiki", WikiHandler)

	// serve up static content
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// start the web server
	http.ListenAndServe(":8080", nil)

	// Configuration
	gitlabURL := "https://gitlab.com"                          
	projectID := "pepito/my-test-repo"                      
	branchName := "main"                                        
	commitMessage := "Add new file via API"                     
	filePath := "hello-world.txt"                                   
	fileContent := "Hello World! This is a new file added via the GitLab API." 

	// Create a GitLab client
	git, err := gitlab.NewClient("glpat-Pi-lkkztP23F8LssgLOV", gitlab.WithBaseURL(gitlabURL))
	if err != nil {
		log.Fatalf("Failed to create GitLab client: %v", err)
	}

	// Prepare the commit action
	actions := []*gitlab.CommitActionOptions{
		{
			Action:   gitlab.FileAction(gitlab.FileCreate),
			FilePath: gitlab.String(filePath),
			Content:  gitlab.String(fileContent),
		},
	}

	// Create the commit
	commitOptions := &gitlab.CreateCommitOptions{
		Branch:        gitlab.String(branchName),
		CommitMessage: gitlab.String(commitMessage),
		Actions:       actions,
	}

	commit, _, err := git.Commits.CreateCommit(projectID, commitOptions)
	if err != nil {
		log.Fatalf("Failed to create commit: %v", err)
	}

	// Print success message
	fmt.Printf("Commit %s created successfully\n", commit.ShortID)
}

