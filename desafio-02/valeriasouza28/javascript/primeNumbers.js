// Código para verificar se um número é primo ou não

const final = 10000
function generate() {
  for (let i = 0; i <= final; i++) {
    if (i === 1) {
      console.log(i + ` Não é um número primo!`)
    } else if (i === 2) {
      console.log(i + ` É um número primo!`)
    } else if (i % 2 === 0) {
      console.log(i + ` Não é um número primo!`)
    } else {
      console.log(i + ` É um número primo!`)
    }
  }
}

generate()
