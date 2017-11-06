# Os Programadores - Projeto do Site
[![Build Status](https://travis-ci.org/mpinheir/op-website-hugo.svg?branch=master)](https://travis-ci.org/mpinheir/op-website-hugo)

Este repositório contém o projeto do site https://osprogramadores.com, desenvolvido por [Marcelo Pinheiro](https://twitter.com/mpinheir). O site serve como ferramenta de suporte ao grupo OsProgramadores.
Se você tem interesse em programação e assuntos relacionados à área em geral, junte-se a nós no [Telegram](https://t.me/osprogramadores).

A página do grupo foi criada usando o [Hugo](https://gohugo.io/) e o tema [Universal](http://themes.gohugo.io/theme/hugo-universal-theme/).

## Instruções para executar o clone do Projeto do site
  1. Crie uma conta no GitHub.
  2. Instale o git na sua máquina: https://git-scm.com/downloads.
  3. Acesse a linha de comando (Windows) ou abra um terminal (Linux ou OSX).
  4. Crie um diretório onde você deseja instalar o projeto do site.
  5. Digite `git clone --recursive https://github.com/OsProgramadores/OsProgramadores-ProjetodoSite.git`

## Instruções para re-criar o website

Esse repositório contem o "fonte" do website, que é baseado num gerador de websites a partir de templates chamado Hugo. Para re-gerar o website ou submeter alterações, siga as instruções abaixo:

  1. Instale o [Hugo](https://gohugo.io/) (gerador de websites). Faça o download da versão apropriada para o seu sistema na [Página de Releases](https://github.com/spf13/hugo/releases) do Hugo.
  2. Siga as instruções de instalação na página do Hugo.
  3. Mude para o diretório onde foi feito o `git clone` desse projeto no passo 5 do processo de anterior.
  4. Digite `hugo` seguido de ENTER. Os arquivos HTML serão gerados no diretório `public`.
  5. Opcional: O Hugo também permite visualização instantânea do site gerado. Para isso, use `hugo server -wv` e aponte o browser para http://localhost:1313.

Envie seus comentários para info@osprogramadores.com ou participe de nosso [grupo no Telegram](https://github.com/spf13/hugo/releases).
