+++
title = "Por onde começar?"
date = "2019-03-12T08:26:00-07:00"
tags = ["programação"]
categories = ["linguagens"]
banner = "img/banners/background-code.webp"
+++

Uma pergunta frequente de vários participantes do [OsProgramadores no Telegram](https://t.me/osprogramadores) é: "Por onde eu começo?" e "como posso me preparar para conseguir o primeiro estágio ou o primeiro emprego na área de programação?" 

Para começar com programação é recomendável usar um computador para criar e testar programas. A prática é essencial para reforçar o que foi aprendido através de vídeos e livros. 
Caso você não tenha um computador, existem vários programas que você pode usar mas recomendamos a instalação e uso da [App do Replit](https://replit.com/mobile).

> **[Recomendamos a leitura deste artigo para lhe ensinar alguns conceitos essenciais que serão muito úteis nesta sua jornada para se tornar um programador.](https://osprogramadores.com/blog/2024/02/29/conceitos-essenciais-para-come%C3%A7ar/)**

Qualquer sistema operacional pode ser usado no aprendizado, [Windows](https://www.microsoft.com/en-ca/windows), [Linux](https://en.wikipedia.org/wiki/Linux) ou [macOS](https://en.wikipedia.org/wiki/MacOS). Use o que você tem acesso para aprender. Também não se preocupe com qual editor de programas ou [IDE](https://en.wikipedia.org/wiki/Integrated_development_environment) usar. Escolha algo simples e se concentre em aprender a programar. Se você não souber qual editor de programas usar, recomendamos o [Visual Studio Code](https://code.visualstudio.com/)

Escolha uma linguagem de programação como [Python](https://wiki.python.org.br/DocumentacaoPython), [JavaScript](https://www.w3schools.com/js/default.asp) ou outra de seu interesse e prossiga com o aprendizado. Se concentre em aprender apenas uma linguagem pois tentar aprender vários tópicos ao mesmo tempo só irá gerar frustrações. 

Se o seu interesse é desenvolvimento front-end, recomendamos seguir a trilha de [HTML + CSS e depois JavaScript do Freecodecamp](https://www.freecodecamp.org/portuguese/learn/).

Se seu interesse for em outras linguagens, na seção de [Links do site Osprogramadores](https://osprogramadores.com/links/) você encontrará tutoriais e documentações que irão ajudar no aprendizado de várias linguagens de programação.

Recomendamos que você assista dois videos que estão no Canal do OsProgramadores no Youtube:
1-[Introdução ao Terminal do Linux](https://www.youtube.com/watch?v=CFWttwWZSAQ&t=1s)
2-[Usando a busca do Google](https://www.youtube.com/watch?v=7Yi7jlbbbsA)

Usaremos Python nos exemplos que serão apresentados abaixo.

Caso o Python não esteja instalado no seu computador, siga as instruções para instalação em um sistema [Windows](https://python.org.br/instalacao-windows/), [Linux](https://python.org.br/instalacao-linux/) ou [macOS](https://python.org.br/instalacao-mac/).

Caso você não deseje instalar o Python em seu computador, recomendamos a criação de uma conta grátis no [Repl.it](https://repl.it/) onde você poderá utilizar um ambiente de edição e execução de programas Python sem a necessidade de instalar software em seu computador.

Vamos começar a programar?

Abra o terminal no Linux ou o [prompt no Windows](https://tecnoblog.net/responde/7-maneiras-de-abrir-o-prompt-de-comando-no-windows-10-e-11/)

Digite ```python``` no terminal ou prompt e pressione a tecla enter para carregar o interpretador interativo da linguagem Python. 

**Note que**:
1. Em alguns sistemas o comando a ser usado pode ser o python3 ao invés de python.
2. Para sair do interpretador python e retornar ao termina/prompt entre  com o seguinte comando e pressione enter:

```python
exit()
```

Comece criando um programinha conhecido tradicionalmente como "Olá Mundo". Imprima uma mensagem simples. Veja um exemplo do programa Hello World em Python abaixo. Pressione a tecla enter ao final de cada linha digitada e note a reação do interpretador do Python.

```python
print("Olá Mundo.")
```

Uma vez entendido o funcionamento do programa "Olá Mundo" você poderá criar um programa com duas instruções, conforme o exemplo abaixo:

```python
print("Olá Mundo.")
print("Escreva algo aqui.")
```

Uma vez entendido o conceito de sequência de instruções do exemplo acima, o próximo passo é usar variáveis. Crie um programa conforme o exemplo abaixo e atribua o valor 0 (zero) a uma variável chamada _a_ e exiba o valor de _a_ na tela do seu computador.

```python
a = 1
print(a)
```

Após a execução do programa anterior, pesquise como você pode adicionar 1 ao valor da variável `a` e exiba o novo valor de `a` na tela do seu computador. 

```python
a = 0
print(A)
a = a + 1
print(a)
```

A seguir, insira uma segunda variável, chamada `b`, no seu programa e imprima o valor de `a + b`. 

```python
a = 0
print(a)
a = a + 1
print(a)
b = 35
a = a + b
print(a)
```

O próximo passo seria usar o comando `for` para repetir uma sequência de linhas do seu programa várias vezes. Sugestão: Imprima uma mensagem com o seu nome 10 vezes.

Exemplo:

```python
for x in range(10):
  print("Escreva o seu nome aqui")
```

A seguir vamos usar o comando `IF` para tomar uma decisão. Pesquise o que é decisão e como funcionam comandos do tipo `IF THEN ELSE`. Faça um programa simples usando `IF THEN ELSE` e procure entender como funciona. 

Exemplo:

```python
a = 5
if a == 5:
    print ("O valor de a = 5")
else:
    print ("Jamais executarei esta linha neste programa")
```

Uma vez que esta parte básica tenha sido entendida, você pode partir para projetos mais complexos pois o que foi descrito acima é o básico de qualquer programa.

Recomendamos que você faça os [desafios do grupo](https://osprogramadores.com/desafios/).

Caso tenha dificuldades com os exercícios, faça perguntas no [grupo no telegram](https://t.me/osprogramadores)


