const fs = require('fs')
const a = fs.readFileSync('./Invictus.txt')
let result = Buffer.from('')
const base = 4087382455 - 8
let resultString = ''
for (let i = 0; i < a.length; i++) {
  const b = a[i]
  if (b === 243) {
    const intValue1 = parseInt(a.slice(i, i +4).toString('hex'), 16)- base + 65
    const intValue = intValue1 > 260 ? intValue1 - 278 + 86 : intValue1
    console.log(intValue, String.fromCharCode(intValue))
    resultString += String.fromCharCode(intValue)
    result = Buffer.concat([result, a.slice(i, i + 4)])

  }
}
console.log(result, result.length / 4)
fs.writeFileSync('./filtered.txt', result)
fs.writeFileSync('./filtered-string.txt', resultString)
// od -c -b Invictus.txt