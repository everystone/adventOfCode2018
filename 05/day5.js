
class Stack {
  constructor() {
    this.items = []
  }
  pop() {
    const item = this.items[this.items.length - 1]
    this.items = this.items.slice(0, this.items.length - 1)
    return item
  }
  peek() {
    return this.items[this.items.length - 1]
  }
  push(val) {
    this.items.push(val)
  }
  len() {
    return this.items.length
  }
}


var fs = require('fs');
fs.readFile('input.txt', 'utf8', function (err, input) {
  if (err) throw err;
  input = input.trimRight()

  const react = (str) => {
    let s = new Stack()
    for (let i = 0; i < str.length; i++) {
      c = str[i]
      if (s.len() == 0) {
        s.push(c)
      } else {
        last = s.peek()
        if (last == c.toLowerCase() && c == c.toUpperCase() ||
          last == c.toUpperCase() && c == c.toLowerCase()) {
          s.pop()
        } else {
          s.push(c)
        }
      }
    }
    return s.len()
  }

  function replaceAll(str, find, replace) {
    return str.replace(new RegExp(find, 'g'), replace);
  }
  const process = (str, unit) => {
    str = replaceAll(str, unit, "")
    str = replaceAll(str, unit.toUpperCase(), "")
    const result = react(str)
    console.log(`${unit}:  ${result}`)
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


