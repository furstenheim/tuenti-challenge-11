const fs = require('fs')
const content = fs.readFileSync(0).toString()
const input = content.split('\n')
const size = parseInt(input[0])
const moment = require('moment')
console.error(size)
if (input.length < size + 1) {
  throw new Error(`Expected size ${size} got ${input.length}`)
}
const languageMap = {
  CA: 'CA',
  CZ: 'CS',
  DE: 'DE',
  DK: 'DA',
  EN: 'EN',
  ES: 'ES',
  FI: 'FI',
  FR: 'FR',
  IS: 'is',
  GR: 'EL',
  HU: 'HU',
  IT: 'IT',
  NL: 'NL',
  VI: 'VI',
  PL: 'PL',
  RO: 'RO',
  RU: 'RU',
  SE: 'SV',
  SI: 'SL',
  SK: 'SK',
}

const testCases = input.slice(1, size + 1)
const result = testCases.map(t => t.split(':')).map(function ([date, language]) {
  if (date.match(/^[0-9]{4}\-[0-9]{2}\-[0-9]{2}$/)) {
    return computeDay(date, 'YYYY-MM-DD', language)
  } else if (date.match(/^[0-9]{2}\-[0-9]{2}\-[0-9]{4}$/)) {
    return computeDay(date, 'DD-MM-YYYY', language)
  } else {
    throw new Error(`Unknown date '${date}'`)
  }
}).map((r, i) => { return `Case #${i + 1}: ${r}`})
console.log(result.join('\n'))
console.error(result.join('\n'))

// TODO iso days
function computeDay (date, format, language) {
  if (!languageMap[language]) {
    return 'INVALID_LANGUAGE'
  }
  const parsedDate = moment(date, format)
  if (!parsedDate.isValid()) {
    return 'INVALID_DATE'
  }
  return parsedDate.locale(languageMap[language]).format('dddd').replace('ț', 'ţ').toLowerCase()
}