+++
date = "2017-08-24T23:34:06-03:00"
title = "Ordenação bubble sort"
categories = ["ordenação","algorítimo"]
tags = ["bubblesort"]
banner = "img/banners/bubbles.jpg"

+++

## Algorítimo de ordenação: _bubble sort_[^fa]

Os dados de um programa podem ser armazenados na memória, porém para que passem a ser considerados informação e ter sentido, muitas vezes é necessário ordená-los. Existem várias formas de ordenar os dados, algumas são mais eficientes do que outras, porém as mais eficientes também podem ser as mais complexas em implementar.

Desse forma, neste artigo vamos aprender um dos mecanismos de ordenação mais simples para quem está estudando estruturas de dados. O _bubble sort_ é um dos algorítimos menos eficientes, contudo mais fáceis de implementar. O termo _bubble_ advém da forma como o vetor é ordenado, como bolhas[^f1] que vão buscando seu lugar emergindo na ordenação desejada do vetor.

### Demonstração

Considere um algorítimo de vetor em _c++_ com "_n_" elementos:

```c++
#include <iostream>

using namespace std;

int main() {

  int vet[]={3,0,1,8,7,2,5,4,6,9};

  int len = sizeof(vet)/sizeof(*vet);
  // Tamanho 4 bytes para cada elemento inteiro no array = 40 bytes / 4 bytes do tamanho do deferenciamento do endereço do vetor = 10 unidades.

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

Explicando o algorítimo acima, podemos ver que o vetor _vet_ possui 10 (dez) elementos desordenados. A função _main_ do algorítimo _bubble sort_, ordenadamente varre o vetor e estabecele comparações sucessivas entre pares de elementos.

Veja que o _loop_ _for_ externo faz a iteração através do vetor para cada iteração do _loop_ interno, aplicando a comparação entre o elemento inicial e seu sucessor no vetor _vet_. Se a condição _if_ é atingida, isto é, o sucessor é maior que o predecessor, eles trocam de posição com a ajuda da variável _aux_. Dependendo do tipo de ordenação, do maior para o menor ou o inverso, podemos mudar a condição _if_.

## Executando o algorítimo _bubble sort_

Ao executar, temos a seguinte saída para o Vetor Desordenado _vet_:

```sh
3 0 1 8 7 2 5 4 6 9               

Quantas vezes no "for": 90
Quantas  vezes no "if": 32        
9 8 7 6 5 4 3 2 1 0 

Quantas vezes no "for": 90        
Quantos vezes no "if": 45         
0 1 2 3 4 5 6 7 8 9 

```
A Complexidade é (O) = N², isto é, tivemos 10 iterações do loop externo  e 9 do loop interno. Tivemos 32 trocas para ordenador do vetor desordenado para ordem decrescente, enquanto tivemos 45 trocas para ordenador do vetor desordenado para ordem crescente.

De maneira mais lúdica, a ordenação será uma dança das cadeiras entre pares de números. 
Veja uma demonstração neste vídeo:

{{< youtube  lyZQPjUT5B4 >}}

## Considerações finais

A ordenação de dados é uma das intervenções mais frequentes em qualquer conjunto de dados. Dominar os algorítimos é essencial para tornar-se um programador de sucesso. Dependendo do volume de dados, há outros algorítimos de ordenação mais apropriados.

Veja que os números inteiros do vetor _vet_ correspondem aos números nos trajes dos dançarinos. ;)


[^fa]: David Gesrob: [E-mail david81br@gmail.com](mailto: david81br@gmail.com);
[^f1]: Imagem do Banner: http://all-free-download.com/free-vector/download/bubble-bubble-vector-ornate-halos_296230.html;
