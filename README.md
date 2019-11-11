---
header-includes:
	\usepackage[czech]{babel}
	\usepackage{a4wide}
---
# Chemická jednotka

## Řešení

Na začátku si vzestupně seřadím dny ,kdy je otevřeno, podle dne kdy je otevřeno. Následně proiteruji každý den až do N. 
V každém dni zkontroluji zda nemá drogerie otevřeno. Pokud ano přidám si lahvičku do dvou seznamů. Jeden je seřazený 
podle ceny, druhý podle dne kdy látka vyprchá, ten získám tak, že ke dni kdy můžu lahvičku nakoupit přičtu trvanlivost. 
Počítám s tím, že až do tohoto dne je lahvička dostupná, a že potřebný počet nakoupím v ten jeden den kdy ji nabízejí. 
Potom vezmu tu lahvičku, jejíž cena je aktuálně nejnižší, je první v seznamu dle ceny. Následně odstraním lahvičky 
s prošlou dobou trvanlivosti z obou seznamů. Takto pokračuji dokud neprojdu všechny seznamy, nebo mi nedojdou lahvičky. 
Ve zdrojácích je trochu jiné pořadí operací, ale myšlenka je stejná.

### Složitost

Načtení dat má složitost `N`^2^, při načítání si dny řadím.
Následně každé přidání, či odebrání lahvičky má složitost `N`, procházím celý seznam.
A jelikož iteruji každý den, asymptomická složitost bude asi: `O(N`^2^`)`.
