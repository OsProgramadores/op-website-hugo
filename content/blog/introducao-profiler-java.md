+++
date = "2023-03-09T10:00:00-03:00"
title = "Profiling - Introdução com Java, VisualVM e VSCode"
categories = ["ferramentas", "linguagens"]
tags = ["programação","java", "profiler", "vscode", "visualvm", "debug", "artigos"]
banner = "img/banners/visualvm_logo_big.webp"
+++

Neste artigo, vamos fazer uma introdução acerca da técnica de `profiling`. Esta técnica é uma forma de análise dinâmica do código, isto é, que verifica e recolhe informações do programa em tempo de execução.

Ela é muito importante, pois nos permite ter informações sobre o nosso programa que às vezes são impossíveis de se perceber apenas lendo o código, como a quantidade de vezes que uma função é executada, quais funções que demoram mais para serem executadas, como a memória é utilizada internamente no código e etc.

Apesar de ser uma técnica importante e bastante útil, ela não é bastante difundida nos cursos para iniciantes e muitos desenvolvedores não sabem tirar proveito dessa ferramenta.

Dessa forma, aqui faremos uma breve explicação de como uma ferramenta de profiling funciona utilizando o Java e como ela nos auxilia a entregar um código de qualidade ao final do processo.

## Requisitos Mínimos

Para este artigo, irei utilizar apenas ferramentas open-source, dessa forma, você poderá reproduzir facilmente na tua máquina local:

- Linux. Irei utilizar o Ubuntu 22.04, mas estes processos funcionam em outras distribuições e sistemas operacionais, porém os comandos no terminal poderão diferir dos utilizados aqui.

- [VSCode](https://code.visualstudio.com/). Um editor de textos open-source.

- Java 17
  - No Ubuntu 22.04, para instalá-lo, basta utilizar o comando: `sudo apt install openjdk-17-jdk`

- Qualquer programa de terminal desejado(Konsole, Tlix, Terminator, Gnome Terminal, etc.)

## Configurando o VSCode

### GraalVM

Antes de adicionarmos os plugins no VSCode, devemos instalar o GraalVM, porque ele é necessário para utilizarmos o plugin que possui a ferramenta de profiling que iremos utilizar: o VisualVM.

Os passos para instalar o GraalVM no Ubuntu 22.04 são esses:

1. Baixar o arquivo contendo o GraalVM dentro da [página de releases](https://github.com/graalvm/graalvm-ce-builds/releases) que eles possuem no Github.
   1. Lembrem-se de escolher a opção que corresponde a JVM que possui em teu computador, o sistema operacional e a arquitetura. No meu caso, como estou utilizando Java 17 e o Ubuntu 22.04, fiz o download da versão **Linux(amd64)** na coluna **Java 17**

1. Descompactar o arquivo baixado dentro da pasta `Documentos` no teu computador.

## Plugin do VSCode

Pronto, agora que baixamos a pasta do GraalVM, vamos adicionar o plugin necessário no VSCode e configurá-lo:

1. Na parte lateral direita do teu editor, vá ao ícone para instalar as extensões, o que se parece como uma peça do jogo do Tetris:

![icone-extensoes](/img/conteudos-de-artigos/introducao-profiler-java/icone_extensions-vscode.webp)


2. Na barra de pesquisa, digite **GraalVM Tools for Java** e faça a instalação deste plugin.

3. Ao instalar esse plugin, na barra direita do teu editor irá aparecer um ícone com as letras **Gr**. Clique nele.

4. Ao clicar no ícone do plugin, ele pedirá na parte de cima para fazer o Download e Instalação do GraalVM ou escolher uma existente. Selecione a segunda opção e escolha a pasta do GraalVM que descompactou na pasta de **Documentos** na seção anterior.

![add-graalvm](/img/conteudos-de-artigos/introducao-profiler-java/add_existing_graalvm.webp)

5. Agora, precisamos instalar o componente do VisualVM, que é a ferramenta de profiler que iremos utilizar, basta expandir a opção do GraalVM que você adicionou no plugin e clicar para instalar o componente do VisualVM como na imagem abaixo:

![add-visualvm-component](/img/conteudos-de-artigos/introducao-profiler-java/add-visualvm-plugin.webp)

6. Pronto, ao fim da instalação do componente, você deve ver uma seção chamada **VisualVM** na parte inferior da aba do plugin com algumas opções *Process*, *CPU sampler*, *JFR*, etc.

7. Ao clicar no botão de executar o VisualVM, ele deverá abrir o programa de profiling, que possui esta estética:

![visual-vm-gui](/img/conteudos-de-artigos/introducao-profiler-java/visualvm_gui.webp)

Pronto, agora que configuramos e instalamos todas as ferramentas necessárias, podemos seguir com os passos seguintes, a criação do programa em Java que utilizaremos para testar a ferramenta de profiling e o uso dela efetivamente!

## VisualVM

Nota-se que nos passos anteriores instalamos duas ferramentas, o GraalVM e o VisualVM. O GraalVM é um projeto que cria uma máquina virtual(JVM) que visa melhorar o desempenho dos programas em Java e outras linguagens. Não vou entrar em mais detalhes porque aqui utilizamos ele apenas para instalar o plugin do VSCode e o componente do VisualVM, mas vale a pena conhecer este projeto, que é bastante interessante: https://www.graalvm.org/

O VisualVM é a ferramenta de profiling que iremos utilizar, desenvolvida pela Oracle, com sua primeira release sendo lançada no ano de 2008. Ela é uma ferramenta open-source e grátis, distribuída através da licença GPL e era embutida dentro da própria JVM até a versão 8, agora segue como um projeto separado.

## Programa de Teste

Para testarmos como o VisualVM funciona, vamos criar um programa bastante simples, com apenas uma classe, que executa os seguintes passos:

1. Recebe o path de um arquivo como parâmetro.
   1. Caso o usuário não passe o arquivo como parâmetro, o programa irá enviar uma mensagem de erro.

2. Abre o arquivo que o usuário passou como argumento e o percorre linha por linha.
   1. Caso o arquivo não exista ou é inválido, ele irá lançar uma mensagem de erro especificando que o arquivo é inválido.

3. Para cada linha do arquivo, ele verifica, utilizando uma regex, se a linha corresponde a um formato de email válido ou não.

4. Caso a linha corresponda a um email com formato válido, iremos escrever no terminal que o email é válido, caso contrário, iremos escrever que o email é inválido.

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

Para rodar esse programa, basta criar um arquivo chamado **Main.java** em tua máquina local, copiar e colar o programa acima, baixar o arquivo com o dataset, ir à pasta onde criou o arquivo e rodar o seguinte comando no terminal:

```shell
$ java Main.java [caminho-para-o-dataset]/emails_400k.csv
```

Lembre-se que você deve utilizar o Java 17, ou o comando acima pode dar erro durante a execução.

## Profiling na prática

Agora que criamos a primeira versão do programa na versão anterior, vamos executar o VisualVM e executar o profiling, na prática.

### Instalando o Startup Profiler

Antes de executarmos o VisualVM no nosso código, precisamos adicionar o plugin de **Startup Profiler** dentro dele, que nos permite conectar a ferramenta de profiling com o programa desde o momento que ele começa a rodar. Para instalar o plugin, basta fazer os seguintes passos:

1. Abrir o VSCode.

2. Clicar no ícone do plugin do GraalVM, com as letras **Gr**, na lateral esquerda do editor.

3. Executar o VisualVM.

4. Ao abrir o VisualVM ir à parte superior em **Tools** -> **Plugins**.

5. Uma janela adicional irá se abrir. Basta ir em **Available Plugins** escolher o plugin **Startup Profiler** e instalá-lo. O processo pode demorar alguns minutos.

![startup-profiler](/img/conteudos-de-artigos/introducao-profiler-java/startup_profiler_jvm.webp)

### Executando o VisualVM

Agora que instalamos o Startup Profiler no VisualVM, podemos executá-lo.

1. Abra o VisualVM

2. Na parte superior, que contém vários ícones, clique no ícone que se assemelha a um pequeno relógio. Ao passar o mouse em cima dele estará escrito *Starts a new process and profile its startup*.

3. Uma nova janela irá se abrir, pedindo para adicionar as configurações que serão utilizadas. Adicione as seguintes configurações abaixo, exceto a seção 3, que é gerada automaticamente na tua máquina:

![profiler-config](/img/conteudos-de-artigos/introducao-profiler-java/profile_config.webp)

4. Clique no botão **Copy to clipboard** dentro da seção 3 da tela e salve esse texto em algum local, que iremos utilizar no futuro.

5. Clique no botão **Profile**

6. Caso seja sua primeira vez utilizando o profiler, ele irá pedir para você indicar qual o path do arquivo Java que está utilizando, basta adicioná-lo, como na tela abaixo(pode diferir na tua máquina caso tenha instalado o Java em outro local).

![calibrate-java](/img/conteudos-de-artigos/introducao-profiler-java/calibrate_java.webp)

7. Ao clicar em OK, ele irá efetuar o processo de identificar o programa e versão do Java que está utilizando por alguns segundos e irá direcioná-lo para uma tela onde ele espera o teu programa começar a rodar.

![connecting-to-jvm](/img/conteudos-de-artigos/introducao-profiler-java/connecting_jvm.webp)

8. Pronto, agora o VisualVM está esperando o teu programa começar a rodar para ele executar o profiling.

9. Abra o terminal e vá ao diretório onde você salvou o arquivo com o programa.

10. Execute o programa com o parâmetro adicional que você copiou no passo 4 desta seção, você irá adicioná-lo antes do nome do arquivo. Será algo similar ao comando abaixo:

```sh
$ java -agentpath:[CAMINHO-GRAALVM]/graalvm-ce-java17-22.3.1/lib/visualvm/visualvm/lib/deployed/jdk16/linux-amd64/libprofilerinterface.so=[CAMINHO-GRAALVM]/graalvm-ce-java17-22.3.1/lib/visualvm/visualvm/lib,5140 Main.java [CAMINHO-DATASET]/emails_400k.csv
```

11. Ao rodar o programa, você verá que ele começa a adicionar a saída no terminal e logo você é redirecionado para o VisualVM, que vai mostrando em tempo real as informações do programa:

![profiler](/img/conteudos-de-artigos/introducao-profiler-java/main_area_profiler.webp)

12. O programa leva em torno de 12-20 segundos para executar completamente, quando ele termina de executar, aparece uma caixa perguntando se você deseja criar um snapshot, basta selecionar que sim:

![application-finished](/img/conteudos-de-artigos/introducao-profiler-java/application_finished.webp)

Pronto, com isso, temos os primeiros resultados do profiling do nosso código.

### Monitor

![monitor](/img/conteudos-de-artigos/introducao-profiler-java/profile_monitor.webp)

A primeira tela que é importante mencionar, é a tela do **Monitor**, ela nos dá algumas informações gerais sobre o nosso programa.

**CPU**: Aqui nos mostra o quanto o programa consumiu de recursos de processamento enquanto esteve rodando, além de mostrar os momentos que o Garbage Collector(GC) está executando, o responsável por excluir os objetos que foram criados e não estão sendo utilizados, de modo a controlar a quantidade de memória utilizada na JVM. Este programa como é bastante simples e roda em um tempo bastante rápido, nota-se que não tem nenhum momento que o GC é executado.

**Heap**: A memória heap, é o local dentro da máquina virtual que roda o Java(JVM) onde os objetos criados são "guardados". O responsável por gerenciar o ciclo de vida desses objetos, excluindo aqueles não utilizados é o Coletor de Lixo(GC). Esta parte do Monitor mostra a quantidade de memória existente, além da quantidade de memória utilizada ao longo do programa.

**Threads**: As threads são processos dentro do programa que rodam paralelamente. Esta parte do Monitor mostra quantas threads existem dentro da máquina virtual do Java.

### Analisando o Snapshot

Agora que vimos a tela de Monitor, que mostra as informações gerais, vamos analisar a tela de Snapshot, que contém informações mais detalhadas sobre o nosso programa.

![snapshot_first](/img/conteudos-de-artigos/introducao-profiler-java/profiler_snapshot.webp)

Primeiramente, se notar o passo 3 da [sessão de executar o VisualVM](#executando-o-visualvm), selecionamos a opção **cpu**. Essa opção nos dá informações sobre o tempo de execução dos métodos dentro do código e a quantidade de vezes que eles foram chamados.

Na imagem do Snapshot, nota-se três colunas:

- **Total Time**: mostra durante o tempo de execução, quanto tempo o programa gastou executando aquele método. É importante citar que esse tempo não é o tempo de uma invocação, mas sim de todas as invocações somadas.

- **Total Time(CPU)**: similar ao Total Time, mas mostra quanto tempo foi efetivamente gasto com processamento. Isso acontece porque existem operações de I/O, isto é, de transferência de dados com serviços externos, como enviar uma requisição na WEB ou escrever em um arquivo, que não são contabilizadas nessa contagem.

- **Invocations**: A quantidade de vezes que o método foi executado.

A partir dessas colunas, podemos ver claramente que o gargalo do nosso código está no método **printEmailLine**, porque é onde nosso programa mais gastou tempo durante o processamento. Além disso, ele possui o mesmo número de invocações do método **isValidEmail**, mas consumiu 10x mais tempo que ele. Logo ele será nosso ponto de atenção para otimização.

Outra coisa interessante de se notar é que dos 10s que o método consumiu, apenas 1.2s foram utilizados em processamento, se visualizar a coluna Total Time(CPU) e o resto foi de operações de I/O. Isso acontece porque, se olhar o exemplo do código mostrado, basicamente ele está escrevendo uma mensagem no console/terminal, que é uma operação de I/O, o que nos mostra que temos que achar algo que seja mais performático que essa função **System.out.println** para escrever as mensagens no terminal.

## Otimizando os códigos

Pronto, agora que executamos a ferramenta de profiling, analisamos os resultados obtidos e identificamos o gargalo principal do código, chegou a hora de aplicar as otimizações e executar o profiler novamente, para ver se houve alguma mudança.

### Otimizando o método printEmailLine

Notamos duas coisas:

- O método **printEmailLine** é o maior gargalo na execução do programa.

- Ele utiliza apenas o **System.out.println**, então temos que achar uma alternativa mais performática que ele para escrever as mensagens no terminal.

Como o foco do artigo é o profiling, já entrego o spoiler do que vamos utilizar para otimizar este método, o **BufferedWriter**.

As operações de I/O geralmente tem um custo bem significativo, porque dependem de alguns fatores, como velocidade da internet, ou a velocidade de ler ou escrever no disco. Neste caso, quando chamamos o método do *println*, escrevemos uma linha por vez no terminal, que multiplicado pela quantidade de vezes que o método é chamado, causa esse gargalo que estamos vendo.

Com o **BufferedWriter**, nós vamos acumulando essas linhas em um "buffer" em memória e escrevemos várias delas de uma vez só quando o buffer atinge o limite ou caso a gente force isso acontecer, o que reduz bastante as operações de I/O no código. Tem um [artigo(em Inglês)](oracle.com/technical-resources/articles/javase/perftuning.html) da Oracle bastante interessante, que mostra várias estratégias de aumentar a performance do programa quando se lida com I/O.

A primeira modificação que vamos fazer é adicionar uma variável global dentro da classe Main, antes do método main:

```
public class Main {

    private static final BufferedWriter WRITER = new BufferedWriter(new OutputStreamWriter(System.out));

    // Mantém os outros métodos iguais.
}
```

Agora, basta modificarmos o printEmailLine para utilizarmos essa variável que criamos.

```
private static void printEmailLine(String line, boolean isValidEmail) throws IOException {
    if (isValidEmail) {
        WRITER.write("O email eh valido: " + line);
        WRITER.newLine();
    } else {
        WRITER.write("O email eh invalido: " + line);
        WRITER.newLine();
    }
}
```

É importante ressaltar que o BufferedWriter não adiciona a quebra de linha de forma automática, de forma semelhante ao println, por isso adicionamos o método **newLine**. É possível adicionar o caractere da quebra de linha dentro da string que vamos escrever, mas utilizar este método é melhor, pois a quebra de linha é representada de forma diferente em alguns sistemas operacionais, o que este método já trata.

O último passo é chamar o método **flush** do BufferedWriter ao fim da execução para garantir que todos os dados foram escritos no terminal e não resta nenhum no buffer.

```
try (var fileReader = getFileReader(args[0])) {
    String line;
    while ((line = fileReader.readLine()) != null) {
        boolean isValidEmail = isValidEmail(line);

        printEmailLine(line, isValidEmail);
    }
} catch (IOException e) {
    System.err.println("Erro enquanto inicializa o arquivo");
} finally {
    WRITER.flush();
    WRITER.close();
}
```

Agora com todas as modificações feitas, podemos realizar novamente os passos da seção de executar o VisualVM e executar o programa com as novas alterações para ver se elas impactaram o código de forma positiva:

![optimization-print-email-line](/img/conteudos-de-artigos/introducao-profiler-java/optimized_print_snapshot.webp)

UAU! Agora, nota-se que o tempo de execução do programa e do método caíram quase 10s. Além disso, agora o programa está consumindo mais tempo total na execução do método **isValidEmail** do que no método anterior, o que mostra que as modificações surtiram efeito.

### Mais otimizações

Depois de realizar a primeira otimização no código, vimos que ela surtiu efeito, diminuindo o tempo de execução do nosso método e consequentemente do programa.

Agora, o método que causa mais impacto na execução do programa é o **isValidEmail**, então vamos verificar uma forma de otimizá-lo também.

Ao analisar o método, nota-se que ele basicamente "compila" uma [regex](https://medium.com/cwi-software/e-o-regex-como-vai-657f94388dc) e aplica ela na string passada no método, para ver se ela satisfaz as condições impostas e é um email no formato válido.

Outro spoiler, para otimizar esse método, basta criarmos uma variável global com o resultado da chamada **Pattern.compile**, já que a regex não muda, e no método apenas chamarmos os outros dois métodos, o **matcher** e o **matches**.

A razão disso é que "compilar" regex tem um certo custo, ainda mais quando estamos executando esse passo várias vezes, como acontece nesse código. Dessa forma, como a regex sempre vai ser a mesma, podemos cortar economizar o processamento, compilando-a apenas uma vez.

Primeiro, adiciona-se uma nova variável de ambiente com o resultado da compilação da regex:

```
public class Main {

    private static final BufferedWriter WRITER = new BufferedWriter(new OutputStreamWriter(System.out));

    private static final Pattern PATTERN = 
        Pattern.compile("^(?=.{1,64}@)[A-Za-z0-9_-]+(\\.[A-Za-z0-9_-]+)*@[^-][A-Za-z0-9-]+(\\.[A-Za-z0-9-]+)*(\\.[A-Za-z]{2,})$");

    // Os outros métodos não mudam
}
```

Agora, basta modificarmos o método *isValidEmail* para tirarmos a chamada para o *Pattern.compile* e utilizar a variável que adicionamos:

```
private static boolean isValidEmail(String line) {
    return PATTERN
        .matcher(line)
        .matches();
}
```

Agora, basta executarmos o novo código com essas modificações adicionadas e ver no VisualVM o impacto delas no código:

![optimization-is-valid-email](/img/conteudos-de-artigos/introducao-profiler-java/optimization_isValidEmail.webp)

Pronto, tivemos uma queda de aproximadamente 50% do tempo de execução do método e o aumento da performance do programa. O método *printEmailLine* voltou a ser o método que mais consome tempo do programa, mas acredito que as modificações estão boas o suficiente por agora, conseguimos abaixar o tempo do programa em 10s, o que é ótimo.

A versão final do código otimizado, se encontra neste [gist](https://gist.github.com/anthonylgf/4b40b8920c678a80b307a75612c7717c).

## Analisando a Memória

Até agora, utilizamos a nossa ferramenta apenas para analisar a questão relacionada ao processamento. Mas é possível ver também como o nosso programa está utilizando a memória e alocando os objetos.

Para fazer isto, basta seguir os passos descritos da configuração do VisualVM, porém as opções de configuração serão diferentes, semelhantes à imagem abaixo:

![memory-profiler-setup](/img/conteudos-de-artigos/introducao-profiler-java/memory_profiler_setup.webp)

As diferenças são as seguintes:

- Nas configurações do profiler, seleciona-se **Memory** ao invés de CPU.

- Na caixa de **Profile classes**, nós colocamos um asterisco para indicar que queremos ver todos os objetos criados. Você pode limitar as classes dos objetos que quer ver, conforme a necessidade.

- A opção **Track only live objects** indica que apenas os objetos que ainda não foram excluídos pelo coletor de lixo serão mostrados.

Os próximos passos são os mesmos das etapas anteriores, para executar e conectar o VisualVM ao programa.

O que aparece agora no snapshot quando executamos o programa é o seguinte:

![memory-profiler-response](/img/conteudos-de-artigos/introducao-profiler-java/memory_profiler.webp)

Aqui, notamos que as informações mostradas são acerca da quantidade de objetos alocados, separados por cada classe que eles pertencem, e o quanto de memória eles ocupam na heap da máquina virtual do Java.

Essas informações são bastante importantes quando você está analisando o teu código procurando informações relacionadas ao consumo de memória do teu programa.

## Considerações finais

Neste artigo, tivemos uma pequena introdução acerca das ferramentas de profiling, desde a configuração até como utilizá-la para resolver problemas, utilizando o Java e o VisualVM como exemplo.

Obviamente a ferramenta possui bem mais recursos disponíveis, que auxiliam bastante, então vale a pena dar uma olhada na documentação oficial para ver tudo que ela pode oferecer.

Além disso, existem outras ferramentas de profiling pagas com mais recursos à disposição, mas o uso do VisualVM ajuda bastante na compreensão do mecanismo de um profiler e você pode aproveitar isso quando for utilizar outras ferramentas distintas.

Dessa forma, vale a pena entender como essas ferramentas funcionam, pois elas irão te ajudar bastante a melhorar a performance e qualidade do software que você cria.