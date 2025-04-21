# Medoc-Kms

this is the KMS (Key Management Service) used by every other application for encryption and decription

# Version : 1.2.0

Authour [@vinayakgupta29](https://www.github.com/vinayakgupta29)

#### The Key , HMAC and encrypted data are stored and sent as base64 String.

#### The Key Struct is stored after encryption in the Keystore Folder/Dir as Binary Data.

#### For the most part you don't have to edit anything in this.

## Requirements

### `Go env setup or Docker setup.`

1. <b>GO Setup</b>

```sh
 go mod download
```

```sh
go build -o ./build/kms
```

```sh
./build/kms
```

> [!TIP]
>
> For windows systems you might wanna rename the object file as `./build/kms.exe`

2. <b>Docker Setup</b>

Setup docker in your system and then clone this repository.

```sh
docker build go-kms
```

```sh
docker run -p host_port:container_port go-kms
```

> [!NOTE]
>
> The container port is defined in the Dockerfile with the `EXPOSE` command and the host port would be your machine's port which will serve the traffic

### API ROUTES

#### Create Key

- **URL** : `http://<host>/create`
- **Method** : `POST`
- **Request Header** :

```json
{
  "X-Client": "<your app's name>",
  "Authorization": "<app's associate HMAC>"
}
```

- **Response** :
  `status`:`200`
  - **Respons Body** : `<Encrypted base64 Key ID>`

### Get Key

- **URL** : `http://<host>/get`
- **Method** : `POST`
- **Request Header** :

```json
{
  "X-Client": "<your app's name>",
  "Authorization": "<app's associate HMAC>"
}
```

- **Request Body** :

```json
{
  "keyId": "<KeyId of the key to be fetched>"
}
```

- **Response** :
  `status`:`200`
  - **Respons Body** : `<Encrypted base64 Key Data>`

