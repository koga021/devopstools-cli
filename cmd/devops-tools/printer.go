package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/douglas.ribas/devops-tools/internal/catalog"
)

func printItems(items []catalog.MenuItem) {
	for _, item := range items {
		if len(item.Items) > 0 {
			printItems(item.Items)
		} else {
			if item.Description != "" {
				fmt.Printf("%s \t %s\n", item.ID, item.Description)
			} else {
				fmt.Printf("  - %s\n", item.ID)
			}
		}
	}
}

func ListCatalog() {

	config, err := catalog.Load("configs/catalog.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Navega na árvore: menu > templates > terraform
	// Primeiro, encontra "templates" no primeiro nível
	templatesItem := findMenuItem(config.Menu, "templates")
	if templatesItem == nil {
		fmt.Println("Categoria 'templates' não encontrada no catálogo.")
		return
	}

	// Depois, encontra "terraform" dentro de "templates"
	terraformItem := findMenuItem(templatesItem.Items, "terraform")
	if terraformItem == nil {
		fmt.Println("Categoria 'terraform' não encontrada no catálogo.")
		return
	}

	// Agora itera sobre os items de terraform e cria as opções do menu
	var options []huh.Option[string]
	for _, item := range terraformItem.Items {
		// Usa o Name como label e o ID como valor
		options = append(options, huh.NewOption(item.Name, item.ID))
	}
	// Adiciona opções fixas de navegação
	options = append(options, huh.NewOption("Voltar", "voltar"))
	options = append(options, huh.NewOption("Sair", "sair"))

	for {
		var opcao string

		err := huh.NewSelect[string]().
			Title("Templates Terraform").
			Options(options...).
			Value(&opcao).
			Run()

		if err != nil || opcao == "sair" {
			os.Exit(0)
		}
		if opcao == "voltar" {
			return
		}

		// Aqui você pode tratar o template selecionado
		selected := findMenuItem(terraformItem.Items, opcao)
		if selected != nil {
			fmt.Printf("\nSelecionado: %s\n", selected.Name)
			fmt.Printf("Descrição:   %s\n", selected.Description)
			if selected.Source != nil {
				fmt.Printf("Repositório: %s (ref: %s)\n", selected.Source.URL, selected.Source.Ref)
			}
			fmt.Println("\nDeseja instalar o template? (yes/no)")
			fmt.Println(selected.ID)
			fmt.Scanln()
		}
	}
}

// findMenuItem busca um MenuItem pelo ID em um slice de MenuItems
func findMenuItem(items []catalog.MenuItem, id string) *catalog.MenuItem {
	for i := range items {
		if items[i].ID == id {
			return &items[i]
		}
	}
	return nil
}
