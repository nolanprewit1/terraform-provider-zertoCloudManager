###################
Terraform Provider for Zerto Cloud Manager
###################

- Website: https://www.terraform.io

<img  src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg"  width="600px">


- Website: https://zerto.com

<img  src="https://www.zerto.com/wp-content/themes/zerto_com/zerto-main-logo.png"  width="600px">

  
###################
Requirements
###################

-  [Terraform](https://www.terraform.io/downloads.html) 0.10.x

-  [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

###################
Building The Provider
###################

LINUX:
Clone the repository:
```sh

$ mkdir -p $GOPATH/src/github.com/nolanprewit1; cd  $GOPATH/src/github.com/nolanprewit1

$ git clone https://github.com/nolanprewit1/terraform-provider-zertoCloudManager.git

```
Enter the provider directory and build the provider

```sh

$ cd  $GOPATH/src/github.com/nolanprewit1/terraform-provider-zertoCloudManager

$ make deps

$ make compile

``` 

WINDOWS:
Clone the repository:

```sh

$ mkdir -p $GOPATH/src/github.com/nolanprewit1; cd  $GOPATH/src/github.com/nolanprewit1

$ git clone https://github.com/nolanprewit1/terraform-provider-zertoCloudManager.git

```
Enter the provider directory and build the provider
```sh

$ cd  $GOPATH/src/github.com/nolanprewit1/terraform-provider-zertoCloudManager

$ go build -o .terraform/plugins/windows_amd64/terraform-provider-zertoCloudManager.exe

```
###################
Using the provider
###################

1. Create a new terraform configuration file (terraform.tf)
```sh

provider  "zertocloudmanager"  {
	address = var.zertocloudmanager_address
	port = var.zertocloudmanager_port
	username = var.zertocloudmanager_username
	password = var.zertocloudmanager_password
}

resource  "zertocloudmanager_zorg"  "test"  {
	name = "test12345"
	crmidentifier = "test0111"
}

```
2. Create a new terraform variables configuration file (variables.tf)

```sh

variable  "zertocloudmanager_address"  {
  default = "https://192.168.1.202"
}
variable  "zertocloudmanager_port"  {
  default = "9989"
}
variable  "zertocloudmanager_username"  {
  default = "zcm_admin"
}
variable  "zertocloudmanager_password"  {
  default = "putPasswordHere"
}

```

3. Initialize the provider

```sh

terraform init

```  

4. Apply the terraform configuration

```sh

terraform apply

```
###################
Resource Arguments
###################
```
zertocloudmanager_zorg
	- name          | (Required,String) | The name used to identify the ZORG. The name must be unique
	- crmidentifier	| (Required,String) | A internal CRM identification for the ZORG"
```
###################
Directory Structure
###################
```
.
├── .terraform/             			# Directory where compiled terraform providers should live
├── api						
|	├── client.go				# Code related to the generation of the api url and getting the header key after basic auth is performed
├── provider   
|	├── provider.go				# Code defining the provider and the resources available
|	├── resource_zorg.go			# The CRUD specific actions for the zorg resource
├── .gitignore              			# Files and directories for git to ignore
├── main.go                 			# The main initialization of the terraform provider application
├── Makefile					# A make file use for easier compiling of the application and reference for multiple operating systems
├── terraform.tf				# The terraform configuration file
├── variables.tf				# Variables used by the terraform configuration file. Not kept in source code. 
└── README.md					# Application readme file
```
###################
TODO
###################


###################
Code References
###################
- https://github.com/sethvargo/terraform-provider-googlecalendar
- https://github.com/ContainerSolutions/terraform-provider-template/tree/v0.1.1
- https://github.com/spaceapegames/terraform-provider-example
