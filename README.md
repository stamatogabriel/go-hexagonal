# Go-Hexagonal

Projeto realizado dentro do Curso Full Cycle com o intuito de estudar e entender a lógica por trás de uma arquitetura hexagonal.

É uma implementação simples de um serviço de controle de produtos, onde criamos um interface com as regras que de implementação do produto e a abstração do serviço para conexão com os métodos de persistência (banco de dados ou File Systems).

## O que é Arquitetura Hexagonal
Também chamada de Ports and Adapters, a arquitetura hexagonal é uma forma de organizar o código em camadas, cada qual com a sua responsabilidade, tendo como objetivo isolar totalmente a lógica da aplicação do mundo externo. Ou seja, define limites e proteção nas regras da aplicação, mantendo os serviços o máximo possível desacoplados e componentizados, facilitando assim manutenção, melhoria, crescimento e escalabilidade do projeto.

## Tecnologias utilizadas
- GoLang

## Como executar

### Necessário
- Ter Docker e docker-compose instalados

### Rodando a aplicação
- abra o terminal na raiz do projeto
- execute o comando `docker-compose up -d`