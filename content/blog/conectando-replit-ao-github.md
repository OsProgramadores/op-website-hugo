+++
title = "Como Conectar Replit ao GitHub"
date = "2023-09-19T07:20:00-00:00"
tags = ["replit", "github"]
categories = ["artigos"]
banner = "img/banners/replitGit.jpg"
+++

# Como Conectar seu Projeto do Replit à sua Conta no GitHub

Conectar seu projeto no Replit à sua conta no GitHub por meio de uma chave SSH é uma maneira eficiente de gerenciar seus repositórios remotamente. Com esta configuração, você poderá realizar commits e push de código para seus repositórios GitHub diretamente do Replit. Neste tutorial, vamos orientá-lo passo a passo sobre como realizar essa configuração.

1. Primeiro, acesse o [site do Replit](https://replit.com/~).

2. No Replit, clique no botão "+ Create Repl" para criar um novo projeto.

3. Escolha o modelo "Blank Repl" e dê um nome ao seu projeto.

4. Se o projeto criado incluir um README.md padrão que não será necessário, você pode excluí-lo.

5. Na aba "Console", localizada à direita da interface, execute os seguintes comandos para configurar sua conta Git:

```
$ git config --global user.name "Seu Nome"
$ git config --global user.email "seu_email@example.com"
```

6. Instale o pacote OpenSSH, no seu ambiente Replit com o seguinte comando:

```
$ nix-env -iA nixpkgs.openssh
```

7. Agora, crie uma nova chave SSH executando o seguinte comando:

```
$ ssh-keygen -t rsa -b 4096 -C "seu_email@example.com"
```

Quando for requisitado, basta pressionar a tecla "Enter" para confirmar a localização padrão onde a chave será salva.

Uma vez que você está gerando uma chave privada em um servidor que não está sob seu controle, sugerimos que restrinja o uso desta chave somente a projetos de aprendizado.

Você também será convidado a definir uma senha para a sua chave. É altamente aconselhável escolher uma senha robusta nesta ocasião.

8. Imprima a chave SSH pública no terminal com o seguinte comando:

```
$ cat ~/.ssh/id_rsa.pub
```

Copie a chave SSH pública gerada, pois você a usará no próximo passo.

9. Adicione a Chave SSH ao GitHub
- Faça login na sua conta do GitHub.
- Clique na sua foto de perfil no canto superior direito e vá para "Settings".
- No menu à esquerda, clique em "SSH and GPG keys".
- Clique em "New SSH key".
- Cole a chave SSH pública que você copiou anteriormente na caixa "Key".
- Dê um nome descritivo à chave no campo "Title".
- Clique em "Add SSH key".

Pronto! Sua chave SSH agora está configurada e vinculada à sua conta no GitHub. Você pode usar essa configuração para fazer push e pull de código de forma segura entre o Replit e o GitHub.