# build the terraform provider plugin in the current directory
build:
	go build -o terraform-provider-zertoCloudManager main.go

# compile the terraform provider plugin in the correct directories for usage with terraform
compile:
	echo "Compiling for linux and windows"
	GOOS=linux GOARCH=amd64 go build -o .terraform/plugins/linux_amd64/terraform-provider-zertoCloudManager main.go
	# GOOS=windows GOARCH=amd64 go build -o .terraform/plugins/windows_amd64/terraform-provider-zertoCloudManager main.go

# install all dependencies required to build or run the terraform provider
deps:
	go get github.com/hashicorp/terraform/

# run the terraform provider on demand
run:
	go run main.go