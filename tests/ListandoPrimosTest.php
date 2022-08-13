<?php

use Desafio\klisostom\php\ListandoPrimos;
use Desafio\klisostom\php\PrimosAteDezMil;

test('Listando números primos.', function () {
    // Arrange
    $listandoPrimos = new ListandoPrimos;
    $primosAteDezMil = new PrimosAteDezMil;

    // Act
    $expected = $primosAteDezMil->primes();

    // Assert
    expect($listandoPrimos->lista())->toBeArray();
    expect($listandoPrimos->lista())->toEqualCanonicalizing($expected);
});
