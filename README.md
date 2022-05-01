# lsystemstr
Generate strings for L-systems (Lindenmayer systems)
https://en.wikipedia.org/wiki/L-system


How to generate a Sierpinski triangle sentence
```Go
ls := lsystemstr.New("F-G-G")
ls.AddRules(
  lsystemstr.NewRule("F", "F-G+F+G-F"),
  lsystemstr.NewRule("G", "GG"))
ls.Iterate(2)
ls.Sentence()
```
F-G+F+G-F-GG+F-G+F+G-F+GG-F-G+F+G-F-GGGG-GGGG
