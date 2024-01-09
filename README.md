CRM Backend em Go com GIN, SQL e Swagger
Bem-vindo ao repositório do backend para um CRM leve desenvolvido em Go, utilizando o framework GIN para a construção da API, SQL para persistência de dados e Swagger para documentação.

Pré-requisitos
Certifique-se de ter o Go instalado em sua máquina. Caso ainda não tenha, você pode baixá-lo aqui.

Além disso, é necessário instalar as dependências do projeto. Para isso, utilize o seguinte comando:

bash
Copy code
go mod download
Ferramenta Swagger
A documentação Swagger é gerada utilizando o Swag. Certifique-se de instalá-lo em sua máquina através do seguinte repositório: Swaggo.

Executando o Projeto
Você pode executar o projeto de duas maneiras diferentes:

1. Usando Makefile
Se você tiver o utilitário make instalado, basta executar o seguinte comando:

bash
Copy code
make run
2. Usando o comando go
Se preferir, você pode executar diretamente o arquivo principal do projeto:

bash
Copy code
go run main.go
O servidor será iniciado e estará acessível em http://localhost:8080.

Documentação Swagger
A documentação da API é gerada automaticamente usando o Swagger. Para acessar a documentação, siga os passos abaixo:

Certifique-se de que o servidor está em execução.
Abra o navegador e vá para http://localhost:8080/swagger/index.html.
A documentação Swagger fornecerá detalhes sobre os endpoints disponíveis, parâmetros necessários e exemplos de solicitações.

Contribuindo
Se você quiser contribuir para este projeto, sinta-se à vontade para abrir uma issue ou enviar uma solicitação pull. Esperamos que este projeto evolua com a ajuda da comunidade!
