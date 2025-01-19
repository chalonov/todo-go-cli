# Todo CLI (GO)

Una aplicación de línea de comandos simple pero potente para gestionar tus tareas pendientes, implementada en Go. Esta aplicación te permite crear, editar, completar y eliminar tareas, manteniendo un registro persistente de todas tus actividades.

## Características

La aplicación ofrece las siguientes funcionalidades principales:

- Agregar nuevas tareas con título y fecha de creación automática
- Marcar tareas como completadas, registrando la fecha de finalización
- Editar el título de tareas existentes
- Eliminar tareas
- Listar todas las tareas en una tabla formateada
- Almacenamiento persistente en formato JSON

## Requisitos

Para ejecutar esta aplicación necesitas:

- Go 1.18 o superior
- El paquete `github.com/aquasecurity/table` para el formato de tabla

## Instalación

1. Clona el repositorio:
```bash
git clone https://github.com/tu-usuario/todo-cli
cd todo-cli
```

2. Instala las dependencias:
```bash
go mod tidy
```

3. Compila la aplicación:
```bash
go build
```

## Uso

La aplicación soporta varios comandos a través de flags. Aquí están todos los comandos disponibles:

### Listar todas las tareas
```bash
./todo-cli -list
```
Este comando mostrará una tabla con todas tus tareas, incluyendo:
- Número de índice
- Título de la tarea
- Estado de completitud (✅ o ❌)
- Fecha de creación
- Fecha de finalización (si está completada)

### Agregar una nueva tarea
```bash
./todo-cli -add "Comprar leche"
```
La tarea se agregará con la fecha actual como fecha de creación.

### Editar una tarea existente
```bash
./todo-cli -edit "0:Comprar leche descremada"
```
El formato es `índice:nuevo_título`. El índice debe ser un número válido de tu lista de tareas.

### Marcar una tarea como completada
```bash
./todo-cli -toggle 0
```
Este comando alternará el estado de la tarea entre completada y pendiente. Si se marca como completada, se registrará la fecha de finalización.

### Eliminar una tarea
```bash
./todo-cli -del 0
```
Elimina permanentemente la tarea en el índice especificado.

## Estructura del Proyecto

El proyecto está organizado en varios archivos:

- `main.go`: Punto de entrada de la aplicación
- `command.go`: Manejo de flags y comandos de línea de comandos
- `todo.go`: Implementación de la lógica de tareas
- `storage.go`: Manejo del almacenamiento persistente en JSON

## Almacenamiento

Las tareas se almacenan en un archivo `todos.json` en el directorio de la aplicación. El formato JSON incluye:
```json
[
    {
        "Title": "Ejemplo de tarea",
        "Completed": false,
        "CreatedAt": "2024-12-29T20:04:42.0258377-05:00",
        "CompletedAt": null
    }
]
```
## Licencia

Este proyecto está licenciado bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.