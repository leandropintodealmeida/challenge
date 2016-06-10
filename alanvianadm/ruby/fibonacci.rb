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

puts "Digite o tamanho da sequencia de fibonacci desejada: "
seq = gets.to_i

seq.times do |number|
    result = fibonacci(number)
    puts result
end
