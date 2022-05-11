# pull-request-finder

The pull-request-finder finds all pull request not older than three days from a repo and displays the latest statuses of the HEAD-commits of these pull requests.

## How to use it


In the configuration folder there is a configuration.yaml file. Adjust the name and the owner of the repository you want to check with your own values.
After you set the correct repository values you move into the pull-request-finder folder and run 

```
go run main.go
```

When no error is displayed there should be a file named pullRequests.yaml which contains the information.