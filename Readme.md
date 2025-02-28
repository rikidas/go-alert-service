# Go Email Service

Este es un servicio de envío de correos electrónicos en Go. Permite enviar correos electrónicos con un asunto y mensaje utilizando la línea de comandos.

## Requisitos

- Go 1.18 o superior

## Instalación

1. Clona el repositorio y ve al directorio del proyecto:

    ```sh
    git clone https://github.com/rikidas/go-email-service.git
    cd go-email-service
    ```

2. Inicializa los módulos de Go:

    ```sh
    go mod tidy
    ```

## Configuración

1. Crea un archivo [.env](http://_vscodecontentref_/1) en la raíz del proyecto con las siguientes variables:

    ```env
    SMTP_HOST=smtp.example.com
    SMTP_PORT=587
    SMTP_USER=tu_usuario
    SMTP_PASS=tu_contraseña
    ```

2. Asegúrate de que el archivo [.env](http://_vscodecontentref_/2) esté en el [.gitignore](http://_vscodecontentref_/3) para no subir tus credenciales al repositorio.

## Uso

Ejecuta el servicio desde la línea de comandos pasando los argumentos necesarios:

```sh
./email-service --email="correo@ejemplo.com" --subject="Prueba" --message="Este es un mensaje"
```