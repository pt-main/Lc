# String Parsing

Package with string parsers, implements ParserInterface for engine: 

- `Lexer`: Simple lexer based on regexp2. Find tokens in full text. Supports multichar bracket balance.
- `Parser1`: Parser based on regexp. Find tokens in lines. Supports simple bracket balance, line continuation (with `line \ \n...`), and block space trimming.
- `Parser2`: The simplest one parser parsing grammar like `cmd args ...`.
- `Parser3`: The hardest one peg-like parser. See `parsing/stringParsing/parser3` for more docs