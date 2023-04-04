# Email service

Email service é um microserviço que consome mensagem de uma fila e envia emails usando AWS SES.

## Antes de começar

1 - Este projeto segue [essa estrutura de pastas](https://github.com/golang-standards/project-layout). <br/>
2 - Eu acredito que seria interessante você saber que Quess não possui somente este projeto.
Nós temos outros projetos, você pode visitá-los em:

- [Toolkit - Módulos auxiliares](https://github.com/QuessApp/toolkit)
- [Core API - Onde a mágica acontece](https://github.com/QuessApp/core-go)
- [Geo Service - Serviço para buscar a localização do usuário por IP](https://github.com/QuessApp/trusted-geo-service)
- [Ban Service - Serviço para banir usuários](https://github.com/QuessApp/ban-service)
- [Web App - Nossa bela interface Web!](https://github.com/QuessApp/web-app)

[Confira todos os projetos aqui](https://github.com/orgs/QuessApp/repositories)

## Tecnologias

- GO
- RabbitMQ
- AWS SES <br/>
  ...

## Rodando localmente

Clone o projeto

```bash
$  git clone https://github.com/QuessApp/email-service
```

Vá para o diretório do projeto

```bash
$  cd email-service
```

Execute os comandos para iniciar o projeto:

```bash
$ make start
```

ou

```bash
$ ./scripts/start.sh
```

Ao executar o comando acima, ele executará algumas ações, como:

- Irá verificar se o arquivo `.env` existe
- Copia o arquivo `.env.example` para `.env` se ainda não existir
- Execute contêineres oriundos do arquivo `docker-compose.yml` (você precisa executar o Docker no PC)
- Por fim, inicie o projeto com base na propriedade `ENV` do arquivo `.env`

Se você quiser destruir tudo, você pode executar o seguinte comando:

```bash
$ make destroy
```

ou

```bash
$ ./scripts/destroy.sh
```

Ao executar o comando acima, ele executará algumas ações, como:

- Excluir todos os contêineres criados anteriormente
- Excluir pasta `tmp`

## Planos para o futuro

- Escrever testes

## Contribuindo

Contribuições são sempre bem vindas!

Veja [contributing.md](https://github.com/QuessApp/email-service/blob/master/.github/CONTRIBUTING_pt.md) para saber como começar.

Por favor, siga o `código de conduta` desse projeto.

## Autores

- [Caio Augusto (dono & mantenedor)](https://www.github.com/caioaugustoo)

## Suporte

Para suporte, mande um email para caioamfr@gmail.com

## Licença

[MIT](https://choosealicense.com/licenses/mit/)
