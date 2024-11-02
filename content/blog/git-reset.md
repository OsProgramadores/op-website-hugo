+++
title = "Help! Múltiplos desafios no meu PR!"
date = "2024-11-02T13:00:00-07:00"
tags = ["programação", "git"]
categories = ["git"]
banner = "img/banners/git.webp"
+++

## Help! Múltiplos desafios no meu PR!

Este artigo contem uma explicação para uma pergunta frequente: "Por que o meu
PR tem múltiplos desafios se eu já enviei os anteriores?"

TL;DR: Este problema acontece quando o usuário começa um novo desafio sem
sincronizar o repositório local com o conteúdo do repositório _upstream_
(nesse caso, osprogramadores).

## Conceitos importantes

1. **upstream** se refere ao repositório principal do os-programadores.
   Usuários não tem acesso a push a esse repositório.
1. **origin** Aponta para o seu _fork_. O usuário dono do repositório tem
   acesso completo a esse repositório e pode usar `git push` para enviar
   modificações.
1. **Pull Request (PR)** é uma requisição de modificação do repositório
   upstream a partir de modificações no seu fork. O PR é aprovado pelos
   administradores do os-programadores, que então passa a conter as suas
   modificações.

## Explicação detalhada

Imagine uma situação inicial onde você faça o clone do
`osprogramadores/op-desafios` (upstream), e que existam apenas dois commits lá.
Usaremos números simples ao invés de hashes para facilitar a compreensão:

```
002 Desafio de outro alguém
001 Desafio de alguém
```

Depois do `git clone`, o repositório local fica igual ao upstream (dica:
use `git log --oneline --all` para verificar):

```
002 Desafio de outro alguém
001 Desafio de alguém
```

Nesse momento, você começa a trabalhar em um novo desafio (não importa qual o
branch, vamos supor que seja `master`, para efeito de simplicidade). Depois de
um commit inicial e uma correção, o repositório local agora contém os seguintes
commits:

```
301 Correção do meu novo desafio.
300 Meu novo desafio
002 Desafio de outro alguém
001 Desafio de alguém
```

Novamente, observe que os números de commit acima são uma simplificação. Usamos
300 em diante para commits feitos a partir de uma modificação local. Na prática
são hashes. Use `git log --oneline --all` para ver a lista de commits no
repositório.

Com tudo confirmado, é hora de fazer um push para o seu fork e criar um PR. O
comando `git push origin/master --force` envia as modificações para o seu fork
no github. A partir daí, cria-se um PR, que basicamente é um diff entre o seu
fork e o upstream.

O PR é então analisado pelos administradores do repositório. Se tudo estiver
certo, os administradores aprovarão e farão o "merge" do repositório.

No caso do os-programadores, os administradores optaram por usar
**rebase/squash**, para manter o repositório linear. Em termos simples, isso
quer dizer que o commit 300 e 301 (vindos do seu fork) vão ser reescritos
como apenas *um* commit antes do merge. Após o rebase/squash, o upstream fica
da seguinte forma:

```
003 Desafio do Gabriel  <-- Nome do commit mudado. Inclui 300 e 301!
002 Desafio de outro alguém
001 Desafio de alguém
```

Repare que o commit 003 agora no upstream é a combinação de dois commits seus
(300 e 301). O upstream não tem a mínima ideia do que sejam 300 e 301, já que
eles não existem lá.

Nesse momento, você começa OUTRO desafio, e se esquece de fazer o _reset_
do repositório! No seu segundo desafio, cria mais um commit (400). O seu
repositório local agora tem a seguinte estrutura de commits:

```
400 Mais um desafio novinho!       <-- Novo desafio
301 Correção do meu novo desafio.
300 Meu novo desafio
002 Desafio de outro alguém
001 Desafio de alguém
```

Um novo `git push origin/master --force` faz o upload do repositório local
para o seu fork (que naturalmente fica idêntico ao repositório local). Quando
o novo PR request é criado, algo interessante acontece: O upstream só tem os
commits 001, 002 e 003, logo ele ignora o 001 e 002 no seu fork (ele já sabe
que tem esses, eles não criam um diff), mas dessa vez o seu  repositório também
contem os commits 300, 301, e 400. O resultado? **No PR aparecem os três
commits e os dois arquivos**, mostrando dois desafios no PR quando apenas um
foi adicionado.

## Evitando o problema

Lembre-se que usamos rebase+squash no merge dos PRs para manter o repositório
principal linear. O efeito colateral disso é que essa operação **sempre cria
um commit novo a partir dos seus commits**. Logo, **sempre use `git reset
upstream master` antes de começar a trabalhar em um novo desafio**, e sempre
use um `git pull upstream master -r` para fazer um rebase das suas
modificações antes de enviar o PR.

Complicado? Vamos por partes:

Antes de começar um novo desafio (ou qualquer outra modificação), tenha
certeza de que o seu repositório está idêntico ao upstream/master:

```bash
git remote update
git reset upstream/master --hard
```

Trabalhe normalmente. Antes de enviar o push, execute:

```bash
git pull -r upstream master
```

Isso fará um *rebase* das suas mudanças em cima do upstream.  Basicamente ele
lê o upstream para dentro do seu repositório e reaplica suas mudanças no topo
do upstream.

É importante nunca usar merge (`git pull` sem `-r` ou `git merge`), já que ele
cria um merge commit que vai complicar a visualização do seu repositório.

## Como resolver o problema em um PR já existente?

Se um PR problemático já foi criado, a solução é relativamente simples:

1. Use `git remote update` para atualizar as cópias locais dos repositórios
   remotos.
1. Use `git reset upstream/master`. Esse comando joga o HEAD do seu repositório
   para o HEAD do upstream/master (atenção: **não use `--hard` aqui**).
1. Use o comando `git status`. É provável que várias alterações não
   relacionadas com as suas modificações existam (arquivos deletados, criados,
   etc). Use `git restore` em todos os arquivos que não rejam relacionados com
   o seu PR.
1. Adicione as suas modificações novamente com `git add` (elas estão intactas,
   mas precisam ser adicionadas após o `git reset`)
1. Faça o commit novamente com `git commit`.
1. Faça o upload para o seu fork com `git push origin master --force`.
1. Deixe um comentário no PR (github) informando que o problema foi corrigido.
   **Não é necessário fechar o PR e/ou criar outro.**

## Se tudo falhar?

Pergunte a um dos admins no grupo [os-programadores](http://t.me/os-programadores).
