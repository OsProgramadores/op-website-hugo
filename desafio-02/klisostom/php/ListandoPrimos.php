<?php

namespace Desafio\klisostom\php;

class ListandoPrimos
{
    public function lista(): array
    {
        $foundPrimes = [2];

        $index = 3;
        while (count($foundPrimes) < 10000) {
            if ($this->ehImpar($index)) {
                $ehPrimo = true;

                for ($z = 2; $z < $index; $z++) {
                    if ($index % $z === 0) {
                        $ehPrimo = false;
                        break;
                    }
                }

                if ($ehPrimo) {
                    $foundPrimes = [...$foundPrimes, $index];
                }
            }

            $index++;
        }

        return $foundPrimes;
    }

    private function ehImpar(int $number): bool
    {
        return $number % 2 !== 0;
    }
}
