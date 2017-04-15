+++
title = "Contribuindo para o site"
date = "2017-04-14T23:40:00-07:00"
tags = ["osprogramadores"]
categories = ["git"]
banner = "img/banners/octocat.png"
+++

## Como contribuir para o site?

A melhor maneira de contribuir com um artigo ou correção para o site é criar um
*Pull Request* no github. *Pull Request*, ou simplesmente "PR" é o mecanismo
usado pelo github para enviar modificações a serem integradas em um repositório.

O objetivo desse artigo é sugerir um fluxo de trabalho *extremamente
simplificado* que permita usuários com pouca experiência enviar um PR para os
administradores de "osprogramadores.com".

## Convenções básicas

Em todos os exemplos, o comando é mostrado junto com a saída esperada. O comando
a ser digitado é precedido pelo simbolo "$" (prompt default do Linux).

Em todos os exemplos nesse artigo, troque "USUÁRIO" pelo seu usuário no github.

## O que é preciso

1. Uma conta no [github](http://github.com). Acesse o site e crie sua conta
   gratuitamente. É importante lembrar o seu usuário no github.com, pois ele
   será usado através desse tutorial.

2. Instale o [git](http://git-scm.com) no seu computador.

Apesar de várias GUIs existirem para o git, esse artigo foca no uso de linha
de comando.

## Passo-a-passo

### Crie um fork do repositório principal

O primeiro passo é a criação de um fork do repositório principal de
osprogramadores.com. Um fork é basicamente uma cópia do site inteiro, dentro
da sua conta no github (onde você tem acesso de gravação).

1. Faça o login no github.com
2. Navegue para o website principal do grupo:
   https://github.com/OsProgramadores/op-website-hugo
3. No canto superior direito da tela, clique o botão `Fork`
4. Após alguns segundos, o browser será redirecionado para a cópia do
   repositório osprogramadores.com, usando a sua conta (Observe a URL, que agora
   deve ser algo como https://github.com/USUÁRIO/op-website-hugo, onde "USUÁRIO"
   é o seu usuário no github.

**IMPORTANTE**: Normalmente, só é necessário fazer o fork uma vez, mas devido a
mudanças recentes no nome do repositório principal, recomendamos que todos
refaçam o fork. Se você já tem um fork anterior, abra a página do repositório na
sua conta, clique no botão `Settings` e escolha `Delete This Repository`. Refaça
o fork seguindo as instruções acima.

### Crie um clone do repositório no seu computador

Crie um diretório no seu computador onde o repositório será clonado. Entre no
diretório criado e digite:

```
$ git clone git@github.com:USUÁRIO/op-website-hugo.git

Cloning into 'op-website-hugo'...
remote: Counting objects: 418, done.
remote: Total 418 (delta 0), reused 0 (delta 0), pack-reused 418
Receiving objects: 100% (418/418), 3.11 MiB | 3.07 MiB/s, done.
Resolving deltas: 100% (133/133), done.
Checking connectivity... done.

```

### Verifique o status do seu repositório

Entre no diretório `op-website-hugo` e verifique o status do repositório com:

```
$ git status
On branch master
Your branch is up-to-date with 'origin/master'.

nothing to commit (use -u to show untracked files)
```

Observe a Mensagem "Your branch is up-to-date with 'origin/master'". Essa
mensagem indica que o seu repositório contém uma cópia atualizada do repositório
original.

### Criando uma nova origem

O git permite a sincronização com vários repositórios. No nosso caso, a versão
original será *lida* do repositório OsProgramadores/op-website-hugo, e *gravada*
na sua cópia (fork) embaixo da sua conta no github (USUÁRIO/op-website-hugo):

```
$ git remote add upstream https://github.com/OsProgramadores/op-website-hugo.git

$ git remote -v
origin  git@github.com:marcopaganini/op-website-hugo.git (fetch)
origin  git@github.com:marcopaganini/op-website-hugo.git (push)
upstream        https://github.com/OsProgramadores/op-website-hugo.git (fetch)
upstream        https://github.com/OsProgramadores/op-website-hugo.git (push)
```

O comando `git remote -v` mostra todos as as "origens".

### Crie um branch de trabalho

O git funciona com o conceito de branch, que é uma versão alternativa local dos
dados no repositório. É boa prática manter todas as alterações num branch
separado (fora do branch "master"). No comando `git status` acima, pode-se ver
que o branch default "master" foi criado automaticamente. Crie um outro branch
para as suas contribuições. Qualquer nome é aceitável. No nosso caso, usaremos
o nome "contri":

```
$ git checkout -b contri origin/master
Branch contri set up to track remote branch master from origin.
Switched to a new branch 'contri'

$ git branch
* contri
  master
```

O primeiro comando cria um novo branch, e o segundo mostra o branch corrente com
um asterisco a esquerda (e normalmente, em cor verde).

### Atualize o seu diretório de trabalho

É sempre importante trabalhar com informações atualizadas do repositório
principal. Antes de fazer alterações, sempre atualize o seu diretório de
trabalho com:

```
$ git pull
Already up-to-date.
```

A mensagem "Already up-to-date." indica que o seu diretório local já está
atualizado em relação ao repositório principal. Se modificações existirem
no repositório principal, elas serão reaplicadas no seu diretório local.

Também é possível que o git precise fazer um "merge commit" das mudanças no
repositório remoto com as suas mudanças locais. Neste caso, ele abrirá um
editor indicando que um commit está em progresso. Nesses casos, basta salvar
o arquivo e sair do editor para que a atualização prossiga.

### Modificando um arquivo

Para modificar um arquivo, edite-o diretamente usando o seu editor de textos
preferido. Use o comando `git status` frequentemente para observar o estado
atual do seu repositório local (exemplo com o arquivo README.md modificado):

```
$ git status
On branch contri
Your branch is up-to-date with 'origin/master'.

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        modified:   README.md

no changes added to commit (use "git add" and/or "git commit -a")
```

### Gravando as modificações no repositório local

Quando estiver satisfeito com as alterações, use o comando `git add -A` e `git
commit` para gravar todas as alterações no repositório local:

```
$ git add -A
$ git commit
```

O comando `git commit` abre um editor de textos onde a descrição do commit
(atualização) deve ser digitada. A primeira linha da mensagem de commit deve
conter uma breve descrição da mudança, terminada com um ponto. Se uma descrição
mais detalhada for necessária, adicione uma linha em branco seguida de um ou
mais parágrafos com a descrição mais longa. Exemplo:

```
Correções ortográficas no arquivo README.md.

- Removidos erros de digitação.
- Acertos na formatação.
```

Idealmente, um commit deve ser "atômico", ou seja, conter todas as alterações
necessárias ao funcionamento de uma determinada mudança. Se uma alteração for
necessária após o commit, edite o arquivo novamente e repita o processo acima.
Os administradores do osprogramadores.com se encarregarão de agregar todos os
commits relevantes em um só durante o Pull request.

### Enviando para o github

Uma vez que as suas modificações estejam completas, envie o repositório para a
sua conta no github:

```
$ git push origin contri
Counting objects: 5, done.
Delta compression using up to 8 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 280 bytes | 0 bytes/s, done.
Total 3 (delta 2), reused 0 (delta 0)
remote: Resolving deltas: 100% (2/2), completed with 2 local objects.
To git@github.com:USUÁRIO/op-website-hugo.git
   d4759b8..e50b436  contri -> contri
```

Observe como o comando push recebe o nome da origem (origin) e do branch
corrente ("contri", use `git branch` para verificar). Modifique o nome de acordo.
Não use o branch "master" para fazer pushes!

### Criando um Pull Request

O Pull Request lê as modificações feitas no seu repositório no github e envia as
diferenças para os administradores de osprogramadores.com. Basicamente, é como
você pede que as suas alterações sejam consideradas para inclusão no repositório
principal.

1. Faça o login no github.
2. Na tela principal, observe uma tarja amarela mostrando o seu novo branch
   (contri) e um botão verde indicando `Compare & Pull request`. Pressione
   este botão.
3. Verifique se as mudanças estão corretas, assim como as descrições.
4. Pressione o botão `Create Pull Request`.

Os administradores do osprogramadores.com irão receber um email com as suas
modificações. Uma cópia será enviada para o seu email. Os administradores podem
fazer perguntas sobre a sua mudança e requisitar novas mudanças. Se mudanças
forem necessárias, repita o ciclo acima.

### Limpeza do repositório

Uma vez que o Pull Request **tenha sido aprovado**, remova o branch do repositório local e
remoto com:

```
$ git checkout master
Switched to branch 'master'
Your branch is up-to-date with 'origin/master'.

$ git branch -D contri
Deleted branch contri (was e50b436).

$ git push --delete origin contri
To git@github.com:marcopaganini/op-website-hugo.git
 - [deleted]         contri

```

Para fazer outras modificações, não é necessário fazer outro fork (no github) ou
git clone. Basta recomeçar o ciclo acima a partir da criação de um novo branch.
