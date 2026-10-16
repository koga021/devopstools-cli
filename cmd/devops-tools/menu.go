package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

func showMenu() {
	fmt.Println("Usage: devops-tools <command> [args]")
	fmt.Println("Available commands:")
	fmt.Println()

	fmt.Println("list template \t\t\t\t List all templates")
	fmt.Println("install template <template-name> \t Install the template selected")
}

func checkItemMenu(parametro []string, item string) bool {
	if contains(parametro, item) {
		return true
	}
	return false
}

func contains(arr []string, item string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func FirstMenu() {
	for {
		var opcao string

		err := huh.NewSelect[string]().
			Title("DevOps Tools").
			Options(
				huh.NewOption("Templates", "templates"),
				huh.NewOption("Sair", "sair"),
			).
			Value(&opcao).
			Run()

		if err != nil || opcao == "sair" {
			os.Exit(0)
		}

		if opcao == "templates" {
			MenuTemplates()

		}
	}
}

func MenuTemplates() {
	for {
		var opcao string

		err := huh.NewSelect[string]().
			Title("DevOps Tools - Templates").
			Options(
				huh.NewOption("Terraform", "templates_terraform"),
				//huh.NewOption("Ansible", "templates_ansible"),
				//huh.NewOption(".gitlab-ci.yml", "templates_gitlab_ci_yaml"),
				huh.NewOption("Voltar", "voltar"),
				huh.NewOption("Sair", "sair"),
			).
			Value(&opcao).
			Run()

		if err != nil || opcao == "sair" {
			os.Exit(0)
		}

		if opcao == "templates_terraform" {
			ListCatalog()
			fmt.Println("Escolha o Template a ser instalado")
			fmt.Scanln()
		}
		if opcao == "voltar" {
			return
		}
	}
}
