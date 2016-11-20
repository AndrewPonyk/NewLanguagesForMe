fib <- function(n) {
  if (n== 1 || n == 2){
    return (1);
  }
  return (fib(n-1) + fib(n-2))
}

for (i in 1:10){
  print(fib(i))
}

