# pr-status-go

Hobby project for learning Go. Command line program that checks pull request statuses on a bitbucket server.
There is something that I am not using, called https://github.com/go-bitbucket/bitbucket.
Not really interested in this since this is about learning Go, and I also believe this is for connecting to bitbucket.org only.

## Configuration

You need to add your connection details to config.go.

- username - Your bitbucket username
- token - Generated oauth2 token for the user account
- serverHostname - Hostname/address of server
- project - Name of your project (as part of the url)
- repoPath - Name of your repository (as part of the url)
