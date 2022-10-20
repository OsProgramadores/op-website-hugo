def n_prime(num):
    for value in range(2,num):
        if num % value == 0:
            return False
    return True

def print_upto():
  up_to=10000
  for value in range(2,up_to+1):
    if(n_prime(value)):
      print(value)
while True:
  print_upto() 
  break