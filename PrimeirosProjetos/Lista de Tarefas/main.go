package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

var funcao int

type ToDoList struct {
	ToDo  string
	Date  time.Time
	Added time.Time
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
			arrumaDataTarefa := v.Date.Format("02/01/2006")
			arrumaDataAdd := v.Added.Format("02/01/2006")
			fmt.Printf("%v\t%s\t%s\t%s\t%s", i, arrumaDataTarefa, arrumaDataAdd, status, v.ToDo)
			fmt.Println("")
		} else if bef == true {
			status := "Atrasado"
			arrumaDataTarefa := v.Date.Format("02/01/2006")
			arrumaDataAdd := v.Added.Format("02/01/2006")
			fmt.Printf("%v\t%s\t%s\t%s\t%s", i, arrumaDataTarefa, arrumaDataAdd, status, v.ToDo)
			fmt.Println("")
		} else if v.Done == false {
			status := "Pendente"
			arrumaDataTarefa := v.Date.Format("02/01/2006")
			arrumaDataAdd := v.Added.Format("02/01/2006")
			fmt.Printf("%v\t%s\t%s\t%s\t%s", i, arrumaDataTarefa, arrumaDataAdd, status, v.ToDo)
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
			arrumaDataTarefa := v.Date.Format("02/01/2006")
			arrumaDataAdd := v.Added.Format("02/01/2006")
			fmt.Printf("%v\t%s\t%s\t%s\t%s", i, arrumaDataTarefa, arrumaDataAdd, status, v.ToDo)
			fmt.Println("")
		} else if v.Done == false {
			status := "Pendente"
			arrumaDataTarefa := v.Date.Format("02/01/2006")
			arrumaDataAdd := v.Added.Format("02/01/2006")
			fmt.Printf("%v\t%s\t%s\t%s\t%s", i, arrumaDataTarefa, arrumaDataAdd, status, v.ToDo)
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

func showPorData(t []ToDoList) {
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
		addTarefa()
		voltarMenu()
	case 2:
		sort.Sort(sortByDate(Tarefas))
		alteraStatus(Tarefas)
		voltarMenu()
	case 3:
		show(Tarefas)
		voltarMenu()
	case 4:
		sort.Sort(sortByDate(Tarefas))
		showPorData(Tarefas)
		voltarMenu()
	case 5:
		sort.Sort(sortByDate(Tarefas))
		showToDo(Tarefas)
		voltarMenu()
	case 6:
		fmt.Println("Fim do Programa!")
	default:
		fmt.Printf("Opção Inválida! Digitou errado?\n\n")
		lerEscolha()
	}
}

func alteraStatus(t []ToDoList) {
	showPorData(Tarefas)
	fmt.Println("")
	fmt.Printf("Qua das tarefas gostaria de alterar o Status? (digite o número da tarefa): ")
	var tarefaAlterar int
	_, err := fmt.Scan(&tarefaAlterar)
	if err != nil {
		fmt.Println("Deu xabu, tenta de novo!")
	}
	tarefaAlterar--
	if Tarefas[tarefaAlterar].Done == false {
		Tarefas[tarefaAlterar].Done = true
	} else {
		Tarefas[tarefaAlterar].Done = false

	}
	showPorData(Tarefas)
	fmt.Println("")
	fmt.Println("Status da tarefa alterado!")
}

func addTarefa() {
	fmt.Printf("Descrição da tarefa: ")
	descricao1 := lerDescricao()

	fmt.Printf("Data da entrega da tarefa: ")
	var dataNovaTarefa string
	_, err1 := fmt.Scan(&dataNovaTarefa)
	if err1 != nil {
		fmt.Println("Deu xabu, tenta de novo!")
		fmt.Println(err1)
	}
	dataNovaTarefaTime, err := time.Parse("02/01/2006", dataNovaTarefa)
	if err != nil {
		fmt.Println("Deu xabu, tenta de novo!")
		fmt.Println(err)
	}
	adicao := []ToDoList{{descricao1, dataNovaTarefaTime, time.Now(), false}}
	Tarefas = append(Tarefas, adicao...)
	fmt.Println("Tarefa adicionada com sucesso!")
	show(Tarefas)
}

func lerEscolha() {

	fmt.Printf("Escolha a função que gostaria de executar!")
	fmt.Printf("\n1 - Adicionar nova tarefa\n2 - Alterar status de uma tarefa\n3 - Listar todas as tarefas\n4 - Listar tarefas por data\n5 - Listar tarefas pendentes\n6 - Encerrar programa\n")
	fmt.Printf("Sua escolha: ")
	_, err := fmt.Scan(&funcao)
	if err != nil {
		fmt.Println("Deu xabu, tenta de novo!")

	}
	escolhaFuncao(funcao)
}

func voltarMenu() {

	var voltamenu string
	fmt.Println("")
	fmt.Printf("Deseja voltar ao Menu principal? (s/n): ")
	_, err := fmt.Scan(&voltamenu)
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

func lerDescricao() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	descricao := scanner.Text()
	fmt.Printf(descricao)
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	descricao1 := scanner1.Text()
	return descricao1
}

var Tarefas = []ToDoList{
	{"Lavar roupas brancas", time.Date(2023, time.June, 15, 23, 0, 0, 0, time.UTC), time.Date(2023, time.June, 15, 23, 0, 0, 0, time.UTC), false},
	{"Escrever o código dessa lista ", time.Date(2009, time.January, 07, 23, 0, 0, 0, time.UTC), time.Date(2023, time.June, 15, 23, 0, 0, 0, time.UTC), false},
	{"Levar o cachorro para cagar", time.Date(2022, time.June, 01, 23, 0, 0, 0, time.UTC), time.Date(2023, time.June, 15, 23, 0, 0, 0, time.UTC), false},
}

func main() {

	today := time.Now().Format("02/01/2006")

	fmt.Printf("Bem vindo a SUPER LISTA DE TAREFAS!\n\n")
	fmt.Println("Bom dia, hoje é", today)
	fmt.Println("")

	lerEscolha()

}
