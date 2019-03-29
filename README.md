
## ICA01

## log
Bruk: kompiler og kjør, eller kjør med go run.    

Syntaks: Main.go <log base (2 eller 10)> <ønsket tall>   

Log.go i pakken, functions, regner ut logaritmen til et valgfritt tall med base 2 eller base 10. Det krever to argumenter fra kommandolinjen,
parser tallet vi tar logaritmen av (tall) til float64 og basen (base) til int - begge fra string. Vi bruker en if-setning til å
sjekke hvilke base som skal brukes til å beregne logaritmen. Logaritme-funksjonen kjøres fra main.go. 

## logcli.go 
Tar inn et argument fra kommandolinja, parser det til float64 og regner ut logaritmen av argumentet med base 2. 
