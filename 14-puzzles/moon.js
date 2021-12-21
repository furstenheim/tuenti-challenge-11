const solutions = findSolutions()
console.log(solutions)
function findSolutions () {
  const solutions = []
  const exp = ["M","O","O","N"," ","+"," ","M","O","O","N"," ","+"," ","S","O","O","N"," ","=","="," ","J","U","N","E"]

  for (let e0 = 1; e0 < 10; e0++) {
    exp[0] = e0
    exp[7] = e0
    for (let e1 = 1; e1 < 10; e1++) {
      if (e1 === e0) continue
      exp[14] = e1
      for (let e2 = 1; e2 < 10; e2++) {
        if (e2 === e0) continue
        if (e2 === e1) continue
        exp[22] = e2
        for (let e3 = 0; e3 < 10; e3++) {
          if (e3 === e0) continue
          if (e3 === e1) continue
          if (e3 === e2) continue
          exp[1] = e3
          exp[2] = e3
          exp[8] = e3
          exp[9] = e3
          exp[15] = e3
          exp[16] = e3
          for (let e4 = 0; e4 < 10; e4++) {
            if (e4 === e0) continue
            if (e4 === e1) continue
            if (e4 === e2) continue
            if (e4 === e3) continue
            exp[3] = e4
            exp[10] = e4
            exp[17] = e4
            exp[24] = e4
            for (let e5 = 0; e5 < 10; e5++) {
              if (e5 === e0) continue
              if (e5 === e1) continue
              if (e5 === e2) continue
              if (e5 === e3) continue
              if (e5 === e4) continue
              exp[23] = e5
              for (let e6 = 0; e6 < 10; e6++) {
                if (e6 === e0) continue
                if (e6 === e1) continue
                if (e6 === e2) continue
                if (e6 === e3) continue
                if (e6 === e4) continue
                if (e6 === e5) continue
                exp[25] = e6

                if (eval(exp.join(''))) solutions.push(exp.join('').replace('==', '='))
              }
            }
          }
        }
      }
    }
  }
  return solutions
}
