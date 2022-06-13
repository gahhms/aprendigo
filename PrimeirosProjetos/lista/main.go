package main

import (
	"fmt"
	"sort"
	"time"
)

var funcao int

type ToDoList struct {
	ToDo  string
	Date  time.Time
	Added string
	Done  bool
}

type sortByDate []ToDoList

func (x sortByDate) Len() int           { return len(x) }
func (x sortByDate) Less(i, j int) bool { return x[i].Date.Before(x[j].Date) }
func (x sortByDate) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func statusToDo(t []ToDoList) {
	for i, v := range t {
		i++
		bef := v.Date.Before(time.Now())
		if v.Done == true {
			status := "Feito\t"
			arrumaData := v.Date.Format("02/01/2006")
			fmt.Printf("%v\t%s\t%s\t\t%s\t\t%s", i, arrumaData, v.Added, status, v.ToDo)
			fmt.Println("")
		} else if bef == true {
			status := "Atrasado"
			arrumaData := v.Date.Format("02/01/2006")
			fmt.Printf("%v\t%s\t%s\t\t%s\t\t%s", i, arrumaData, v.Added, status, v.ToDo)
			fmt.Println("")
		} else if v.Done == false { // ARRUMAR AQUI!!!!!!!!!!!!
			status := "Pendente"
			arrumaData := v.Date.Format("02/01/2006")
			fmt.Printf("%v\t%s\t%s\t\t%s\t\t%s", i, arrumaData, v.Added, status, v.ToDo)
			fmt.Println("")
		}
	}
}

func statusOnlyToDo(t []ToDoList) {
	for i, v := range t {
		i++
		bef := v.Date.Before(time.Now())
		if bef == true {
			status := "Atrasado"
			arrumaData := v.Date.Format("02/01/2006")
			fmt.Printf("%v\t%s\t%s\t\t%s\t\t%s", i, arrumaData, v.Added, status, v.ToDo)
			fmt.Println("")
		} else if v.Done == false {
			status := "Pendente"
			arrumaData := v.Date.Format("02/01/2006")
			fmt.Printf("%v\t%s\t%s\t\t%s\t\t%s", i, arrumaData, v.Added, status, v.ToDo)
			fmt.Println("")
		}
	}
}

func show(t []ToDoList) {
	fmt.Println("")
	fmt.Printf("Suas tarefas são:\n")
	fmt.Println("Nº\tData Final:\tAdicionada em:\tStatus:\t\tTarefa:")
	fmt.Println("")
	statusToDo(t)
}

func showAll(t []ToDoList) {
	fmt.Println("")
	fmt.Printf("Suas tarefas ordenadas por data são:\n")
	fmt.Println("Nº\tData Final:\tAdicionada em:\tStatus:\t\tTarefa:")
	fmt.Println("")
	statusToDo(t)
}

func showToDo(t []ToDoList) {
	fmt.Println("")
	fmt.Printf("Suas tarefas não concluídas são:\n")
	fmt.Println("Nº\tData Final:\tAdicionada em:\tStatus:\t\tTarefa:")
	fmt.Println("")
	statusOnlyToDo(t)
}

func escolhaFuncao(i int) {
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		show(Tarefas)
		voltarMenu()
	case 3:
		sort.Sort(sortByDate(Tarefas))
		showAll(Tarefas)
		voltarMenu()
	case 4:
		sort.Sort(sortByDate(Tarefas))
		showToDo(Tarefas)
		voltarMenu()
	case 5:
		fmt.Println("Fim do Programa!")
	default:
		fmt.Printf("\nOpção Inválida! Digitou errado?\n\n")
		lerEscolha()
	}
}

func lerEscolha() {

	fmt.Printf("Escolha a função que gostaria de executar!")
	fmt.Printf("\n1 - Adicionar nova tarefa\n2 - Listar todas as tarefas\n3 - Listar tarefas por data\n4 - Listar tarefas pendentes\n5 - Encerrar programa\n")
	fmt.Printf("Sua escolha: ")
	_, err := fmt.Scanln(&funcao)
	if err != nil {
		fmt.Println("Deu xabu, tenta de novo!")
	}
	escolhaFuncao(funcao)
}

func voltarMenu() {

	var voltamenu string
	fmt.Println("")
	fmt.Printf("Deseja voltar ao Menu principal? (s/n): ")
	_, err := fmt.Scanln(&voltamenu)
	if err != nil {
		fmt.Println("Deu xabu, tenta de novo!")
	}
	switch voltamenu {
	case "s":
		fmt.Println("volta menu")
		lerEscolha()
	case "n":
		fmt.Println("Fim do Programa")
	default:
		fmt.Printf("\nOpção Inválida! Digitou errado?\n")
		voltarMenu()
	}
}

var Tarefas = []ToDoList{
	{"Lavar roupas brancas", time.Date(2023, time.June, 15, 23, 0, 0, 0, time.UTC), "03/06/2022", false},
	{"Escrever o código dessa lista ", time.Date(2009, time.January, 07, 23, 0, 0, 0, time.UTC), "02/06/2022", false},
	{"Levar o cachorro para cagar", time.Date(2022, time.June, 01, 23, 0, 0, 0, time.UTC), "02/06/2022", false},
}

func main() {

	today := time.Now().Format("02/01/2006")

	fmt.Printf("Bem vindo a SUPER LISTA DE TAREFAS!\n\n")
	fmt.Println("Bom dia, hoje é", today)
	fmt.Println("")

	lerEscolha()

}
