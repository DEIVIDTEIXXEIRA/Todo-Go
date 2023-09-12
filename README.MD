


<h1 align="center">
<img src="logo.png">
</h1>


#### Esse Projeto Esta dividido em duas partes API e WEB.

#### Na parte WEB da aplicação, os usuários podem se cadastrar, criar tarefas, formar equipes para seus projetos e organizar tarefas com prazos a serem cumpridos dentro dessas equipes. O restante da API trata de todos os dados e lógica, enviando todas as solicitações e respostas em tempo real utilizando concorrência.


# API REST
### Tecnologias
- [Go](https://go.dev/dl/)
- [MYSQL](https://www.mysql.com/downloads/)

# WEB
### Tecnologias

- [Go](https://go.dev/dl/)
- [JavaScript](https://js.org/index.html)

# Iniciando
### Requisitos


```shell
$ git clone git@github.com:DEIVIDTEIXXEIRA/Todo-Go.git && cd todo
```

### Configurando
- Crie um banco de dados e execute o script do arquivo [SQL](https://github.com/DEIVIDTEIXXEIRA/Todo-Go/blob/main/API/sql/sql.sql)

- Atualize as configurações para sua maquina no arquivo [.env](https://github.com/DEIVIDTEIXXEIRA/Todo-Go/blob/main/API/.env)


# Iniciando o Backend

```shell
# Acesse a pasta da API
$ cd api

# Crie um executavel
$ go build api

# Execute
$ ./api
```

    
# Iniciando o WebApp

```shell
# Acesse a pasta da webapp
$ cd webapp

# Crie um executavel
$ go build webapp

# Execute
$ ./webapp
```

# Em seu navegador
- Acesse http://localhost:9090/

# Funcionalidades

O Todo é uma aplicação completa de organização de tarefas com diversas funcionalidades:

- **Registro e Autenticação**: Após o cadastro, o sistema utiliza autenticação por token, garantindo maior segurança nas comunicações entre o cliente e o servidor.

- **Usuário**: Usuário: Crie seu usuário autenticado e edite-o quando quiser.

 - **Tarefas**: Tarefas: Com um usuário registrado, você pode criar tarefas e definir prazos.

 - **Equipes**: Crie equipes para definir tarefas específicas para temas ou projetos.

 - **Benefícios da aplicação utilizar Go**:A aplicação utiliza concorrência em algumas funcionalidades, tornando as requisições muito mais eficientes.


# Funcionalidades futuras:
- Testes

- Adicionar usuarios a uma equipe e esse usuario poderá visualizar, completar, editar e criar tarefas dentro dessa equipe.

- Pesquisar tarefa com filtro de nome ou data.

- Adicionar paginas organizando as tarefas mediante a data atual.

# Contato
Para qualquer dúvida ou sugestão, sinta-se à vontade para entrar em contato comigo:
- **E-mail**: deividteixeira.go@gmail.com
- **GitHub**: https://github.com/DEIVIDTEIXXEIRA
- **Instagram**: @deivid_tx

