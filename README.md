# GoFoundry

<p align="center">
  <img src="https://raw.githubusercontent.com/ryckmob/GoFoundry/main/media/logo.png" alt="GoFoundry Logo" width="200">
</p>

## Por que este projeto é útil

Chega de perder tempo pensando como vai organizar estruturalmente a arquitetura do projeto.

GoFoundry nasce para resolver exatamente esse gargalo inicial que todo projeto em Go enfrenta. Você quer começar a codar, mas perde horas decidindo pastas, nomes, padrões e convenções.

Este projeto é inspirado nos modelos MVC e MTV, porém com uma abordagem pensada para Go. Nada de copiar frameworks de outras linguagens. Aqui a estrutura respeita simplicidade, composição e organização clara, sem magia escondida.

Resultado
Menos decisão inútil
Mais código entregue
Arquitetura consistente desde o primeiro commit

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

