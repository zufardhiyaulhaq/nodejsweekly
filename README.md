# nodeweekly
Get data from nodeweekly.com and create Weekly CRDs based on community-operator & push to datastore

### Developing locally
- export this variable
```
export GITHUB_TOKEN="GITHUB_TOKEN"
export GITHUB_ORGANIZATION="GITHUB_USERNAME/ORGANIZATION"
export GITHUB_REPOSITORY="GITHUB_REPOSITORY"
export GITHUB_REPOSITORY_PATH="GITHUB_REPOSITORY_PATH"
export GITHUB_BRANCH="BRANCH"

export COMMUNITY="COMMUNITY_NAME"
export TAGS="TAGS"
export NAMESPACE="CRD_NAMESPACE"
export IMAGE="IMAGE_FOR_WEEKLY"
```
for example
```
export GITHUB_TOKEN="token"
export GITHUB_ORGANIZATION="zufardhiyaulhaq"
export GITHUB_REPOSITORY="community-ops"
export GITHUB_REPOSITORY_PATH="./manifest/nodejs-community/"
export GITHUB_BRANCH="master"

export COMMUNITY="NodeJs Indonesia Community"
export TAGS="weekly,node"
export NAMESPACE="node-community"
export IMAGE="https://calebmadrigal.com/images/nodejs-logo.png"
```
- Build & Run
```
go build -o nodeweekly cmd/nodeweekly/*.go
./nodeweekly
```

### Build Docker
- To build your image
```
make build REPOSITORY=username/repository TAG=tag
```
