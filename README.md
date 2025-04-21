# go ci example

Create github fine-grained token for repository

```
Profile > Settings > Developer settings > Personal access tokens
```

Token permissions:
- Read metadata
- Read and write dependabot alerts
- Read and write actions
- Read and write code
- Read and write issues

Add repository secret GH_ACCESS_TOKEN

```
Repository > Settings > Secrets and variables > Actions
```