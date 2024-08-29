import { Octokit } from "octokit";

async function getRepositoryData(repoUrl?: string){
  const octokit = new Octokit({ 
    auth: 'YOUR-TOKEN',
  });
  
  await octokit.request(`GET ${repoUrl}`, {
    owner: "github",
    repo: "docs",
    per_page: 2
  });
}
