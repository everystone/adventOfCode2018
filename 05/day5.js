var fs = require('fs');
fs.readFile('input.txt', 'utf8', function (err, input) {
  if (err) throw err;
  
  input = input.trimRight()
  const react = (str) => {
    for (let i = 0; i < str.length; i++) {
      if (i === str.Length - 1) {
        return [false, str]
      }
      s = str[i]
      if (s == s.toLowerCase() && str[i + 1] == s.toUpperCase() || s == s.toUpperCase() && str[i + 1] == s.toLowerCase()) {
        str = str.replace(s + str[i + 1], "")
        return [true, str]
      }
    }
    return [false, str]
  }

  function replaceAll(str, find, replace) {
    return str.replace(new RegExp(find, 'g'), replace);
  }
  const process = (str, unit) => {
    let run = true
    str = replaceAll(str, unit, "")
    str = replaceAll(str, unit.toUpperCase(), "")
    while (run) {
      [run, str] = react(str)
    }
    console.log(`${unit}:  ${str.length}`)
  }


  const chars = ['', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']
  console.log('read input: ', input.length)

  console.time("run")
  for (let l of chars) {
    copy = input.slice()
    process(copy, l)
  }
  console.timeEnd("run")
});


