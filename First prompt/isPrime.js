function isPrime(number) {
  if (!Number.isSafeInteger(number) || number <= 0) {
    throw new TypeError("Input must be a positive safe integer");
  }

  if (number === 1) return false;

  for (let divisor = 2; divisor * divisor <= number; divisor++) {
    if (number % divisor === 0) return false;
  }

  return true;
}