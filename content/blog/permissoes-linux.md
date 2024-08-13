+++
title = "Breve Introdução e Permissões no Linux"
date = "2024-08-12T23:27:41-03:00"
tags = ["linux","guia","dicas"]
categories = ["linux"]
banner = "img/banners/TODO.webp"
+++

Quanto roda os comandos seguintes, as permissõe das pastas e arquivos aparecem:

```sh
$ pwd
/tmp/teste
$ ls -la
total 0
drwxr-xr-x  2 usuario_1 usuario_1  80 Aug 13 00:29 .
drwxrwxrwt 30 root      root      820 Aug 13 00:28 ..
-rw-r--r--  1 usuario_1 usuario_1   0 Aug 13 00:29 arquivo1.txt
-rw-r--r--  1 usuario_1 usuario_1   0 Aug 13 00:29 imagem.png
```

As linhas como `-rw-r--r--  1 usuario_1 usuario_1` descrevem as permissões do arquivo. Com essa linha, é possível saber se seu usuário pode ler, escrever e/ou executar os arquivos.

Você pode interagir com isso usando o comando chmod (CHange MODe)

R: read (se você pode ler o arquivo)
W: write (se você pode escrever o arquivo)
X: execute (se você pode executar o arquivo)

Você tem permissão de usuário, grupo, e outros, é oq a parte de cima mostra, e é por isso q você vê 3 vezes esse padrão rwx rwx rwx

Basicamente um arquivo que tem rw- em outros (terceira coluna), pode ser escrito por qualquer um. E por padrão por questões de segurança, geralmente é sempre configurado pra 0

Exercícios:

Responda oq cada permissão faz da seguinte forma: "Permissão de usuário permite..., Permissão de grupo permite... Permissão de outros permite..." 

```
rw- r-- ---
r-- --- ---
rw- --- -w-
r-x r-x r-x
```

# Águas de bases numéricas a frente

RWX são binários, então 111 significa rwx, 100 significa r--, 101 significa r-w e por aí vai

Podemos ler esses números binários, então:

4 - 100 r--
6 - 110 rw-
7 - 000 rwx

Podemos então escrever permissões assim 467

Significa r-- para usuário, rw- para grupo rwx para todos

~ Trivia ~

O que a permissão 666 faz? É perigosa?

Todo hacker (lembre-se, hacker não é cracker!) tem que no MÍNIMO saber ler isso, e os script kiddies já espanam o parafuso na primeira página de manpages do chmod


