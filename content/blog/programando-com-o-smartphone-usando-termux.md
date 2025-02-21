+++
title = "Programando com o smartphone utilizando o Termux"
date = "2025-02-20T20:00:00-00:00"
tags = ["termux", "github", "linux"]
categories = ["artigos"]
banner = "img/banners/floating-smartphone.webp"
+++

## Introdução

Olá pessoal, me chamo Rubens dos Santos e sou um membro do grupo Os Programadores.
Neste artigo iremos aprender a configurar o Termux com git e github, além de algumas dicas para programar no smartphone.

Mas o que é Termux? Termux é um emulador de terminal de código aberto para o sistema operacional Android. Não é o único disponivel, mas o que o difere dos outros aplicativos emuladores de terminal é o fato dele se assimilar a uma distro linux, possuindo um gerenciador de pacotes por exemplo, o que nos possibilita instalar bibliotecas e pacotes.


> **Preciso saber linux para poder utilizar o Termux?**
> Depende. Você precisou saber andar de bicleta para aprender a andar de bicicleta? Paradoxal, certo? Em muitas ocasiões da nossa vida aprendemos no **durante**, no **processo**, na **tentativa e erro**. Então respondendo a pergunta: você pode aprender linux enquanto utiliza o Termux, enquanto sua curiosidade o leva a se questionar o que cada comando faz. Não precisa ler um livro de mil páginas sobre linux nem nada do genêro, ao menos que você queira.

Vamos começar?

## Baixando o Termux

Por uma questão de manutenabilidade da ferramenta eu indico que faça o download do aplicativo pela **F-Droid**. Que é uma loja de aplicativos livres e open-source. Você **não precisa** baixar a F-Droid para baixar o Termux. Basta baixar o Termux acessando [o site da F-droid](https://f-droid.org/pt_BR/packages/com.Termux/) e clicar em **baixar APK**. Após isso basta instalar o APK do Termux em seu smartphone.

## Conhecendo a ferramenta

Essa é a tela inicial do Termux

![termux](/img/conteudos-de-artigos/programando-com-o-smartphone-utilizando-o-termux/termux.webp)

Um ambiente linux em linha de comando. Os links no topo são para acessar a documentação, comunidade e donativo para apoiar o projeto caso tenha interesse. Temos algumas instruções de como utilizar o `pkg` e repositórios adicionais.
O Termux utiliza por padrão o shell `bash`, mas você é livre para instalar o `zsh` ou `fish` em seu ambiente (se você não sabe o que são essas coisas, não precisa se preocupar neste momento)

## Atualizando o sistema e instalando o Git e OpenSSH

O Termux utiliza o `pkg` e `apt` como gerenciadores de pacotes. Na ferramenta ambos são análogos mas utilizarei o `apt` nos exemplos por uma questão didática.
Faremos a atualização dos repositórios e do sistema com o comando abaixo:

```bash
apt update && apt upgrade -y
```

> A flag `-y` é utilizada quando não queremos que um ou mais comandos precisem de confirmação do usuário.

Agora vamos instalar o `openssh` e o `git`. O git é fundamental para nós desenvolvedores versionarmos nosso código. Também precisaremos de um comando do openssh para gerar uma chave ssh que iremos utilizar para vincular nosso host, o Termux, com o github.

Instale os pacotes:

```bash
apt install openssh git
```
## Configurando Git

É importante fazemos duas configurações básicas em nosso git: definir nosso e-mail e usuário do github.
Podemos fazer isto com os comandos abaixo:

```bash
git config --global user.name "seu_usuario"
git config --global user.email "seu_email"
```

## Gerando a chave SSH

A próxima etapa é gerar nossa chave SSH. Isso é realizado pelo comando ssh-keygen que executaremos em seguida:

```bash
ssh-keygen
```

Após a execução do comando vocẽ será questionado sobre:

1) Onde Deseja salvar a chave (vamos deixar em branco, desse modo a ferramenta salva num diretório padrão)
2) Se você quer proteger ela com uma senha (o que eu recomendo, embora não seja obrigatório).

> O diretorio padrão onde a chave é salva se chama `.ssh` (todo diretório linux precedido por um _ponto_ [ . ] é uma pasta oculta) e fica localizado na `$HOME`, que aponta para a pasta do usuário.

3) Dentro da pasta `.ssh` teremos alguns arquivos. A chave ssh que utilizaremos fica dentro do arquivo com final `.pub`. Podemos obter a chave da seguinte maneira:

- Acessamos a pasta
```bash
cd $HOME/.ssh
```
- Listamos os arquivos
```bash
ls
```

- Identificamos o arquivo com final `.pub`, exibimos o conteúdo dele com o comando `cat` e copiamos a saída contendo a chave SSH.
```bash
cat id_xpto.pub
```

> O comando cat permite criar, exibir ou redirecionar o conteúdo de arquivos. Nesse caso estamos utilizando ele para exibir nossa chave para que possamos copiá-la. Podemos realizar todo esse processo descrito acima em um único comando `cat $HOME/.ssh/id_xpto.pub`.

**Lembrando que nossa chave SSH, mesmo sendo protegida por senha nunca deve ser compartilhada com outras pessoas**


## Adicionando nossa chave SSH do Termux no Github

1) Navegaremos até as configurações do nosso Github, na seção de [SSH e GPG keys](https://github.com/settings/keys).

2) Clicaremos em **New SSH key**

3) Escolheremos um titulo para nossa chave. Nesse caso escolhi `termux-ssh-key`. O **key type** podemos manter como `Authentication` e por fim colaremos o conteúdo copiado da nossa chave ssh  no campo **key** e clicaremos em **Add SSH Key**.

![add-new-ssh-key-to-github](/img/conteudos-de-artigos/programando-com-o-smartphone-utilizando-o-termux/add-ssh-key-to-github.webp)


## Verificando a conexão ssh.

A última etapa do processo é verificar se nosso vinculo do Termux com o github via ssh foi bem sucedido, faremos isso por meio do comando abaixo (que deve ser executado no Termux):

```bash
ssh -T git@github.com
```

Se a saída do comando for:

_Hi [Nome do seu usuário]! You've successfully authenticated, but GitHub does not provide shell access._

Parabéns! Agora você pode utilizar o Termux como ferramenta de desenvolvimento podendo trabalhar tanto em seus projetos pessoais como em projetos open-source com o GitHub.

## Dicas Extras

1) Recomendo fortemente estudar a documentação do Termux. Ele é muito mais do que parece ser.

2) Recomendo estudar alguns comandos básicos de linux, a maioria deles possui um manual próprio integrado mas se o idioma for barreira... há livros, videos e todo tipo de conteúdo sobre linux.

2) O Termux é um projeto comunitário. Descobriu um bug? Abra uma issue no repositório deles (caso não exista). Se você descobrir como resolver o bug, melhor ainda: abre um PR.

3) Este último tópico está relacionado ao uso pessoal de cada um. Nós desenvolvedores precisamos de editores de código para programar, há muitos editores que funcionam dentro do terminal. No contexto do Termux isso acaba sendo positivo pois smartphones não possuem a mesma flexibilidade de gerenciar janelas como temos no Desktop/Notebook. Alguns dos editores via linha de comando são o [Nano](https://www.nano-editor.org/), [Vim](https://www.vim.org/), [NeoVim](https://neovim.io/), [Micro](https://micro-editor.github.io/) e [Emacs](https://www.gnu.org/software/emacs/). Minha recomendação — pessoal — para um iniciante é o **Micro**. Apesar dele possuir uma comunidade pequena, seus comandos são muito simples (semelhantes  aos do [VSCode](https://code.visualstudio.com/)) o que o torna amigável e prático, além de suportar plugins assim como o Emacs e o NeoVim. Existem também editores de código na PlayStore (alguns são bem completos). Escolha o que melhor se sentir confortável para utilizar. O importante é aprender.

> “Se eu vi mais longe, foi por estar sobre ombros de gigantes”_ — Sir Isaac Newton
