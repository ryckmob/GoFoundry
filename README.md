# GoFoundry

<p align="center">
  <img src="https://raw.githubusercontent.com/ryckmob/GoFoundry/main/media/logo.png" alt="GoFoundry Logo" width="200">
</p>

## Por que este projeto é útil

Chega de perder tempo pensando como vai organizar estruturalmente a arquitetura do projeto.

GoFoundry nasce para resolver exatamente esse gargalo inicial que todo projeto em Go enfrenta. Você quer começar a codar, mas perde horas decidindo pastas, nomes, padrões e convenções.

Resultado: Menos decisão inútil, Mais código entregue, Arquitetura consistente desde o primeiro passo;

<div style="background-color:#eef4ff; border-radius:8px; padding:12px 16px; border:1px solid #c7d7ff; font-family:Arial, sans-serif;">
  <strong>Atenção</strong><br>
  Todo app criado com o GoFoundry precisa ser registrado manualmente em <code>internal\http\routes.go</code>.<br>
  Se não importar o app e chamar o <code>RegisterRoutes</code>, nenhuma rota será exposta, mesmo com o servidor rodando.
</div>


GoFoundry is a command line tool written in Go designed to speed up backend development. It creates the project structure, generates apps, and automatically writes full CRUD code from a single command. You define the app name and struct fields, and it generates models, handlers, services, routes, and database integration following a consistent pattern. The goal is to eliminate repetitive work and human error. Instead of wasting time on boilerplate code, you focus on business logic. Pure productivity, no hidden magic.

<div style="background-color:#eef4ff; border-radius:8px; padding:12px 16px; border:1px solid #c7d7ff; font-family:Arial, sans-serif;">
  <strong>Warning</strong><br>
  Every app created with GoFoundry must be manually registered in <code>internal\http\routes.go</code>.<br>
  If you don’t import the app and call <code>RegisterRoutes</code>, no routes will be exposed, even if the server is running.
</div>


## Como instalar

Clonar o repositório

`git clone https://github.com/ryckmob/GoFoundry.git`

Entrar na pasta do projeto

`cd GoFoundry`

Baixar dependências

`go mod tidy`

Gerar o executável

`go build -o GoFoundry.exe .`

Opcionalmente adicionar o executável ao PATH do Windows para usar em qualquer pasta.

## Modos de usar

Criar um novo projeto

`GoFoundry.exe new nome-do-projeto`

Criar um novo app dentro de um projeto existente

`GoFoundry.exe app nome-do-app`

Simples assim.
Você executa o comando, o GoFoundry cria a estrutura e você começa a programar.

