def fibonacci(number)
    a = 0
    b = 1

    number.times do
	temp = a
	a = b

	b = temp + b
    end

    return a
end

puts "Digite quantos números da sequência de Fibonacci você quer ver:"
seq = gets.to_i

puts "Você quer ver os #{seq} números da sequência."
seq.times do |number|
    result = fibonacci(number)
    puts result
end
