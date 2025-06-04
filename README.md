# SonarCloud API CLI

## Descripción

Este proyecto es una herramienta de línea de comandos (CLI) para interactuar con la API de SonarCloud. Permite gestionar y obtener información sobre proyectos, perfiles de calidad, etc.

## Requisitos Previos

- Go 1.x o superior
- Una cuenta de SonarCloud y un token de API

## Instalación

1. **Clona el repositorio:**

    ```sh
    git clone https://github.com/warckor/sonar-api.git
    cd sonar-api
    ```

2. **Construye el ejecutable:**

    ```sh
    go build -o sonarcli main.go
    ```

3. **(Opcional) Mueve el ejecutable a una ruta en tu PATH:**
    Para poder ejecutar `sonarcli` desde cualquier directorio:

    ```sh
    sudo mv sonarcli /usr/local/bin/
    ```

## Configuración

Antes de usar la CLI, necesitas configurar tu token de API de SonarCloud y, opcionalmente, tu organización por defecto.

La CLI buscará un archivo de configuración `config.json` en los siguientes directorios (en orden de precedencia):

1. Directorio actual (`./config.json`)
2. Directorio de configuración del usuario (`~/.config/sonar-api/config.json` o `%APPDATA%\sonar-api\config.json` en Windows)
3. Directorio del ejecutable.

El archivo `config.json` debe tener el siguiente formato:

```json
{
  "organization": "tu-organizacion-por-defecto",
  "token": "tu-token-de-api-de-sonarcloud"
}
```

Si no se especifica la organización mediante el flag `--org` en un comando, se utilizará la `organization` del archivo de configuración. El `token` es siempre requerido para autenticarse con la API de SonarCloud.

## Uso

A continuación se muestran los comandos disponibles y cómo usarlos.

### Comandos Principales

La CLI se invoca con el comando `sonarcli`.

```sh
sonarcli [comando] [subcomando] [flags]
```

### Proyectos (`project`)

#### Obtener detalles de un proyecto específico

Obtiene los detalles de un proyecto específico en SonarCloud utilizando su clave de proyecto o nombre.

**Uso:**

```sh
sonarcli get project [flags]
```

**Flags:**

- `--org, -o string`: Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica).
- `--project-key, -p string`: Clave del proyecto de SonarCloud.
- `--name, -n string`: Nombre del proyecto de SonarCloud.

**Ejemplo:**

```sh
sonarcli get project --org mi-organizacion --project-key mi-clave-de-proyecto
sonarcli get project --name "Mi Proyecto Asombroso"
```

*Nota: Se requiere al menos `--project-key` o `--name`.*

#### Listar todos los proyectos

Lista todos los proyectos disponibles en SonarCloud para una organización específica.

**Uso:**

```sh
sonarcli list project [flags]
```

**Flags:**

- `--org, -o string`: Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica).

**Ejemplo:**

```sh
sonarcli list project --org mi-organizacion
```

*(Más comandos y subcomandos serán documentados aquí a medida que se implementen, como `quality`, `user`, `actions create`, `actions get`, `actions list`, `actions status` etc.)*

## Contribuciones

Las contribuciones son bienvenidas. Por favor, abre un issue o un pull request para discutir los cambios.

## Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.
