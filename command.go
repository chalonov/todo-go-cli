package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Agrega una tarea por hacer")
	flag.StringVar(&cf.Edit, "edit", "", "Edita una tarea ya creada. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Indica una tarea para eliminar")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Indica una tarea para marcar como completa")
	flag.BoolVar(&cf.List, "list", false, "Muestra una lista de todas la tareas creadas")

	flag.Parse()

	return &cf
}	

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {

	case cf.List: 
		todos.print()

	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, formato inválido para editar. Usar el id de la tarea")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error, indice equivocado para editar")
			os.Exit(1)
		}
		
		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	
	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Comando equivocado")
	}
}