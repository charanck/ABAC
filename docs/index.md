# Branching Strategy

main -> is the main and base branch and the development branch with latest changes

feature/{FEATURE_NAME} -> is for developing a new feature based of main branch

bugfix/{BUG_NAME} -> is for fixing a bug

# Changelog and version

* Semantic versioning will be used
* Changes for all the version must be maintained in the changelog.md and must be updated before each PR merge
* Version.md will contain the lastest version of the application

# Pull request checklist

* [ ] Use Meaningful branch name and PR title
* [ ] Update unit test
* [ ] Update changelog and version.md

# Proto and grpc generation

Define the rpc service in the protobuf/abac.proto file and dont make changes to the files in the generated folder and commit

## Pre requestics to run the grpc code generation:

protoc compiler -> install it using the below commands

* go install google.golang.org/protobuf/cmd/protoc-gen-go
* go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

Use the below command to generate the grpc code from the proto definition

```
 protoc --go_out=./protobuf/generated --go-grpc_out=./protobuf/generated  ./protobuf/abac.proto
```
