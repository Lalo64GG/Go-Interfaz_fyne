Simulador de Estacionamiento en GO

Este proyecto es un simulador de estacionamiento desarrollado en Go, utilizando el framework de interfaz gráfica Fyne. El simulador permite observar el estado de los vehículos en el estacionamiento, mostrando información sobre los espacios disponibles, vehículos en espera y estado de cada vehículo.

Requisitos

- Go (https://golang.org/doc/install) (versión 1.16 o superior)
- Fyne (https://fyne.io/) (versión 2.0 o superior)

Instalación

1. Clona este repositorio
```bash
git clone https://github.com/Lalo64GG/Simulador-Estacionamiento-GO.git
cd Simulador-Estacionamiento-GO
```

2. Instala las dependencias de Fyne
Asegúrate de tener Go instalado y configura el módulo de Go para el proyecto:
```bash
go mod init Simulador-Estacionamiento-GO
go mod tidy
```

Instala el paquete de Fyne:
```bash
go get fyne.io/fyne/v2
```

3. Ejecuta el simulador
```bash
go run ./src/main.go
```

Uso del Proyecto

El simulador muestra una interfaz gráfica con los siguientes componentes:
- Tabla de Vehículos: muestra el estado de cada vehículo en el estacionamiento.
- Botón "Iniciar Simulación": inicia la simulación de los vehículos entrando y saliendo del estacionamiento.
- Indicadores de Espacios y Vehículos en Espera: indican la cantidad de espacios disponibles y vehículos esperando por un espacio.

Controles
- Haz clic en el botón "Iniciar Simulación" para comenzar a generar vehículos que intentarán estacionarse.
- Observa cómo cambia el estado de cada vehículo en la tabla.

Estructura de Carpetas

La estructura del proyecto está organizada en carpetas para facilitar la modularidad y escalabilidad del código:

```plaintext
simulador-estacionamiento/
│
├── src/
│   ├── controllers/
│   │   ├── ParkingController.go        # Controlador para la lógica del estacionamiento
│   │   └── VehicleController.go        # Controlador para la gestión de los vehículos
│   │
│   ├── models/
│   │   └── VehicleStatus.go            # Modelo que representa el estado de cada vehículo
│   │
│   ├── services/
│   │   ├── ParkingService.go           # Lógica de negocio para la gestión de espacios de estacionamiento
│   │   └── VehicleService.go           # Lógica de negocio para la gestión de vehículos
│   │
│   ├── utils/
│   │   ├── Poisson.go                  # Utilidad para generación de tiempos aleatorios (ej. distribución de Poisson)
│   │   └── RandomTime.go               # Utilidad para generar tiempos aleatorios
│   │
│   ├── main.go                         # Punto de entrada principal de la aplicación
│   └── car.png                         # Imagen de un coche (opcional para visualización en la interfaz)
│
└── README.md                           # Documentación del proyecto
```

Explicación de Carpetas

- controllers/: Contiene los controladores que manejan la interacción entre la interfaz gráfica y los servicios de negocio.
- models/: Contiene los modelos de datos que representan los estados y propiedades de los vehículos.
- services/: Contiene la lógica de negocio principal, como la gestión de espacios de estacionamiento y el control de vehículos.
- utils/: Contiene utilidades generales, como la generación de tiempos aleatorios.
- main.go: Archivo principal que inicializa la aplicación y configura la interfaz gráfica utilizando Fyne.

Uso de Go y Fyne

Go
Go es un lenguaje de programación eficiente y simple que permite desarrollar aplicaciones de alto rendimiento. En este proyecto, Go se utiliza para manejar la concurrencia en la simulación de vehículos, gestionando múltiples goroutines para representar los vehículos que intentan estacionarse al mismo tiempo.

Fyne
Fyne es un framework de interfaz gráfica para Go, diseñado para ser multiplataforma. En este proyecto, Fyne se utiliza para crear una interfaz de usuario interactiva, donde se puede observar el estado del estacionamiento en tiempo real. Algunos componentes importantes de Fyne en este proyecto son:

- widget.NewLabel: Para mostrar textos simples como los indicadores de estado.
- widget.NewTable: Para mostrar una tabla con los detalles de cada vehículo.
- widget.NewButton: Para crear el botón que inicia la simulación.
- container.NewVBox y container.NewVScroll: Para organizar los elementos en la ventana.

Recursos
Para obtener más información sobre Go y Fyne, consulta los enlaces oficiales:
- Documentación de Go (https://golang.org/doc/)
- Documentación de Fyne (https://developer.fyne.io/)

Contribuciones

Si deseas contribuir a este proyecto, por favor abre un pull request o crea un issue para discutir cambios o mejoras.

Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.
