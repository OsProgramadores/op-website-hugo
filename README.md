# Os Programadores - Projeto do Site
[![Build Status](https://travis-ci.org/OsProgramadores/op-website-hugo.svg?branch=master)](https://travis-ci.org/OsProgramadores/op-website-hugo)

Este repositório contém o projeto do site https://osprogramadores.com, desenvolvido por [Marcelo Pinheiro](https://twitter.com/mpinheir). O site serve como ferramenta de suporte ao grupo OsProgramadores.

Se você tem interesse em programação e assuntos relacionados à área em geral, junte-se a nós no [Telegram](https://t.me/osprogramadores).

A página do grupo foi criada usando o [Hugo](https://gohugo.io/) e uma versão modificada do tema [Universal](http://themes.gohugo.io/theme/hugo-universal-theme/).

## Como contribuir

### Instale o Hugo

Este repositório contém o fonte do website, que usa o [Hugo](http://gohugo.io) para gerar os arquivos em HTML, CSS e javascript a partir de conteúdo em markdown e arquivos de configuração.

1. Instale o [Hugo](https://gohugo.io/) (gerador de websites). Faça o download da versão apropriada para o seu sistema na [Página de Releases](https://github.com/spf13/hugo/releases).
1. Siga as instruções de instalação na página do Hugo.

### Faça um clone do repositório

1. Crie uma conta no GitHub.
1. Faça um *Fork* deste repositório no Github (Botão `Fork` no canto superior direito da tela).
1. Instale o git (CLI) na sua máquina: 
    1. Instruções genéricas: https://git-scm.com/downloads
    1. Debian ou variações (Ubuntu, Mint, etc), use: `sudo apt-get install git`
1. Acesse a linha de comando (Windows) ou abra um terminal (Linux ou OSX).
1. Crie um diretório onde você deseja instalar o projeto do site e faça dele o diretório corrente.
1. Faça um clone do seu fork no github no diretório corrente:
    ```
    git clone --recursive https://github.com/<seu_usuário_no_github>/op-website-hugo.git
    ```
1. Crie um "remote" apontando para o repositório oficial (upstream):
    ```
    git remote add upstream https://github.com/OsProgramadores/op-website-hugo.git
    ```
### Modificando o conteúdo e fazendo o preview

1. **IMPORTANTE**: Antes de *qualquer* modificação (com um repositório limpo), digite `git pull upstream master`. Isso garantirá que a cópia do seu repositório local contem todas as alterações presentes na página inicial.
1. Faça as alterações desejadas, com o seu editor preferido.
1. Digite `hugo -wv` seguido de ENTER. Os arquivos HTML serão gerados no diretório `public`.
1. Visualize suas modificações visitando http://localhost:1313 e certifique-se de que estão corretas (Dica: O Hugo recarrega alterações automaticamente. Se uma mudança for necessária, faça a correção em outra janela e recarregue a página. Não é necessário parar o hugo e rodar novamente).

### Enviando o Pull Request (PR)

Uma vez satisfeito com as suas alterações:
    
1. Verifique que os arquivos corretos foram alterados com `git status`
1. Verifique o diff com `git diff`
1. Faça um commit das alterações com `git commit`. Faça uma descrição breve da alteração na primeira linha do commit (terminado com um ponto). Se maiores informações forem necessárias, use as linhas subsequentes.
1. Faça um push das alterações para o seu repositório com `git push origin master -f`
1. Visite a sua página no github.com (github.com/<seu_usuário_no_github>) e abra um "Pull Request". Esse passo irá enviar um email para os admins com as suas mudanças.
1. Fique atento a emails vindos dos admins, com sugestões de alteração. Os admins pedem encarecidamente que respostas sejam dadas no menor tempo possível. Pedidos sem resposta em 7 dias causarão o fechamento do PR.
1. Uma vez aprovado, visite http://osprogramadores.com após 15 minutos e verifique a sua mudança no ar.

Envie seus comentários para info@osprogramadores.com ou participe de nosso [grupo no Telegram](https://github.com/spf13/hugo/releases).
