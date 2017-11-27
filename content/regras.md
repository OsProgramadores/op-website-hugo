# Introdução

[OsProgramadores](http://osprogramadores.com) é uma entidade formada por um grupo de profissionais experientes na área de computação com o objetivo de promover a discussão de temas relacionados ao desenvolvimento de software (algoritmos, ferramentas, linguagens, etc).

A maior parte das discussões acontece no nosso grupo no Telegram ([#OsProgramadores](https://t.me/OsProgramadores)). O objetivo do grupo é oferecer um local informal e descontraído para a troca de informações, dúvidas, e conversas em geral sobre computadores e tópicos relacionados.

Este documento contém um conjunto de regras básicas para o uso do grupo. Os administradores **fortemente recomendam a leitura deste documento para todos os participantes do grupo**.

# Regras básicas

Tópicos e comportamentos **expressamente proibidos** no grupo (resultam em banimento imediato na maioria dos casos):

* Spam de qualquer forma (comercial, religioso, político, etc).
* Pregação religiosa.
* Pregação política.
* **Qualquer atividade ilegal**, como promoção de pirataria de qualquer forma, estelionato, "hacking", "carding", troca de dados obtidos de forma ilegal, venda ou consumo de drogas, etc (independente da sua opinião sobre o assunto).
* Ataques pessoais ou perseguição de qualquer forma.

Comportamentos não recomendados:

* Anúncios de ofertas de emprego. Ao invés de postar um anúncio inteiro, sugerimos uma mensagem discreta para o grupo na linha de "Tenho empregos imediatos pra posição X, maiores detalhes contate-me em privado".
* Envio de muitas mensagens com apenas uma linha: Tente separar as mensagens em unidades que façam sentido. Uma linha por mensagem torna a leitura cansativa.
* Envio excessivo de mensagens.
* Envio de mensagens sem sentido (como sequências longas de stickers ou memes que não se encaixam com o tópico sendo discutido)
* Discussões acaloradas fora do tópico do grupo, especialmente aquelas envolvendo religião, política e outros assuntos polêmicos.
* Envio de clips de áudio (clips serão sumariamente apagados.)
* "Drive by Shooting" (apenas aparecer no grupo para promover o seu produto ou evento, sem contribuir com mais nada).
* Paste de trechos longos de código (veja mais detalhes em "Como obter ajuda" abaixo)
* Discussões sobre "pentesting" e tópicos similares (por favor, procure um grupo sobre segurança da informação para discutir esses tópicos).

# Como obter ajuda

## Mantenha a sua pergunta dentro do tópico do grupo

O grupo foi feito para ajudar a todos com problemas relacionados a programação, assim como discutir linguagens, algoritmos, etc. Mantenha a sua pergunta dentro do tópico do grupo, sempre que possível.

## O grupo não vai fazer o seu dever de casa

Perguntas como "Meu trabalho da faculdade é fazer X" são bem vindas, **desde que o trabalho já tenha sido iniciado**. Em outras palavras, é OK perguntar porque um programa ou pedaço de código não funciona, mas não é OK perguntar como fazer sem ter tentado _nada_ antes.

## Demonstre interesse

É difícil ajudar aqueles que não tem qualquer interesse em resolver os próprios problemas. Infelizmente, várias pessoas esperam que uma solução apareça, sem fazer qualquer tipo de esforço, __mesmo quando alguém recomenda um caminho a ser tomado__. Lembre-se, o grupo está aqui pra ajudar, não pra resolver os seus problemas.

## Faça-se disponível para aqueles tentando ajudar

Uma vez feita a sua pergunta, esclarecimentos serão necessários. Responda as dúvidas daqueles tentando ajudar. Por favor, não abandone o tópico ou saia do grupo sem explicação.

## Evite fazer a mesma pergunta repetidamente

Todos os membros do grupo são voluntários. Se a sua pergunta não despertou interesse imediato, espere algumas horas antes de perguntar novamente. Evite fazer a mesma pergunta repetidamente.

## Não envie screenshots ou pastes do seu código

Screenshots mostrando o seu código não são a melhor forma de obter ajuda. Coloque o código (ou pedaço de código) no http://repl.it e então coloque o link para o grupo, com uma descrição sumária do problema. Screenshots de *sintomas* são aceitáveis (por exemplo, mostrando um problema de formatação criado pelo seu programa).

Se a sua linguagem não é uma linguagem aceita pelo repl.it, recomendamos o http://termbin.com ou um gist no github (http://gist.github.com).

## Não envie pedidos via áudio

Áudio clips não são tolerados no grupo por qualquer razão e serão imediatamente removidos.

### Não chame outros membros no privado para dúvidas comuns

Salvo casos especiais, discuta as suas dúvidas em público, no grupo. Isso permite que mais de uma pessoa o ajude.

## Problemas frequentes com a linguagem C

## Indente o seu programa corretamente

Programas mal-indentados são difíceis e cansativos de ler, e escondem bugs mais facilmente. Por favor, use um estilo de indentação consistente no seu programa.

**Dica #1**: Se o seu programa estiver no repl.it, o bot do grupo pode fazer indentação automaticamente através do comando `/indent <url_do_replit>`. O comando deve ser enviado *diretamente* ao bot (em private), que vai responder com uma URL atualizada contendo o seu programa reformatado.

**Dica #2**: Também é possível reformatar o seu programa automaticamente usando um site como o [Code Beautify](https://codebeautify.org/c-formatter-beautifier).

## Algo estranho acontece com os dados que eu digito (via scanf/gets/fgets)

Infelizmente, a maioria dos professores pedem exercícios em C usando `scanf()` e `gets()`. Essas funções necessitam de cuidados adicionais que não são muito óbvios para iniciantes (e lamentavelmente, para muitos professores).

Se o seu programa precisa ler dados do teclado, recomendamos o uso das funções `get_string()` e `get_float()` em https://repl.it/@marcopaganini/Simple-Data-Entry. Sinta-se a vontade para copiar o código ou (melhor ainda) entender como o código funciona.

## Meu programa em C só roda em Windows...

Os administradores acreditam que a linguagem C é fortemente ligada ao ambiente Unix/Linux e não tem tempo livre ou motivação para ajudar com problemas específicos do ambiente Windows. Por favor, coloque o seu código no repl.it e modifique-o até que ele compile sem erros.

Alguns problemas comuns em programas C feitos para Windows:

* `system("clear")` ou `system("cls")`: Totalmente desnecessários e tornam o seu programa imediatamente incompatível com outros ambientes. Por favor, remova. Se for absolutamente necessário limpar a tela, use uma library específica (curses/ncurses ou similar).

* Header files (includes) inexistentes: O Windows possui um número de header files que não existem nas implementações usuais do C. Por favor, modifique o seu programa até que ele compile sem erros no repl.it.

* Chamadas para `set_locale()` e similares: Por alguma razão, alguns acreditam que acentos só serão mostrados na tela se o "locale" do computador for configurado para pt_BR. Isto é incorreto e desnecessário.

## Eu não tenho um computador rodando Linux!

Existem várias alternativas a ter um desktop Linux para testar o seu programa. Algumas:

* http://repl.it - permite o uso imediato de várias linguagens de programação, com ou sem uma conta. Altamente recomendado pelo grupo.

* http://c9.io - Ambiente completo de programação gratuito.

* http://codenvy.io - Outro Cloud IDE completo e gratuito.

# Ainda tem dúvidas?

Sinta-se a vontade para contactar um dos administradores listados na nossa home page, [osprogramadores.com](http://osprogramadores.com).
