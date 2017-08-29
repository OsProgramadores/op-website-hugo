+++
date = "2017-08-24T23:34:06-03:00"
title = "Ordenação bubble sort"
categories = ["ordenação","algoritmo"]
tags = ["bubblesort"]
banner = "img/banners/bubbles.jpg"

+++

## Algoritmo de ordenação: _bubble sort_[^fa]

Os dados de um programa podem ser armazenados na memória, porém para que passem a ser considerados informação e ter sentido, muitas vezes é necessário ordená-los. Existem várias formas de ordenar os dados, algumas são mais eficientes do que outras, porém as mais eficientes também podem ser as mais complexas de implementar.

Neste artigo vamos aprender um dos mecanismos de ordenação mais simples para quem está estudando estruturas de dados. O _bubble sort_ é um dos algoritmos menos eficientes, contudo mais fáceis de implementar. O termo _bubble_ advém da forma como o vetor é ordenado, como bolhas[^f1] que vão buscando seu lugar emergindo na ordenação desejada do vetor.

### Demonstração

Considere um algoritmo de vetor em _c++_ com "_n_" elementos:

```c++
#include <iostream>

using namespace std;

int main() {

  int vet[]={3,0,1,8,7,2,5,4,6,9};

  int len = sizeof(vet)/sizeof(*vet);
  
  // O número de unidades do vertor foi obtido pelo tamanho 4 bytes para cada elemento inteiro no array = 40 bytes / 4 bytes
  // ref. ao tamanho do dereferenciamento do endereço do vetor = 10 unidades. Assim sabemos o limite "len" do vetor "vet".

  int j, i, n, aux, count1=0, count2=0;

  for(n = 0; n < len; n++)
    cout << vet[n] << " ";

  cout << "\n";

  // Do maior para o menor

  for(i = 0; i < len; i++)
    for(j = 0; j < len-1; j++) {

	  count2+=1;

      if(vet[j] < vet[j + 1]) {

		count1+=1;

        aux = vet[j];

        vet[j] = vet[j + 1];

        vet[j + 1] = aux;
      }
    }
	cout << "\nQuantas vezes no \"for\": " << count2 << endl;
	cout << "Quantas vezes no \"if\": " << count1 << endl;

  for(n = 0; n < len; n++)
    cout << vet[n] << " ";

  cout << "\n";

  count1=0;
  count2=0;

  // Do menor para o maior
  
  for(i = 0; i < len; i++)
    for(j = 0; j < len-1; j++) {

      count2+=1;

      if(vet[j + 1] < vet[j]) { // Aqui troca a avaliação lógica.

	count1+=1;

        aux = vet[j];

        vet[j] = vet[j + 1];

        vet[j + 1] = aux;
      }

    }
	cout << "\nQuantas vezes no \"for\": " << count2 << endl;
	cout << "Quantas vezes no \"if\": " << count1 << endl;


  for(n = 0; n < len; n++)
	cout << vet[n] << " ";

  cout << "\n";

  return 0;
}
```

Explicando o algoritmo acima, podemos ver que o vetor _vet_ possui 10 (dez) elementos desordenados. A função _main_ do algoritmo _bubble sort_, ordenadamente varre o vetor e estabecele comparações sucessivas entre pares de elementos.

Veja que o _loop_ _for_ externo faz a iteração através do vetor para cada iteração do _loop_ interno, aplicando a comparação entre o elemento inicial e seu sucessor no vetor _vet_. Se a condição _if_ é atingida, isto é, o sucessor é maior que o predecessor, eles trocam de posição com a ajuda da variável _aux_. Dependendo do tipo de ordenação, do maior para o menor ou o inverso, podemos mudar a condição _if_.


## Executando o algoritmo _bubble sort_

Ao executar, temos a seguinte saída para o vetor desordenado _vet_:

```sh
3 0 1 8 7 2 5 4 6 9

Quantas vezes no "for": 90
Quantas  vezes no "if": 32
9 8 7 6 5 4 3 2 1 0 

Quantas vezes no "for": 90
Quantos vezes no "if": 45
0 1 2 3 4 5 6 7 8 9 

```

A complexidade é (O) = N², isto é, tivemos 10 iterações do _loop_ externo  e 9 do _loop_ interno. Tivemos 32 trocas para ordenador do vetor desordenado para ordem decrescente, enquanto tivemos 45 trocas para ordenador do vetor desordenado para ordem crescente.

De maneira mais lúdica, a ordenação será uma dança das cadeiras entre pares de números. 
Veja uma demonstração neste vídeo:

{{< youtube  lyZQPjUT5B4 >}}

Veja que os números inteiros do vetor _vet_ correspondem aos números nos trajes dos dançarinos. ;)

## Considerações finais

A ordenação de dados é uma das operações rotineiras em qualquer conjunto de dados. Dominar os algoritmos é essencial para tornar-se um programador de sucesso. Dependendo do volume de dados, há outros algoritmos de ordenação mais apropriados e è medida do crescimento do volume de dados outras formas de ordenação serão mais convenientes.

[^fa]: David Gesrob: [E-mail david81br@gmail.com](mailto: david81br@gmail.com);
[^f1]: Imagem do Banner: http://all-free-download.com/free-vector/download/bubble-bubble-vector-ornate-halos_296230.html;
