void main() {
  int nb = 80;
  print('fibonacci($nb) = ${fibonacci(nb)}');
}

int fibonacci(int number) {
  return number < 2 ? number : ((number - 1) + fibonacci(number - 2));
}
