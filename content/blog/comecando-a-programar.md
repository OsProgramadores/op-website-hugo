+++
date = "2020-10-26T00:00:45-00:00"
title = "Como começar a programar?"
categories = ["artigos"]
tags = ["artigos", "programação"]
banner = "img/banners/learn.jpg"
+++

🪄 Uma pergunta frequente é _"como começar a programar?"_.

Sei que muitos já escreveram sobre, mas como essa pergunta ainda aparece, vou respondê-la em primeira pessoa. Espero que você possa aproveitar essa experiência para se inspirar ✨

## Começando

A primeira vez que "escrevi algum código" foi ainda criança, usando ferramentas bem visuais: Delphi e RPG Maker. É fantástica a empolgação em rapidamente ter um feedback do que estou criando, resolver puzzles, e ter a sensação de poder mandar o computador fazer o que quero.

Como primeiro passo, creio que o mais importante é criar gosto. Linguagens visuais (como RPG Maker) ou educaionais (como Scratch e Logo), além de IDE gráficas (como Delphi) são boas primeiras opções para o primeiro contato. Assim já cria familiaridade de como deve-se comunicar com o computador.

A propósito, apesar de Logo ter sido feita voltada para ensinar crianças, ela é usado inclusive em graduação, como foi no caso da minha faculdade, o IFCE. Com Logo escrevi [esse artigo de iniciação científica](http://hp.unifor.br/encontros2014/PDFs/17883%20-%20Resumo.pdf) - veja só o quão Logo é maravilhoso!

Nesse primeiro momento, o instinto de querer criar "o novo aplicativo tipo o Uber" precisa ficar em segundo plano. **Primeiro faça algo simples e divertido usando uma linguagem com feedback rápido.** Por exemplo, desenvolva um Flappy Bird usando Logo. Ou um jogo como Final Fantasy I no RPG Maker. São primeiros passos factíveis, divertidos, e ao final terá um artefato finalizado. Isso é bem mais gratificante do que um projeto colossal inacabo.

Depois de ter experimentado essas linguagens visuais, quais seriam os próximos passos?

## Open Source

O que me fez migrar de projetos pessoais simples foi querer colaborar num projeto open source.

Sei que diversas pessoas aprenderam a programar de diferentes formas: conheço quem começou para fazer site de hentai, outro que foi para fazer warez... porém, o meu divisor de águas foi querer colaborar por 4fun no OpenKore - bot para o jogo Ragnarok, da qual gosto muito. Me sentir fazer parte da comunidade do OpenKore e ter feedback dos usuários a cada commit no SVN eram os meus motivadores para programar.

Assim, sempre quando me perguntam "como começar a programar?", eu sempre respondo: colabore num projeto open source que tenha uma comunidade para te apoiar, e que você se sinta motivado. Funcionou comigo e tenho esperança que dê para ti também.

No meu caso, eu já tinha um projeto em mente que queria colaborar, e você certamente já usa algum software open source que também gostaria de fazer parte. Mas se ainda não tiver ideia de algum para ajudar, falarei mais disso adiante.

Não se acanhe se a sua colaboração não for aceita de primeira, ou se for simples. Recaptulando o que disse antes: **ter artefatos finalizados é mais relevante do que algo colossal inacabado** - e isso também vale para contribuições em open source.

A minha primeira contribuição no OpenKore nem foi de código, mas sim de tradução. Depois fiz breves correções de bugs. O meu primeiro commit maior recebeu um feedback bem negativo de um colaborador experiente e, com base no feedback dele, aprimorei e corrigi os pontos levantados. Toda essa experiência foi gratificante! Haviam desenvolvedores experientes me ajudando de graça, eu tinha em mãos um complexo projeto para explorar, e diversos usuários fornecendo feedback no que eu criava. Assim conseguia errar e aprender num ambiente seguro e com rápido feedback loop... O famoso mantra _"Fail Fast. Learn Fast"_.

Ainda hoje aplico essa abordagem para aprender algo novo. Por exemplo: queria aprender C#, então decidi colaborar no Ryujinx, um emulador de Nintendo Switch escrito em C#. Abri alguns PRs no Ryujinx, e inclusive um deles foi citados no release notes! Outro exemplo: queria aprender mais sobre compiladores, então abri um PR no Babel corrigindo um simples bug, e também fui citado no release notes! A simples sensação do PR mergeado e ser citado no release notes é bem motivador, especialmente pelo sentimento de fazer parte.

## Projetos Open Source para colaborar

👥 Quais projetos Open Source seriam bom para começar?

Os projetos abaixo podem ser grandes, mas não se acanhe! Normalmente há issues com tag de "beginner friendly", já servindo de norte para onde começar.

### Battle for Wesnoth

<center>
  <img src="https://pbs.twimg.com/media/E0I-ABhXEAon67m?format=jpg&amp;name=small" alt="Imagem do Battle for Wesnoth" height="250px">
</center>
<br />

Battle for Wesnoth é um dos meus jogos open source favoritos. Joguei muito quando criança, e ainda hoje é bem mantido. O gameplay é de estratégia baseado em turnos, há diversas campanhas, algumas com história bem elaborada. Além disso, a IA é impecável e tem bom multiplayer.

O jogo é escrito principalmente em C++ com o SDL2. A parte da IA usa Lua, e alguns trechos estão em Python. Inclusive, ele tem uma linguagem de marcação própria, WML, para facilitar a escrita das campanhas.

- Site: https://www.wesnoth.org/
- GitHub: https://github.com/wesnoth/wesnoth

### Meus projetos

Fazendo uma auto-referência, eu tenho dois projetos que gosto de recomendar, pois estão bem documentados.

#### MacroCompiler

<center>
  <img src="https://user-images.githubusercontent.com/9501115/138568585-6b64824e-f0d0-45fb-8bb9-7763b8b7c9c7.png" alt="Talk que demonstro em uso o MacroCompiler" height="300px">
</center>
<br />

Esse projeto está relacionado ao Ragnarok e o bot para ele, o OpenKore.

Com o OpenKore é possível executar macros para automatizar uma sequências de ações que o bot fará. Já o MacroCompiler é um compilador de macros para Perl, para melhorar a performance e desenvolvimento das macros.

O projeto é quase todo escrito em Elixir, e algumas partes dos testes estão em Perl.

Tenho palestras que explicam em detalhes o MacroCompiler e sobre como compiladores funcionam. Ademais, também vale a leitura do readme.

- GitHub: https://github.com/macabeus/macro-compiler
- Talk falando de compiladores e do MacroCompiler: https://www.youtube.com/watch?v=t77ThZNCJGY

#### klo-gba.js

<center>
  <img src="https://camo.githubusercontent.com/4ef7552db91afd55a1f120034e1f4c60c38d7b9e1cb8b998955f05422896df7a/68747470733a2f2f692e696d6775722e636f6d2f565763394b336b2e706e67" alt="Talk que demonstro em uso o MacroCompiler" height="300px">
</center>
<br />

Caso queira aprender sobre engenharia reversa, ou React, ou WebGL, ou WebAssembly, esse projeto é o que você precisa.

O klo-gba.js é uma ferramenta web para customizar as fases do jogo de Gameboy Advance "Klonoa: Empire of Dreams". Ele usa WebGL para exibir a fase e permitir edição dela, também faz load da ROM e salva a versão customizada dela. Toda a interface web é feita com React.

O projeto é principalmente em JS, e algumas partes em C, que são compiladas para WebAssembly.

Tenho palestras e blog posts que explicam em detalhes todo o desenvolvimento do klo-gba.js.

- GitHub: https://github.com/macabeus/klo-gba.js
- Posts: https://macabeus.medium.com/pt-br-engenharia-reversa-num-jogo-de-gameboy-advance-introdu%C3%A7%C3%A3o-21ebffe2f794
- Talk (em inglês, e há um breve problema no áudio corrigido em ~2:00): https://www.youtube.com/watch?v=RMM_5bq3Ct8

### Awesome For Beginners

Se ainda quiser mais recomendações, tem essa awesome list com vários projetos beginner friendly:

https://github.com/MunGell/awesome-for-beginners

## Como criar foco?

<center>
  <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT-8sPjkCsBEtVo_Za11d_2aL31v1M6-XrGof-aVKW3KX1gw8bgN0BcsN8yK3U68j4ZFV4&usqp=CAU" alt="Meme ironizando ter uma série de projetos inacabados e mesmo assim começar mais um novo" height="300px">
</center>
<br />

🔍 Okay, agora temos um monte de coisa para programar... mas não consigo me focar! Como posso criar foco para terminar algo?

Essa é uma dúvida comum. Sempre faziam para mim quando eu palestrava sobre o klo-gba.js, da qual dediquei mais que um ano focado programando esse único projeto, sem pular de galho em galho.

Inclusive, há situações que dominar a habilidade de foco é mais relevante do que a de desenvolvimento. Programador mediano com foco consegue um elegante portfólio no GitHub por ter vários projetos concluídos, enquanto um programador excelente sem foco nunca termina algo para apresentar num portfólio.

Ao invés de pular de galho em galho entre projetos inacabados, é muito mais elegante terminar um projeto.

Segue breve explicações de quais são as minhas estratégias - e que você pode replicar!

### Skinner Box

Acima comentei que uma excelente forma para aprender a programar é colaborar num projeto open source.

Isso favorece o engajamento numa comunidade, o que é importante, pois você criou uma responsabilidade não apenas consigo mesmo, mas com outras pessoas. Fazer parte de uma comunidade gera sentimento de grupo. Não é à toa que jogos onlines incentivam a colaboração, para criar amizades que prendem o jogador. Porém, ainda falta algo para garantir mais foco, e um outro mecanismo importante e também presente em game design é o _"Skinner Box"_.

Resumo do conceito: pequenas recompensas contínuas garantem mais tempo aplicado numa atividade. A recompensa libera dopamina, e passaremos a querer novamente, assim faremos mais da atividade para obtê-la... num loop de reforço positivo.

<figure style="display: flex; flex-direction: column; align-items: center; font-size: 75%;">
  <img
    src="https://upload.wikimedia.org/wikipedia/commons/thumb/b/b4/Skinner_box_scheme_01.svg/300px-Skinner_box_scheme_01.svg.png"
    alt="Desenho esquemático do Skinner Box"
    height="300px"
    style="object-fit: contain;"
  >
  <figcaption>
    Desenho esquemático do skinner box
  </figcaption>
</figure>

<br />

E como isso se aplica na programação?

Mostra-se relevante criar gatilhos de recompensa no desenvolvimento de um projeto pessoal. Pode ser simples, como uma checklist. Marcar um item como finalizado na checklist será gostoso!

Um gatilho de recompensa que me foi muito útil é o feedback. Com cadência, lançava uma nova versão do klo-gba.js e a divulgava para a comunidade. Assim via as pessoas comentando, usando... o que é gratificante. Isso soma-se com o que falei antes que é útil: engajar em comunidade.

### Objetivos claros

Outro importante conceito em game design é oferecer objetivos claros. Num jogo pode ser simples, como coletar peças ou ir até o outro lado da fase. O que importa é ter um objetivo claro em mente do que fazer. Além disso, precisa ter um caminho claro para alançar o objetivo.

Por exemplo, no Mario o objetivo é claro: salvar a princesa. Mas até isso chegar, temos pequenas fases, e um mapa mostrando cada uma das fases que ainda devemos passar e onde estamos.

E como isso se aplica na programação?

Algo simples é ter uma checklist suficientemente completa. Não precisa ter todos os items, mas a quantia para se localizar e onde quer chegar.

Por exemplo, o klo-gba.js tem um objetivo complexo: customizar as fases de um jogo. Então precisamos quebrá-lo em objetivos menores. Um mais simples é inicialmente só exibir a fase e, em seguida, editar uma parte especifica de uma única fase. Depois podemos generalizar para incluir mais fases. E assim por diante.

### Deadline

Outro quesito relevante é criar deadline externo.

Por exemplo, tenho um amigo que paga inscrições para certificações, pois pagando caro numa prova o motiva a estudar para passar.

Já o deadline externo que eu costumo adotar é outro, gratuito e divertido: talkear em meetups e conferências.

Como é impossível alterar a data do evento, torna-se um deadline externo. Eu preciso ter algo foda para apresentar, então me foco em desenvolver o projeto até a data do evento. O objetivo final é a conf, e para ter uma cadência busco apresentar em meetups menores o projeto.

Além de ajudar a criar foco, outro benefício é que enquanto preparado a apresentação reviso diversos conteúdos que eu achava que sabia, e por vezes descubro que eu não sabia tanto assim. Além de que ao final da apresentação terei aprendido mais e criado um material de referência que eu mesmo poderei consultar no futuro para relembrar o conteúdo.

## Conclusão

E aqui encerro a postagem! Esperto ter te ajudado a se nortear 😋

Caso queria ler mais, abaixo está o link de minha postagem que explico em mais detalhes tanto o aspecto de como desenvolver um projeto pessoal, como também as recompensas que eles trazem para si: https://macabeus.medium.com/engenharia-reversa-num-jogo-de-gameboy-advance-e-a%C3%AD-o-que-obtemos-parte-final-9fc30311f22e

Me mande tweet com os projetos que decidiu começar a colaborar, ou qual projeto pessoal decidiu começar a desenvolver com foco.

Minhas redes sociais:

- Twitter: https://twitter.com/bmacabeus
- GitHub: https://github.com/macabeus
