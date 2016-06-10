initialNumber = int(raw_input("Digite o primeiro numero:  "))
lastNumber = int(raw_input("Digite o ultimo numero "))

def fibonacci(n):
    a = 0
    b = 1
    for i in range(0, n):
	temp = a
	a = b
	b = temp + b
    return a

for c in range(initialNumber, lastNumber):
    print(fibonacci(c))
