import gitlab

GITLAB_URL = 'https://gitlab.com'  
PRIVATE_TOKEN = 'glpat-Pi-lkkztP23F8LssgLOV'  
PROJECT_ID = 'some-project' 
BRANCH_NAME = 'main'  
COMMIT_MESSAGE = 'this is to add a new file in the'

# Initialize the GitLab API client
gl = gitlab.Gitlab(GITLAB_URL, private_token=PRIVATE_TOKEN)

# Get the project object
project = gl.projects.get(PROJECT_ID)

# Prepare the commit data
commit_data = {
    'branch': 'main',
    'commit_message': 'Add hello.txt file',
    'actions': [
        {
            'action': 'create',
            'file_path': 'hello.txt',
            'content': 'Hello, World!',
        }
    ]
}

# Create the commit
commit = project.commits.create(commit_data)
print(f"Commit {commit['id']} created successfully")
