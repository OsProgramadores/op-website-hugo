+++
date = "2023-03-09T10:00:00-03:00"
title = "Profiling - Introdução com Java, VisualVM e VSCode"
categories = ["ferramentas"]
tags = ["programação","java", "profiler", "vscode", "visualvm", "debug"]
banner = "img/banners/visualvm_logo_big.webp"
+++

Neste artigo, vamos fazer uma introdução acerca da técnica de `profiling`. Esta técnica é uma forma de análise dinâmica do código, isto é, que verifica e recolhe informações do programa em tempo de execução.

Ela é muito importante, pois nos permite ter informações muito importantes sobre o nosso programa que as vezes são impossíveis de se perceber apenas lendo o código, como a quantidade de vezes que uma função é executada, quais funções que demoram mais para serem executadas, como a memória é utilizada internamente no código e etc.

Apesar de ser uma técnica importante e bastante útil, ela não é bastante difundida nos cursos para iniciantes e muitos desenvolvedores não sabem tirar proveito dessa ferramenta.

Dessa forma, aqui faremos uma breve explicação de como uma ferramenta de profiling funciona utilizando o Java e como ela nos auxilia a entregar um código de qualidade ao final do processo.

## Requisitos Mínimos

Para este artigo, irei utilizar apenas ferramentas open-source, dessa forma, você poderá reproduzir facilmente na tua máquina local:

- Linux. Irei utilizar o Ubuntu 22.04, mas estes processos funcionam em outras distribuições e sistemas operacionais, porém os comandos no terminal poderão ser diferentes dos utilizados aqui.

- [VSCode](https://code.visualstudio.com/). Um editor de textos open-source.

- Java 17
  - No Ubuntu 22.04, para instalá-lo, basta utilizar o comando: `sudo apt install openjdk-17-jdk`

- Qualquer programa de terminal a tua escolha.

## Configurando o VSCode

### GraalVM

Antes de adicionarmos os plugins no VSCode, devemos instalar o GraalVM, porque ele é necessário para utilizarmos o plugin que possui a ferramenta de profiling que iremos utilizar: o VisualVM.

Os passos para instalar o GraalVM no Ubuntu 22.04 são esses:

1. Baixar o arquivo contendo o GraalVM dentro da [página de releases](https://github.com/graalvm/graalvm-ce-builds/releases) que eles possuem no Github.
   1. Lembrem-se de escolher a opção que corresponde a JVM que possui em teu computador, o sistema operacional e a arquitetura. No meu caso, como estou utilizando Java 17 e o Ubuntu 22.04, fiz o download da versão **Linux(amd64)** na coluna **Java 17**

1. Descompactar o arquivo baixado dentro da pasta `Documentos` no teu computador.

## Plugin do VSCode

Pronto, agora que baixamos a pasta do GraalVM, vamos adicionar o plugin necessário no VSCode e configurá-lo:

1. Na parte lateral direita do teu editor, vá no ícone para instalar as extensões, que se parece como uma peça do jogo do Tetris:

![icone-extensoes](/img/conteudos-de-artigos/introducao-profiler-java/icone_extensions-vscode.webp)


2. Na barra de pesquisa, digite **GraalVM Tools for Java** e faça a instalação deste plugin.

3. Ao instalar esse plugin, na barra direita do teu editor irá aparecer um ícone com as letras **Gr**. Clique nele.

4. Ao clicar no ícone do plugin, ele pedirá na parte de cima para fazer o Download e Instalação do GraalVM ou escolher uma existente. Selecione a segunda opção e escolha a pasta do GraalVM que descompactou na pasta de **Documentos** na seção anterior.

![add-graalvm](/img/conteudos-de-artigos/introducao-profiler-java/add_existing_graalvm.webp)

5. Agora, precisamos instalar o componente do VisualVM, que é a ferramenta de profiler que iremos utilizar, basta expandir a opção do GraalVM que você adicionou no plugin e clicar para instalar o componente do VisualVM como na imagem abaixo:

![add-visualvm-component](/img/conteudos-de-artigos/introducao-profiler-java/add-visualvm-plugin.webp)

6. Pronto, ao fim da instalação do componente, você deve ver uma seção chamada **VisualVM** na parte inferior da aba do plugin com algumas opções *Process*, *CPU sampler*, *JFR* e etc.

7. Ao clicar no botão de executar o VisualVM, ele deverá abrir o programa de profiling, que possui esta estética:

![visual-vm-gui](/img/conteudos-de-artigos/introducao-profiler-java/visualvm_gui.webp)

Pronto, agora que configuramos e instalamos todas as ferramentas necessárias, podemos seguir com os passos seguintes, que é a criação do programa em Java que utilizaremos para testar a ferramenta de profiling e o uso dela efetivamente!

## VisualVM

Nota-se que nos passos anteriores instalamos duas ferramentas o GraalVM e o VisualVM. O GraalVM é um projeto que cria uma máquina virtual(JVM) que visa melhorar a performance dos programas em Java e outra linguagens. Não vou entrar em mais detalhes porque aqui utilizamos ele apenas para instalar o plugin do VSCode e o componente do VisualVM, mas vale a pena conhecer este projeto, que é bastante interessante: https://www.graalvm.org/

O VisualVM é a ferramenta de profiling que iremos utilizar, desenvolvida pela Oracle, com sua primeira release sendo lançada no ano de 2008. Ela é uma ferramenta open-source e grátis, distribuída através da lincença GPL e era embutida dentro da própria JVM até a versão 8, agora segue como um projeto separado.

## Programa de Teste

Para testarmos como o VisualVM funciona, vamos criar um programa bastante simples, com apenas uma classe, que executa os seguintes passos:

1. Recebe o path de um arquivo como parâmetro.
   1. Caso o usuário não passe o arquivo como parâmetro, o programa irá enviar uma mensagem de erro.

1. Abre o arquivo que o usuário passou como argumento e percorre ele linha por linha.
   1. Caso o aquivo não exista ou é inválido, ele irá lançar uma mensagem de erro especificando que o arquivo é inválido.

1. Para cada linha do arquivo, ele verifica, utilizando uma regex, se a linha corresponde a um formato de email válido ou não.

1. Caso a linha corresponda a um email com formato válido, iremos escrever no terminal que o email é válido, caso contrário, iremos escrever que o email não é válido.

Pronto, este pequeno programa já nos permite testar como o VisualVM funciona e como podemos utilizá-lo para melhorar nosso código. A primeira versão feita pode ser visualizada logo abaixo:

```
import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.util.regex.Pattern;

public class Main {
    public static void main(String[] args) {
        if (args.length != 1) {
            System.err.println("Deve-se passar o caminho do arquivo com os emails como parametro!");
            System.exit(1);
        }

        try (var fileReader = getFileReader(args[0])) {
            String line;
            while ((line = fileReader.readLine()) != null) {
                boolean isValidEmail = isValidEmail(line);

                printEmailLine(line, isValidEmail);
            }
        } catch (IOException e) {
            System.err.println("Erro enquanto inicializa o arquivo");
        }
    }

    private static void printEmailLine(String line, boolean isValidEmail) {
        if (isValidEmail) {
            System.out.println("O email eh valido: " + line);
        } else {
            System.out.println("O email eh invalido: " + line);
        }
    }

    private static boolean isValidEmail(String line) {
        return Pattern
            .compile("^(?=.{1,64}@)[A-Za-z0-9_-]+(\\.[A-Za-z0-9_-]+)*@[^-][A-Za-z0-9-]+(\\.[A-Za-z0-9-]+)*(\\.[A-Za-z]{2,})$")
            .matcher(line)
            .matches();
    }

    public static BufferedReader getFileReader(String path) throws IOException {
        var file = new File(path);

        if (!file.exists() || !file.canRead() || file.isDirectory()) {
            System.err.println("Arquivo invalido: " + path);
            System.exit(1);
        }

        return new BufferedReader(new FileReader(file, StandardCharsets.UTF_16));
    }
}
```

Para testar este programa, iremos utilizar um dataset que é uma pequena modificação do dataset [Eron](https://www.cs.cmu.edu/~enron/), que é disponível abertamente, sendo utilizado em alguns cursos online, e possui uma lista com mais de 400 mil emails. Ele pode ser encontrado neste [gist](https://gist.github.com/anthonylgf/f5d76ecc54199a89c5073772ca7e35ce).

Para rodar esse programa, basta criar um arquivo chamado **Main.java** em tua máquina local, copiar e colar o programa acima, baixar o arquivo com o dataset, ir na pasta onde criou o arquivo e rodar o seguinte comando no terminal:

```shell
$ java Main.java [caminho-para-o-dataset]/emails_400k.csv
```

Lembre-se que você deve utilizar o Java 17, ou o comando acima pode dar erro durante a execução.