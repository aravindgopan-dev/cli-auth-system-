package cli

import "fmt"



type CLIHandler struct{

}

func NewCLIHandler() *CLIHandler{
	return &CLIHandler{

	}
}

func Run(h *CLIHandler){
	fmt.Println("--welcome to cli-auth-system--")
	

}