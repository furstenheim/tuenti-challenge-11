const fs = require('fs')
const content = fs.readFileSync(0).toString()
const input = content.split('\n')
const size = parseInt(input[0])
const _ = require('lodash')
console.error(size)
if (input.length < size + 1) {
  throw new Error(`Expected size ${size} got ${input.length}`)
}


const testCases = input.slice(1, size + 1)

const result = testCases.map(function (d, i) {
  if (!d.match(/^[-0-9A-Z */=+]+$/)) {
    throw new Error(`Unknown expression ${d}`)
  }
  console.error(i, d)
  const functionCode = parseExpression(d)
  eval(functionCode)
  const solutions = findSolutions()
  solutions.sort()
  return solutions.join(';')
}).map((d, i) => {
  if (!d) {
    return `Case #${i + 1}: IMPOSSIBLE`
  }
  return `Case #${i + 1}: ${d}`
}).join('\n')

console.info(result)

// SEND + MORE = MONEY
function parseExpression (exp) {
  const words = exp.match(/([A-Z])[A-Z]*/g)
  const initialLetters = _.uniq(words.map(w => w[0]))
  const initialAsMap = _.mapKeys(initialLetters)
  const restLetters = _.filter(_.uniq(exp.match(/([A-Z])/g)), l => !initialAsMap[l])

  const expressionAsArray = exp.replace('=', '==').split('')
  const functionCodeStart = [
    `function findSolutions () {
       const solutions = []
       const exp = ${JSON.stringify(expressionAsArray)}
    `
  ]
  const functionCodeEnd = [
    `}`,
    `return solutions`
  ]

  const allLetters = initialLetters.concat(restLetters)
  for (let i = 0; i < allLetters.length; i++){
    const letter = allLetters[i]
    const start = i < initialLetters.length ? 1 : 0
    functionCodeStart.push(`for (let e${i} = ${start}; e${i} < 10; e${i}++) {`)
    for (let j = 0; j < i; j++) {
      functionCodeStart.push(`if (e${i} === e${j}) continue`)
    }
    for (let j = 0; j < expressionAsArray.length; j++) {
      if (expressionAsArray[j] === letter) {
        functionCodeStart.push(`exp[${j}] = e${i}`)
      }
    }
    functionCodeEnd.push(`}`)
  }

  functionCodeStart.push(`
    if (eval(exp.join(''))) solutions.push(exp.join('').replace('==', '='))
  `)
  const functionCode = functionCodeStart.join('\n') + functionCodeEnd.reverse().join('\n')
  return functionCode // console.error(exp, initialLetters, restLetters, functionCode)
}
