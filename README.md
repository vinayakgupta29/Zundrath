# Medoc-Kms

this is the KMS (Key Management Service) used by every other application for encryption and decription

# Version : 1.1.1

Authour [@vinayakgupta29](https://www.github.com/vinayakgupta29)

#### The Key , HMAC and encrypted data are stored and sent as base64 String.

#### The Key Struct is stored after encryption in the Keystore Folder/Dir as Binary Data.

#### For the most part you don't have to edit anything in this.

#### Requirements

### `Go env setup or Docker setup.`

1. <b>GO Setup</b>

```sh
 go mod download
```

```sh
go build main.go -o ./build/kms
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
docker build medoc-kms
```

```sh
docker run -p host_port:container_port medoc-kms
```

> [!NOTE]
>
> The container port is defined in the Dockerfile with the `EXPOSE` command and the host port would be your machine's port which will serve the traffic

### API ROUTES

#### Create Key

- url : http://<host>/create
- method : POST
- Request Header : {
  "X-Client":<app's name>,
  "Authorization":<app's associate HMAC>
  }
- Response : 200 :
  <Encrypted base64 Key Data>

### Get Key

- url : http://\<host\>/get
- method : POST
- Request Header : {
  "X-Client" : <app's name>,
  "Authorization": <app's associate HMAC>
  }
- Request Body : {
  'keyId' : \<KeyId of the key to be fetched\>
  }
- Response : 200 :
  <Encrypted base64 Key Data>

The HMAC for any app can be found in the _main.go_ file with the variable `CLIENTID`

> [!WARNING]
>
> Don't change the getKey Route to a Get Request. It is the way it is for a reason.
> If you think you are smart enough to change the working or code in any manner then feel free to do so but if in the meantime you manage to break the working or lead to a security failure then feel free to drop your name below this in the Hall of Fame alongwith No. of hours wasted.

# Hall of Fame
