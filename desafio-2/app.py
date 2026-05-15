for num in range(1, 10001):

    quantidade_divisores = 0

    for num_da_divisao in range(1, num + 1):

        if num % num_da_divisao == 0:
            quantidade_divisores = quantidade_divisores + 1

    if quantidade_divisores == 2:
        print(num)
    
        