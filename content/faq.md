+++
title = "Perguntas Frequentes (FAQ)"
date = "2018-08-20T23:00:00-07:00"
markup = "markdown"
+++

# Comportamento

## Qual é o foco do grupo?

O grupo foi criado por um grupo de amigos, veteranos em diversas áreas de computação e eletrônica, com o intuito de ajudar aqueles que estão começando a programar. A participação é completamente gratuita, mas pedimos que os novos integrantes leiam as [Regras](https://osprogramadores.com/regras) para evitar problemas.

## Por onde eu começo?

Temos alguns ótimos artigos introdutórios aqui no site! Recomendamos que você comece por [este artigo](https://osprogramadores.com/blog/2024/02/29/conceitos-essenciais-para-come%C3%A7ar/), leia-o com calma e atenção. Após concluí-lo, continue seus estudos através [deste artigo](https://osprogramadores.com/blog/2019/03/12/por-onde-comecar/).

## Posso fazer perguntas sobre o X?

Sim, desde que X seja ligado a programação e a computadores em geral. A conversa no grupo é variada, mas é importante não desviar o assunto do grupo para temas não relacionados. Tópicos particularmente não recomendados incluem política, sexualidade e qualquer outro tópico polêmico que não esteja relacionado à área. Tópicos envolvendo ilegalidade de qualquer forma sujeitam o autor a banimento imediato e irreversível.

## Posso perguntar se posso perguntar?

[Temos um artigo sobre este tópico](https://osprogramadores.com/blog/2020/09/11/como-perguntar/), mas de modo geral, perguntar se alguém sabe de A ou B para então perguntar o que precisa saber sobre A ou B não é uma forma eficiente de se comunicar... Recomendamos a leitura do artigo.

## Posso colocar anúncios de emprego no grupo?

Anúncios de empregos de qualquer forma são expressamente proibidos no grupo, sujeitando o usuário a banimento.

## Posso enviar fotos da minha tela mostrando o meu problema?

Em 90% dos casos, a resposta é **não**. A ordem de preferência é:

1. **Dúvidas com código**: Coloque o seu código (se possível) no [repl.it](http://repl.it) e coloque o link no canal. Procure isolar a parte do código com problemas (ou seja, não coloque um módulo de 1000 linhas, apenas a parte com problemas.)

1. **Problemas de formatação na tela**: Tire um print-screen da tela e envie o printscreen.

O envio de fotos de telas com códigos, em particular sem contexto suficiente, trará a ira dos admins e participantes. Considere-se avisado.

## Posso anunciar o meu grupo Telegram/Facebook?

A promoção de recursos ligados a área de programação, **sem fins lucrativos** é geralmente bem-vinda. Aceitamos a disseminação de grupos do Telegram, canais no Whatsapp, grupos no Facebook, blogs, podcasts, canais no Youtube, distribuição (legal) de livros gratuitos, etc.

Ao promover algo, mantenha em mente que **o grupo valoriza a participação acima de tudo**. Uma ocorrência frequente (e infeliz) são aqueles que entram no grupo com a única função de promover o seu grupo ou canal.  Os admins reservam o direito de apagar os posts que considerem "auto promoção sem participação".

## Estou sendo perseguido por outro membro. O que devo fazer?

Se você acredita estar sendo perseguido por outro participante, ou observou qualquer comportamento indiscreto ou inconsistente com um ambiente saudável, entre em contato com um dos administradores em private e explique o problema. Os administradores do grupo estão listados no final da página principal do site [OsProgramadores](https://osprogramadores.com).

# Podcast

## Onde posso assistir/ouvir o Podcast?

No momento há algumas opções: 

* [Na sessão de Podcasts em nosso site](https://osprogramadores.com/podcast/)
* [No Youtube](https://www.youtube.com/@OsProgramadores?sub_confirmation=1)
  * onde você pode ver os episódios gravados no formato **videocast** ([episódio 80 adiante](https://www.youtube.com/watch?v=4nBKN0eUToI&sub_confirmation=1)); 
  * como também pode ouvir o áudio de todos os nossos episódios.
* [No Spotify](https://open.spotify.com/show/0IrqGbURcNnumdHVkfKIFA)
  * onde você pode ouvir o áudio de todos os episódios do nosso podcast.
* [No Tiktok](https://www.tiktok.com/@osprogramadores)
  * onde recentemente passamos a publicar cortes dos episódios.

## Como faço para participar do Podcast?

Entrando em contato conosco através do [nosso grupo no Telegram](https://t.me/osprogramadores).

## Posso indicar alguém para ser entrevistado?

Pode sim! Recomendamos que você entre em contato conosco através do [nosso grupo no Telegram](https://t.me/osprogramadores) para saber como prosseguir. 

# Perguntas técnicas

## Qual a melhor linguagem para quem está começando?

Esse é um tópico tão frequente que os admins [escreveram uma entrada no blog](https://osprogramadores.com/blog/2017/04/07/qual_linguagem_usar/) sobre o assunto. Recomendamos a leitura. Sinta-se a vontade para discordar ou comentar no grupo.

## Qual o melhor editor/IDE?

A sua escolha de editor e/ou IDE é estritamente pessoal. Existe uma imensa variedade de editores e IDEs gratuitos para quase todas as plataformas. Nossa sugestão é testar todos e escolher o que achar melhor.

## Como indento o meu programa?

Programas mal-indentados são difíceis e cansativos de ler, e escondem bugs. Por favor, use um estilo de indentação consistente no seu programa.

**Dica #1**: Se o seu programa estiver no repl.it, o bot do grupo pode fazer indentação automaticamente através do comando `/indent <url_do_replit>`. O comando deve ser enviado *diretamente* ao bot (em private), que vai responder com uma URL atualizada contendo o seu programa reformatado.

**Dica #2**: Também é possível reformatar o seu programa automaticamente usando um site como o [Code Beautify](https://codebeautify.org/c-formatter-beautifier).

## Não consigo ler dados do teclado em C.

Infelizmente, a maioria dos professores pedem exercícios em C usando `scanf()` e `gets()`. Essas funções necessitam de cuidados adicionais que não são muito óbvios para iniciantes (e lamentavelmente, para muitos professores).

Se o seu programa precisa ler dados do teclado, recomendamos o uso das funções `get_string()` e `get_float()` em https://repl.it/@marcopaganini/Simple-Data-Entry. Sinta-se a vontade para copiar o código ou (melhor ainda) entender como o código funciona.

## Meu programa em C só roda em Windows...

Os administradores acreditam que a linguagem C é fortemente ligada ao ambiente Unix/Linux e não tem tempo livre ou motivação para ajudar com problemas específicos do ambiente Windows. Por favor, coloque o seu código no repl.it e modifique-o até que ele compile sem erros.

Alguns problemas comuns em programas C feitos para Windows:

* `system("clear")` ou `system("cls")`: Totalmente desnecessários e tornam o seu programa imediatamente incompatível com outros ambientes. Por favor, remova. Se for absolutamente necessário limpar a tela, use uma library específica (curses/ncurses ou similar).

* Header files (includes) inexistentes: O Windows possui um número de header files que não existem nas implementações usuais do C. Por favor, modifique o seu programa até que ele compile sem erros no repl.it.

* Chamadas para `set_locale()` e similares: Por alguma razão, alguns acreditam que acentos só serão mostrados na tela se o "locale" do computador for configurado para "pt_BR". Isto é incorreto e desnecessário.

## Eu não tenho um computador rodando Linux! Como faço?

Existem várias alternativas a ter um desktop Linux para testar o seu programa. Algumas:

* http://repl.it - permite o uso imediato de várias linguagens de programação, com ou sem uma conta. Altamente recomendado pelo grupo.

* http://codenvy.io - Outro Cloud IDE completo e gratuito.

# Ainda tem dúvidas?

Sinta-se a vontade para contactar um dos administradores listados na nossa home page, [osprogramadores.com](http://osprogramadores.com).
